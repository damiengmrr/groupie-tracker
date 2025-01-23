package backgo

import (
	/*"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"*/
	"strings"
	//"text/template"
	"sort"
)

type List struct {
	Lists []artists
}

type artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    locations
	ConcertDates dates
	Relations    relations
}

type dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type relations struct {
	Id        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}

type locations struct {
	Id                   int      `json:"id"`
	Locations            []string `json:"locations"`
	Dates                string   `json:"dates"`
	Country              []string
	ConcertCityLocations []string
}

// filtre artistes par nom
func filterArtists(artistsList []artists, query string) []artists {
	var filtered []artists
	for _, artist := range artistsList {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			filtered = append(filtered, artist)
		}
	}
	return filtered
}

// trier les artistes par date de création
func sortArtistsByCreationDate(artistsList []artists, order string) []artists {
	if order == "croi" {
		sort.Slice(artistsList, func(i, j int) bool {
			return artistsList[i].CreationDate < artistsList[j].CreationDate
		})
	} else if order == "dcroi" {
		sort.Slice(artistsList, func(i, j int) bool {
			return artistsList[i].CreationDate > artistsList[j].CreationDate
		})
	}
	return artistsList
}
