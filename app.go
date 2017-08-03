package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetFavoritesClosure() func(w http.ResponseWriter, req *http.Request) {
	provider := FavTeamProvider{}

	provider.init()

	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)

		var favoriteTeams = provider.favoritesByZip(params["zip_code"])

		json.NewEncoder(w).Encode(favoriteTeams)
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/fan/{zip_code}", GetFavoritesClosure()).Methods("GET")

	log.Fatal(http.ListenAndServe(":2020", router))
}
