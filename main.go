package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/thynquest/fizzbuzz/handlers"
)

var (
	frequency map[string]int
	logger    *logrus.Logger
)

func init() {
	frequency = make(map[string]int)
	logger = logrus.New()
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "02-01-2006 15:04:05"
	formatter.FullTimestamp = true
	logger.SetFormatter(formatter)
	logger.Out = os.Stdout
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	type freqSorted struct {
		Query string
		Hits  int
	}

	var fs []freqSorted
	for k, v := range frequency {
		fs = append(fs, freqSorted{k, v})
	}
	sort.Slice(fs, func(i, j int) bool {
		return fs[i].Hits > fs[j].Hits
	})
	if len(fs) > 0 {
		json.NewEncoder(w).Encode(fs[0])
		logger.Info("[frequency]: frequency request .. ok")
	} else {
		fmt.Fprintf(w, "no fizzbuzz request registered")
		logger.Info("[frequency]: no fizzbuzz request registered")
	}
}

func updateFrequency(multint1, multint2, limit, multstr1, multstr2 string) {
	key := strings.Join([]string{multint1, multint2, limit, multstr1, multstr2}, ",")
	frequency[key] = frequency[key] + 1
	logger.Info("[frequency]: Frequency updated")
}

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	multint1, _ := strconv.Atoi(parameters["multint1"])
	multint2, _ := strconv.Atoi(parameters["multint2"])
	limit, _ := strconv.Atoi(parameters["limit"])
	multstr1 := parameters["multstr1"]
	multstr2 := parameters["multstr2"]
	result, err := handlers.FizzBuzz(multint1, multint2, limit, multstr1, multstr2)
	if err != nil {
		logger.Errorf("[fizzbuzz]:error cause : %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updateFrequency(parameters["multint1"], parameters["multint2"], parameters["limit"], multstr1, multstr2)
	json.NewEncoder(w).Encode(result)
	logger.Infof("[fizzbuzz]: request %v, %v, %v, %v, %v .. ok", multint1, multint2, limit, multstr1, multstr2)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/fizzbuzz/{multint1:[0-9]+}/{multint2:[0-9]+}/{limit:[0-9]+}/{multstr1:[a-z]+}/{multstr2:[a-z]+}", fizzbuzzHandler).Methods("GET")
	router.HandleFunc("/stats", statsHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
	logger.Info("[main]: server has started")
}
