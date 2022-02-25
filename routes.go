package main

import (
	"github.com/gorilla/mux"
)

// method: 'GET',
// path: '/api/v2/client/users/{user_id}/recommendations',
// config: {
// 	auth: {
// 		mode: 'required',
// 		strategy: 'client-status',
// 		scope: 'accepted-{params.user_id}-user',
// 	},
// 	handler: handlers.client.recommendations.social_get_v2,
// 	validate: {
// 		params: {
// 			user_id: Joi.string().required(),
// 		},
// 		query: {
// 			first_plus_purchase: Joi.boolean().optional().default(false),
// 		},
// 	},
// },

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/hello", SayHelloHandler).Methods("GET")
	r.HandleFunc("/recommendations", GetRecommendationsHandler).Methods("GET")
}
