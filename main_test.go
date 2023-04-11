package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
)

func TestMain(m *testing.M) {
	// Start a Redis server for testing
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mr.Close()

	// Update Redis configuration to use the test server
	client = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
		DB:   0,
	})

	// Load test data into Redis
	loadPromotions("data/test_data.csv", client)

	// Run tests
	exitCode := m.Run()

	// Exit with the code returned by the tests
	os.Exit(exitCode)
}

func TestGetPromotion(t *testing.T) {
	req, err := http.NewRequest("GET", "/promotions/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getPromotion)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"id":"1","price":"10.00","expiration_date":"2023-05-01"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetNonexistentPromotion(t *testing.T) {
	req, err := http.NewRequest("GET", "/promotions/999", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getPromotion)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
func TestLoadPromotions(t *testing.T) {
	// Start a Redis server for testing
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}
	defer mr.Close()

	// Update Redis configuration to use the test server
	testClient := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
		DB:   0,
	})

	// Load test data into Redis
	loadPromotions("data/test_data.csv", testClient)

	// Verify that the expected data was loaded into Redis
	expectedPromotions := map[string]Promotion{
		"1": {
			ID:             "1",
			Price:          "10.00",
			ExpirationDate: "2023-05-01",
		},
		"2": {
			ID:             "2",
			Price:          "20.00",
			ExpirationDate: "2023-05-01",
		},
		"3": {
			ID:             "3",
			Price:          "30.00",
			ExpirationDate: "2023-05-01",
		},
	}
	for id, expectedPromotion := range expectedPromotions {
		promotionJSON, err := testClient.Get(id).Result()
		if err != nil {
			t.Errorf("error getting promotion from Redis: %v", err)
		}
		var actualPromotion Promotion
		err = json.Unmarshal([]byte(promotionJSON), &actualPromotion)
		if err != nil {
			t.Errorf("error unmarshalling promotion JSON: %v", err)
		}
		if !reflect.DeepEqual(actualPromotion, expectedPromotion) {
			t.Errorf("promotion with ID %s did not match expected value; got %v, expected %v", id, actualPromotion, expectedPromotion)
		}
	}
}
