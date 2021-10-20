package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vSterlin/bookstore/internal/server"
)

func main() {
	db, err := sql.Open("postgres", "user=v dbname=book-store sslmode=disable")
	// db, err := sql.Open("sqlite3", "file:test.db?cache=shared&mode=memory")

	if err != nil {
		fmt.Println(err)
	}

	s := server.NewServer(8080, db)
	s.Init()
	defer s.Shutdown()
}
