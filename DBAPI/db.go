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
	conn, err := pgx.Connect(context.Background(), "host=127.0.0.1 dbname=s user=s")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return &DBAPI{
		conn:   conn,
		schema: schema,
	}
}

func (db *DBAPI) a(funcName string, params ...interface{}) (bool, JSONResult) {
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
	var js []byte

	err := db.conn.QueryRow(context.Background(), sql, params...).Scan(&ok, &js)
	if err != nil {
		return false, JSONResult{}
	}

	var result interface{}
	err = json.Unmarshal(js, &result)
	if err != nil {
		fmt.Println(err)
	}

	switch result.(type) {
	case []interface{}:
		// convert []interface{} to []map[string]interface{}
		var array []map[string]interface{}
		for _, item := range result.([]interface{}) {
			array = append(array, item.(map[string]interface{}))
		}
		return true, JSONResult{A: array}
	case map[string]interface{}:
		return true, JSONResult{O: result.(map[string]interface{})}
	}

	return false, JSONResult{}
}

func main() {
	API := NewDBAPI("zz")

	fmt.Println("PERSON:")
	_, person := API.a("person", 1)
	fmt.Println(person.O["name"])
	_, person = API.a("person", 2)
	fmt.Println(person.O["name"])
	_, person = API.a("person", 999)
	fmt.Println(person.O["error"])

	fmt.Println("PEOPLE:")
	_, people := API.a("people")
	fmt.Println(people.A[0]["name"])

	fmt.Println("GIVE CHARLIE CHOCOLATE: (only once)")
	_, thing := API.a("thing_add", 2, "chocolate", 3.75)
	fmt.Println(thing)

	_, thing = API.a("thing_add", 2, "chocolate", 3.75)
	fmt.Println(thing)

	fmt.Println("GIVE CHARLIE LOVE: (overloaded, priceless)")
	_, thing = API.a("thing_add", 2, "love")
	fmt.Println(thing)
}
