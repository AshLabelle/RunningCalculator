package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type CalculateTimeRequest struct {
	Distance	string `json:distance`
	Pace		string `json:pace`
}

func calculateTimeHandler(w http.ResponseWriter, r *http.Request) {

	var decoded CalculateTimeRequest

	// Try to decode the request into the request struct
	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Marshal(decoded);

	p, err := time.ParseDuration(decoded.Pace)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	min := math.Floor(p.Minutes())
	sec := math.Mod(p.Seconds(), 60)
	fmt.Println(min)
	fmt.Println(sec)

	dist, err := strconv.ParseFloat(decoded.Distance, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	totalSec := (min * 60 + sec) * dist
	totalDuration := time.Duration(math.Round(totalSec) *1e9)
	fmt.Println(totalDuration)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{ "time": "%s", "distance": "%s", "pace": "%s" }`, totalDuration, decoded.Distance, decoded.Pace)))
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{ "message": "%s" }`, "Hello!")))
}

func main() {
	router := mux.NewRouter()	
	router.HandleFunc("/", ping).Methods("GET")
	router.HandleFunc("/api/calculateTime", calculateTimeHandler).Methods("POST")
	http.ListenAndServe(":3000", 
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"POST", "OPTIONS", "GET", "DELETE", "PUT"}),
			handlers.AllowedHeaders([]string{"Access-Control-Allow-Headers", "Content-Type", "Authorization", "Content-Length", "X-Requested-With", "Accept"}),
		) (router))

	fmt.Println("Server listening on port 3000")
}