package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
)

type Promotion struct {
	ID             string `json:"id"`
	Price          string `json:"price"`
	ExpirationDate string `json:"expiration_date"`
}

var client *redis.Client

func main() {
	redisHost := os.Getenv("REDIS_ADDR")
	client = redis.NewClient(&redis.Options{
		Addr:     redisHost, // Change to your Redis instance address
		Password: "",        // Change to your Redis instance password
		DB:       0,
	})

	loadPromotions("data/promotions.csv", client)

	http.HandleFunc("/promotions/", getPromotion)
	log.Fatal(http.ListenAndServe(":1321", nil))
}

func getPromotion(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Path[len("/promotions/"):]

	promotionJSON, err := client.Get(idParam).Result()
	if err == redis.Nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(promotionJSON))
}

func loadPromotions(filename string, client *redis.Client) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Clear existing promotions in Redis
	client.FlushDB()

	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		promotion := Promotion{
			ID:             record[0],
			Price:          record[1],
			ExpirationDate: record[2],
		}

		// Set promotion in Redis with 30-minute expiration time
		err = client.Set(promotion.ID, toJSON(promotion), 30*time.Minute).Err()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func toJSON(p Promotion) string {
	b, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
