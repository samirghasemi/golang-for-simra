package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samirghasemi/golang-for-simra/lottery/api/handler"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/lottery", handler.LotteryHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
