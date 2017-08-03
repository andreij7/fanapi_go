package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

//to run tests >>  go test

func TestFavoritesByZipCode72401(t *testing.T) {
	provider := FavTeamProvider{}

	var favoriteTeams = provider.favoritesByZip("72401")

	if favoriteTeams == nil {
		t.Errorf("Favorite teams should not be null for 72401")
	}

	j, e := json.Marshal(&favoriteTeams)

	if e == nil {
		fmt.Println(string(j))
	}

	//{"MLB":"/sport/baseball/team:2975","NCAAM":"/sport/basketball/team:2331","NCAAF":"/sport/football/team:75"}
}

func TestFavoriestByZipCode22182(t *testing.T) {
	provider := FavTeamProvider{}

	var favoriteTeams = provider.favoritesByZip("22182")

	if favoriteTeams == nil {
		t.Errorf("Favorite teams should not be null for 22182")
	}

	//{"MLB":"/sport/baseball/team:2959","NHL":"/sport/hockey/team:15","NBA":"/sport/basketball/team:404067","NFL":"/sport/football/team:27"}
}

func TestFavoritesByZipCodeBadZip(t *testing.T) {
	provider := FavTeamProvider{}

	var favoriteTeams = provider.favoritesByZip("2218255")

	if favoriteTeams != nil {
		t.Errorf("Should Be nil")
	}
}
