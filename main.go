package main

import (
	"golang-ftth/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/analysis", handler.AnalysisHandler)
	log.Println("Server listening on port 15442")
	log.Fatal(http.ListenAndServe(":15442", nil))
}
