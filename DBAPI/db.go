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
	var js []byte

	err := db.conn.QueryRow(context.Background(), sql, params...).Scan(&ok, &js)
	if err != nil {
		return false, nil
	}

	return ok, js
}

type Person struct {
	ID        int     `json:"id"`
	CreatedAt string  `json:"created_at"`
	Admin     bool    `json:"admin"`
	Name      string  `json:"name"`
	Things    []Thing `json:"things"`
}

type Thing struct {
	ID        int     `json:"id"`
	Person_IS int     `json:"person_id"`
	Name      string  `json:"name"`
	Price     float32 `json:"price"`
}

type Error struct {
	Error string `json:"error"`
}

func main() {
	API := NewDBAPI("zz")

	var person Person
	var error Error
	var people []Person
	var thing Thing

	fmt.Println("PERSON:")
	ok, data := API.a("person", 1)
	if !ok {
		json.Unmarshal(data, &error)
		fmt.Println(error)
	} else {
		json.Unmarshal(data, &person)
		fmt.Println(person)
	}

	ok, data = API.a("person", 2)
	if !ok {
		json.Unmarshal(data, &error)
		fmt.Println(error)
	} else {
		json.Unmarshal(data, &person)
		fmt.Println(person)
	}

	ok, data = API.a("person", 999)
	if !ok {
		json.Unmarshal(data, &error)
		fmt.Println(error)
	} else {
		json.Unmarshal(data, &person)
		fmt.Println(person)
	}

	fmt.Println("PEOPLE:")
	ok, data = API.a("people")
	if !ok {
		json.Unmarshal(data, &error)
		fmt.Println(error)
	} else {
		json.Unmarshal(data, &people)
		fmt.Println(people)
	}

	fmt.Println("GIVE CHARLIE CHOCOLATE: (only once)")
	ok, data = API.a("thing_add", 2, "chocolate", 3.75)
	if !ok {
		json.Unmarshal(data, &error)
		fmt.Println(error)
	} else {
		json.Unmarshal(data, &thing)
		fmt.Println(thing.ID)
	}
	if ok {
		ok, data = API.a("thing_add", 2, "chocolate", 3.75)
		if !ok {
			json.Unmarshal(data, &error)
			fmt.Println(error)
		} else {
			json.Unmarshal(data, &thing)
			fmt.Println(thing.ID)
		}
	}

	fmt.Println("GIVE CHARLIE LOVE: (overloaded, priceless)")
	ok, data = API.a("thing_add", 2, "love")
	if !ok {
		json.Unmarshal(data, &error)
		fmt.Println(error)
	} else {
		json.Unmarshal(data, &thing)
		fmt.Println(thing.ID)
	}
}
