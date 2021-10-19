package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
	"github.com/vSterlin/bookstore/internal/controller"
	"github.com/vSterlin/bookstore/internal/repository"
	"github.com/vSterlin/bookstore/internal/service"
)

func main() {

	db, _ := sql.Open("postgres", "user=v dbname=book-store sslmode=disable")
	db.Ping()

	br := repository.NewBookRepo(db)
	bs := service.NewBookService(br)
	bc := controller.NewBookController(bs)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/books", func(r chi.Router) {
		r.Get("/", bc.GetBooks)
		r.Get("/{id}", bc.GetBook)
		r.Post("/", bc.InsertBook)
		r.Delete("/{id}", bc.DeleteBook)
	})

	http.ListenAndServe(":8080", r)
}
