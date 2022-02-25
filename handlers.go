package main

import (
	"fmt"
	"net/http"
)

func SayHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

// r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	title := vars["title"]
// 	page := vars["page"]

// 	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
// 	SayHello()
// })

// return await deps.interactors.chef.dishes.prepare(
// 	deps,
// 	user_id,
// 	recipe_name,
// 	{
// 		has_fast_forwards,
// 		has_easy_tiger,
// 		first_plus_purchase,
// 	}
// );
func GetRecommendationsHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	userId := queryParams.Get("user_id")
	recipeName := queryParams.Get("recipe_name")

	fmt.Println(userId)
	fmt.Println(recipeName)

	// fmt.Println(query)
}
