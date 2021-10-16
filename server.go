package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	r *mux.Router
}

func (s *server) init(addr string) {

	http.ListenAndServe(addr, s.r)
}
