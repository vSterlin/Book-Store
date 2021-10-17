package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vSterlin/bookstore/controller"
	"github.com/vSterlin/bookstore/repository"
	"github.com/vSterlin/bookstore/service"
)

func main() {

	br := &repository.BookRepo{}
	bs := &service.BookService{Br: br}
	bc := &controller.BookController{Bs: bs}

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
