package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type messages struct {
	Num1 string `json:"num1"`
}

type messResponseData struct {
	Ans   string `json:"Ans"`
	Class string `json:"Class"`
}

func calc(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var messData messages
	var messResData messResponseData

	decoder.Decode(&messData)

	messResData = process(messData)

	fmt.Println(messResData)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(messResData); err != nil {
		panic(err)
	}
}

func main() {

	fmt.Println("Starting...")

	router := chi.NewRouter()
	router.Post("/calc", calc)

	log.Fatal(http.ListenAndServe(":8090", router))
}
