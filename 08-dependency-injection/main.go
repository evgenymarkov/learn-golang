package main

import (
	"log"
	"net/http"

	"github.com/evgenymarkov/learn-golang/08-dependency-injection/di"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "Evgeny")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(GreetHandler)))
}
