package main

import "testing"

func AssertStatus(t *testing.T, gotStatusCode, expectedStatusCode int) {
	if gotStatusCode != expectedStatusCode {
		t.Fatalf("Request failed with statusCode: %d, expected: %d", gotStatusCode, expectedStatusCode)
	}
}
