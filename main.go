package main

import (
	g "groupie-tracker/backgo"
	"fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"

)

func main(){
	response, err := http.Get("https://groupietrackers.herokuapp.com/api")
	if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
	g.Web()
}