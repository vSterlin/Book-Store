package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/vSterlin/bookstore/internal/server"
)

func main() {
	db, _ := sql.Open("postgres", "user=v dbname=book-store sslmode=disable")
	s := server.NewServer(":8080", db)
	s.Init()
	defer s.Shutdown()
}
