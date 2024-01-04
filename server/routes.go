package server

import (
	"net/http"

	"github.com/thynquest/fizzbuzz/pkg/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const ()

func DefineHandlers(f *fizzbuzzServer) {
	svc := service.NewFizzBuzz()
	h := NewHandler(svc)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("checking root.. Ok"))
	})
	r.Get("/fizzbuzz/{multint1:[0-9]+}/{multint2:[0-9]+}/{limit:[0-9]+}/{multstr1:[a-z]+}/{multstr2:[a-z]+}", h.FizzBuzzHandler)
	r.Get("/stats", h.StatsHandler)
	f.server.Handler = r
}
