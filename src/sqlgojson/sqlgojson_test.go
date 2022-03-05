package sqlgojson

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func queryPostgresql(query string) (*sql.Rows, error) {
	connStr := "user=root dbname=sqljson password=test host=localhost sslmode=disable port=5432"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database \n")
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	return rows, err
}

func queryMysql(query string) (*sql.Rows, error) {
	connStr := "root:root@/sqljson"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database \n")
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	return rows, err
}

func TestMysqlTable(t *testing.T) {
	rows, err := queryMysql("SELECT * from test;")
	if err != nil {
		panic(err)
	}
	res, err := SqlGoJson(rows)
	if err != nil {
		panic(err)
	}
	expected := `[[{"id":1},{"num":3.14},{"b":true}],[{"id":2},{"num":4.014},{"b":false}]]`
	fmt.Println(res, expected)
	if res != expected {
		t.Errorf("Failed")
	}
	fmt.Printf("rep:%v", res)
}

func TestUserTable(t *testing.T) {
	rows, err := queryPostgresql("SELECT * from users;")
	if err != nil {
		panic(err)
	}
	res, err := SqlGoJson(rows)
	if err != nil {
		panic(err)
	}
	if res != `[[{"id":1},{"name":"nolan"},{"created_at":"2017-07-01T14:59:55.711Z"}],[{"id":2},{"name":"test"},{"created_at":"2017-07-01T15:59:55.711Z"}],[{"id":3},{"name":"test"},{"created_at":null}]]` {
		t.Errorf("Failed")
	}
	fmt.Printf("rep:%v", res)
}

func TestTestTable(t *testing.T) {
	rows, err := queryPostgresql("SELECT * from test;")
	if err != nil {
		panic(err)
	}
	res, err := SqlGoJson(rows)
	if err != nil {
		panic(err)
	}
	expected := `[[{"id":1},{"num":3.14},{"b":true}],[{"id":2},{"num":4.014},{"b":false}]]`
	fmt.Println(res, expected)
	if res != expected {
		t.Errorf("Failed")
	}
	fmt.Printf("rep:%v", res)
}
