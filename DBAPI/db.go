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

func (db *DBAPI) a(funcName string, params ...interface{}) (bool, interface{}) {
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

	// parse the json result
	// cannot be of type map[string]interface{} as zz.people() returns an array
	var result interface{}
	err = json.Unmarshal(js, &result)
	if err != nil {
		return false, nil
	}

	return ok, result
}

func main() {
	API := NewDBAPI("zz")

	fmt.Println("PERSON:")
	ok, person := API.a("person", 1)
	fmt.Println(ok, person)

	ok, person = API.a("person", 2)
	fmt.Println(ok, person)

	ok, person = API.a("person", 999)
	fmt.Println(ok, person)

	fmt.Println("PEOPLE:")
	ok, people := API.a("people")
	fmt.Println(ok, people)

	fmt.Println("GIVE CHARLIE CHOCOLATE: (only once)")
	ok, n := API.a("thing_add", 2, "chocolate", 3.75)
	fmt.Println(ok, n)
	if ok {
		ok, n := API.a("thing_add", 2, "chocolate", 3.75)
		fmt.Println(ok, n)
	}

	fmt.Println("GIVE CHARLIE LOVE: (overloaded, priceless)")
	ok, n = API.a("thing_add", 2, "love")
	fmt.Println(ok, n)
}
