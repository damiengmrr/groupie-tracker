package backgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request){
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
	var GroupList List
	json.Unmarshal(responseData, &GroupList.Lists)
	t.Execute(w,GroupList)
}


