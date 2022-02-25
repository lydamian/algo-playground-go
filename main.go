package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var PORT string = "6000"
	muxRouter := mux.NewRouter()

	RegisterRoutes(muxRouter)

	log.Printf("Starting server... on PORT:%s\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), muxRouter))

}
