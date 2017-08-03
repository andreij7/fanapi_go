package main

import (
	"encoding/json"
	"io/ioutil"
)

//FavTeamProvider ...
type FavTeamProvider struct {
	objmap map[string]*json.RawMessage
}

func (f *FavTeamProvider) init() {
	file, e := ioutil.ReadFile("./fan_favs.json")

	if e == nil {
		json.Unmarshal(file, &f.objmap)
	}
}

func (f *FavTeamProvider) favoritesByZip(zipCode string) *json.RawMessage {

	if f.objmap != nil {
		var favoriteTeams = f.objmap[zipCode]

		return favoriteTeams
	}

	return nil

}
