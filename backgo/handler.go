package backgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	// "time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./serv/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var GroupList list
	json.Unmarshal(responseData, &GroupList.Lists)
	t.Execute(w, GroupList)
}

func Artists(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./serv/artists.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	var GroupList []artists
	json.Unmarshal(responseData, &GroupList)

	id, _ := strconv.Atoi(r.FormValue("id"))
	t.Execute(w, GroupList[id-1])
}

func SearchBar(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("./serv/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	SearchValue := r.FormValue("RechercheArtiste")
	
	var GroupList []artists
	json.Unmarshal(responseData, &GroupList)
	

	for _, artist := range GroupList {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(SearchValue)) {
			...
			fmt.Print("GG")
		}
	}
}
var GroupList List
    json.Unmarshal(body, &GroupList.Lists)
    queryYear := r.URL.Query().Get("year")
    queryYearAlbum := r.URL.Query().Get("yearAlbum")
    queryMembers := r.URL.Query().Get("members")

    var filteredArtists []Artiste
    for , artist := range GroupList.Lists {
        searchQuery := r.URL.Query().Get("query")
        if searchQuery != "" {
            searchQuery = strings.ToLower(searchQuery)
            tempFilteredArtists := []Artiste{}
            for , artist := range filteredArtists {
                if strings.Contains(strings.ToLower(artist.Name), searchQuery) {
                    tempFilteredArtists = append(tempFilteredArtists, artist)
                }
            }
            filteredArtists = tempFilteredArtists
        }