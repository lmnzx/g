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

func (db *DBAPI) a(funcName string, params ...interface{}) (bool, interface{}, error) {
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
		return false, nil, err
	}

	// parse the json result
	// cannot be of type map[string]interface{} as zz.people() returns an array
	var result interface{}
	err = json.Unmarshal(js, &result)
	if err != nil {
		return false, nil, err
	}

	return ok, result, nil

}

func main() {
	API := NewDBAPI("zz")

	fmt.Println("PERSON:")
	ok, person, err := API.a("person", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ok, person)

	ok, person, _ = API.a("person", 2)
	fmt.Println(ok, person)

	ok, person, _ = API.a("person", 999)
	fmt.Println(ok, person)

	fmt.Println("PEOPLE:")
	ok, people, err := API.a("people")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ok, people)

	fmt.Println("GIVE CHARLIE CHOCOLATE: (only once)")
	ok, n, err := API.a("thing_add", 2, "chocolate", 3.75)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ok, n)
	if ok {
		ok, n, err := API.a("thing_add", 2, "chocolate", 3.75)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(ok, n)
	}

	fmt.Println("GIVE CHARLIE LOVE: (overloaded, priceless)")
	ok, n, _ = API.a("thing_add", 2, "love")
	fmt.Println(ok, n)

}
