package chat

import (
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("./../../templates/index.html")
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
