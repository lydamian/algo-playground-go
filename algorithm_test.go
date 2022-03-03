package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/go-querystring/query"
)

// TestGetRecommendations makes a http request
// to our GET recommendations route, and asserts
// we get a valid response
func TestGetRecommendations(t *testing.T) {

	t.Run("it returns 200 on GET /api/recommendations", func(t *testing.T) {
		type queryParams struct {
			userId string `fake:"{number:1,10000000000}"`
			recipe string `fake:"{sentence:3}"`
		}
		var q queryParams
		gofakeit.Struct(&q)
		v, _ := query.Values(q)
		req := httptest.NewRequest(http.MethodGet, "/api/recommendations"+v.Encode(), nil)
		res := httptest.NewRecorder()
		GetRecommendationsHandler(res, req)
		resJSON := res.Body.String()
		fmt.Println(resJSON)

		type responseBody struct {
			Recommendations Recommendations
		}
		var got responseBody

		err := json.NewDecoder(res.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into responseBody, '%v'", w.Body, err)
		}
		AssertStatus(t, res.Code, http.StatusOK)
	})
}
