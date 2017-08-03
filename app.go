package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetFavoriteTeams(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	var zipCode = params["zip_code"]

	//read file to map
	file, e := ioutil.ReadFile("./fan_favs.json")

	if e == nil {
		//get value for key zipcode
		var objmap map[string]*json.RawMessage

		json.Unmarshal(file, &objmap)

		var favoriteTeams = objmap[zipCode]

		json.NewEncoder(w).Encode(favoriteTeams)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/fan/{zip_code}", GetFavoriteTeams).Methods("GET")
	log.Fatal(http.ListenAndServe(":2020", router))
}
