package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-bristol/beer-model"
	"github.com/julienschmidt/httprouter"
)

// GetBeers returns the cellar
func GetBeers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Cellar)
}

// GetBeer returns a beer from the cellar
func GetBeer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	ID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(fmt.Sprintf("%s is not a valid Beer ID, it must be a number.", ps.ByName("id")))
		return
	}

	for _, v := range Cellar {
		if v.ID == ID {
			json.NewEncoder(w).Encode(v)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("The beer you requested does not exist.")
}

// GetBeerReviews returns all reviews for a beer
func GetBeerReviews(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO
}

// AddBeer adds a new beer to the cellar
func AddBeer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var newBeer model.Beer
	err := decoder.Decode(&newBeer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(http.StatusBadRequest)
		fmt.Println("Bad beer - this will be a HTTP status code soon!")
	} else {
		json.NewEncoder(w).Encode("New beer added.")
	}
}

// AddBeerReview adds a new review for a beer
func AddBeerReview(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO
}
