package backgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type artists struct {
	Id		int			`json:"id"`
	Image	string		`json:"image"`
	Name	string		`json:"name"`
	Members	[]string	`json:"members"`
	CreationDate int	`json:"creationDate"`
	FirstAlbum string	`json:"firstAlbum"`
	Locations locations
	ConcertDates dates
	Relations relations
}

type dates struct {
	Id			int				`json:"id"`
	Dates		[]string		`json:"dates"`
}

type relations struct {
	Id			int						`json:"id"`
	Relations	map[string][]string		`json:"datesLocations"`
}

type locations struct {
	Id			int				`json:"id"`
	Locations	[]string		`json:"locations"`
	Dates		string			`json:"dates"`
	Country []string
	ConcertCityLocations []string
}

func GetAPI(){
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

	var responseObject artists
	json.Unmarshal(responseData, &responseObject)


    fmt.Println(responseObject.Name)
    fmt.Println(len(responseObject.Pokemon))

    for i := 0; i < len(responseObject.Pokemon); i++ {
        fmt.Println(responseObject.Pokemon[i].Species.Name)
    }

}
