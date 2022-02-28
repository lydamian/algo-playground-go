package main

import (
	"fmt"
	"log"
	"net/http"
)

func PingPongHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
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
	// get query params
	queryParams := r.URL.Query()
	userId := queryParams.Get("user_id")
	recipeName := queryParams.Get("recipe_name")

	type responseBody struct {
		Recommendations Recommendations
	}

	recommendations, err := GetRecommendationsInteractor(userId, recipeName)
	if err != nil {
		log.Fatalf("Error getting recommendations: %s", err)
		respondWithError(w, http.StatusInternalServerError, "Error getting recommendations")
		return
	}
	resp := responseBody{
		Recommendations: recommendations,
	}
	err = respondWithJSON(w, http.StatusOK, resp)
	if err != nil {
		fmt.Println(err)
	}
}
