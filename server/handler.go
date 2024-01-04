package server

import (
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	domain "github.com/thynquest/fizzbuzz/pkg/entity"
	"github.com/thynquest/fizzbuzz/pkg/logging"
	"github.com/thynquest/fizzbuzz/pkg/service"
	"github.com/thynquest/fizzbuzz/server/status"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var handlerTitle = "[fizzbuzz-api-handler]"

type handler struct {
	fb service.FizzBuzz
}

func NewHandler(fizzbuzz service.FizzBuzz) *handler {
	return &handler{
		fb: fizzbuzz,
	}
}

func (h *handler) FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	int1, _ := strconv.Atoi(chi.URLParam(r, "multint1"))
	int2, _ := strconv.Atoi(chi.URLParam(r, "multint2"))
	limit, _ := strconv.Atoi(chi.URLParam(r, "limit"))
	str1 := chi.URLParam(r, "multstr1")
	str2 := chi.URLParam(r, "multstr2")

	data := domain.FizzBuzz{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  str1,
		Str2:  str2,
	}
	response, err := h.fb.FizzBuzz(data.Int1, data.Int2, data.Limit, data.Str1, data.Str2)
	if err != nil {
		logging.Error(handlerTitle, fmt.Sprintf("error on fizzbuzz operation: %v", err))
		render.Render(w, r, status.ErrRender("fizzbuzz failed", errors.New("fizzbuzz operation failed")))
		return
	}
	render.JSON(w, r, response)
}

func (h *handler) StatsHandler(w http.ResponseWriter, r *http.Request) {
	response := h.fb.GetStats()
	if len(response) == 0 {
		logging.Warning(handlerTitle, "no fizzbuzz request registered")
		return
	}
	type freqSorted struct {
		Query string
		Hits  int
	}

	var fs []freqSorted
	for k, v := range response {
		fs = append(fs, freqSorted{k, v})
	}
	sort.Slice(fs, func(i, j int) bool {
		return fs[i].Hits > fs[j].Hits
	})
	render.JSON(w, r, fs[0])
}
