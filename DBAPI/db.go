package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type DBAPI struct {
	conn   *pgx.Conn
	schema string
}

type JSONResult struct {
	O map[string]interface{}
	A []map[string]interface{}
}

func NewDBAPI(schema string) *DBAPI {
	conn, err := pgx.Connect(context.Background(), "dbname=s user=s")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return &DBAPI{
		conn:   conn,
		schema: schema,
	}
}

func (db *DBAPI) a(funcName string, params ...interface{}) (bool, []byte) {
	qs := "("
	for i := range params {
		if i > 0 {
			qs += ","
		}
		qs += fmt.Sprintf("$%d", i+1)
	}
	qs += ")"

	sql := fmt.Sprintf("SELECT ok, js FROM %s.%s%s", db.schema, funcName, qs)

	var ok bool
	var js []byte // valid json from postgress

	err := db.conn.QueryRow(context.Background(), sql, params...).Scan(&ok, &js)
	if err != nil {
		return false, nil
	}

	return true, js
}

// generic function to decode json
func decodeJSON[T any](data []byte) (T, error) {
	var result T
	var temp interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return result, err
	}

	decoded, err := castToType[T](temp)
	if err != nil {
		return result, err
	}

	return decoded, nil
}

// cast the decoded json to the desired type
func castToType[T any](data interface{}) (T, error) {
	var result T
	bytes, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(bytes, &result)
	return result, err
}

func getNestedField[T any](data interface{}, keys ...string) (T, error) {
	var result T

	if len(keys) == 0 {
		return castToType[T](data)
	}

	switch d := data.(type) {
	case map[string]interface{}:
		if val, found := d[keys[0]]; found {
			return getNestedField[T](val, keys[1:]...)
		} else {
			return result, fmt.Errorf("key %s not found", keys[0])
		}
	case []interface{}:
		var results []T
		for _, item := range d {
			result, err := getNestedField[T](item, keys...)
			if err != nil {
				return result, err
			}
			results = append(results, result)
		}
		return castToType[T](results)
	default:
		return result, fmt.Errorf("unexpected type %T", d)
	}
}

func main() {
	API := NewDBAPI("zz")

	fmt.Println("PEOPLE:")
	ok, jsonData := API.a("people")
	if ok {
		var result []map[string]interface{}
		result, err := decodeJSON[[]map[string]interface{}](jsonData)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
		} else {
			fmt.Println("Decoded JSON:", result[0]["name"])
		}
	}

	fmt.Println("PERSON:")
	ok, jsonData = API.a("person", 1)
	if ok {
		var result map[string]interface{}
		result, err := decodeJSON[map[string]interface{}](jsonData)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
		} else {
			fmt.Println("Decoded JSON:", result)

			// getting the things
			var things []map[string]interface{}
			things, err := getNestedField[[]map[string]interface{}](result, "things")
			if err != nil {
				fmt.Println("Error accessing nested field:", err)
			} else {
				for i, thing := range things {
					fmt.Printf("%d -> %v\n", i, thing["name"])
				}
			}
		}
	}
}
