package apiserver

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/vSterlin/bookstore/internal/controller"
	"github.com/vSterlin/bookstore/internal/repository"
	"github.com/vSterlin/bookstore/internal/service"
)

type Server struct {
	addr string
	db   *sql.DB
}

func NewServer(addr int, db *sql.DB) *Server {

	strAddr := strconv.Itoa(addr)

	return &Server{addr: strAddr, db: db}
}

func (s *Server) Init() {
	s.db.Ping()

	br := repository.NewBookRepo(s.db)
	bs := service.NewBookService(br)
	bc := controller.NewBookController(bs)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/books", func(r chi.Router) {
		r.Get("/", bc.GetBooks)
		r.Get("/{id}", bc.GetBook)
	})

	http.ListenAndServe(":"+s.addr, r)
}

func (s *Server) Shutdown() {
	s.db.Close()

}
