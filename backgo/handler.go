package backgo

import (
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, html string) {
	t, err := template.ParseFiles("./serv/" + html + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w,nil)
}

func Home(w http.ResponseWriter, r *http.Request){
	RenderTemplate(w, "home")
}

