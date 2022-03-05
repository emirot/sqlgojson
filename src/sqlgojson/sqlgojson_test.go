package sqlgojson

import (
	"database/sql"
	"fmt"
	"testing"
)

func queryPostgresql() (*sql.Rows, error) {
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
	rows, err := db.Query("SELECT * from users;")
	if err != nil {
		panic(err)
	}
	return rows, err
}

func TestAddingNumber(t *testing.T) {
	rows, err := queryPostgresql()
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
