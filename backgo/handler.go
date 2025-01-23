package backgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./serv/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	searchQuery := r.URL.Query().Get("search") // recherche par nom
	sortOrder := r.URL.Query().Get("sort")     // "croi" ou "dcroi"

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var GroupList List
	json.Unmarshal(responseData, &GroupList.Lists)

	// Filtrer artistes si rechercher
	if searchQuery != "" {
		GroupList.Lists = filterArtists(GroupList.Lists, searchQuery)
	}

	// Trier par date de création
	if sortOrder != "" {
		GroupList.Lists = sortArtistsByCreationDate(GroupList.Lists, sortOrder)
	}

	t.Execute(w, GroupList)
}

func Artists(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./serv/artists.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
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
