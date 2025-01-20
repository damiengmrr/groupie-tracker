package backgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	// "strings"
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


// func applyFilter(items []Item, filter string) []Item {
// 	var filteredItems []Item
// 	filter = strings.ToLower(filter)

// 	for _, item := range items {
// 		if strings.Contains(strings.ToLower(item.Title), filter) {
// 			filteredItems = append(filteredItems, item)
// 		}
// 	}

// 	return filteredItems
// }