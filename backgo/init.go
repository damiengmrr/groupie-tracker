package backgo

import (
	"net/http"
	"fmt"
	)

const port = ":8080"

func Web() {
	http.HandleFunc("/" , Home)
	http.HandleFunc("/artists" , Artists)

	fs := http.FileServer(http.Dir("serv/"))
	http.Handle("/serv/", http.StripPrefix("/serv/", fs))

	fmt.Println("(http://localhost:8080) - server started on port", port)
	http.ListenAndServe(port, nil)

}