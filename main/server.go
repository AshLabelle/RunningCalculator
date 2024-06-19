package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type calculateTimeRequest struct {
	distance	string	`json:distance`
	pace		string	`json:pace`
}

func calculateTimeHandler(w http.ResponseWriter, r *http.Request) {

	var decoded calculateTimeRequest

	// Try to decode the request into the request struct
	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// paceSplit := strings.Split(decoded.pace, ":")
	// paceMin, err := strconv.Atoi(paceSplit[0])
	// paceSec, err := strconv.Atoi(paceSplit[1])

	// var result = paceMin *60 + paceSec * decoded.distance

	// Pass back the screenshot URL to the frontend.
	fmt.Println("Returning Response")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{ "time": "%s", "distance": "%s", "pace": "%s" }`, "600", decoded.distance, decoded.pace)))
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
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