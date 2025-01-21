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
	// "time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./serv/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
	searchQuery := r.URL.Query().Get("search")

	
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

	
	t.Execute(w, GroupList)
}

	
func Artists(w http.ResponseWriter, r *http.Request){
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
	t.Execute(w,GroupList[id-1])
}
