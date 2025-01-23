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

	searchQuery := r.URL.Query().Get("search")
	sortOrder := r.URL.Query().Get("sort")

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

	if searchQuery != "" {
		GroupList.Lists = filterArtists(GroupList.Lists, searchQuery)
	}
	if sortOrder != "" {
		GroupList.Lists = tri_selection(GroupList.Lists, sortOrder)
	}

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

func filterArtists(artistsList []artists, query string) []artists {
	var filtered []artists
	for _, artist := range artistsList {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			filtered = append(filtered, artist)
		}
	}
	return filtered

}

func tri_selection(artistsList []artists, order string) []artists {
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

	for i := range GroupList {
		mini := i
		for j := i + 1; j < len(GroupList); j++ {
			if GroupList[j].CreationDate < GroupList[mini].CreationDate {
				mini = j
			}
		}
		if mini != i {
			GroupList[i], GroupList[mini] = GroupList[mini], GroupList[i]
			tmp := GroupList[i]
			GroupList[i] = GroupList[mini]
			GroupList[mini] = tmp
		}
	}
	return GroupList
}
