package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type todo struct {
	Name string `json:"name"`
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/gettodos", getTodos)
	fmt.Println("listening on 3000")
	http.ListenAndServe(":3000", nil)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/index.html"))
	tmpl.ExecuteTemplate(w, "index.html", struct{ PageTitle string }{PageTitle: "Go Example"})
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	todo1 := &todo{
		Name: "go to the grocery store",
	}
	todo2 := &todo{
		Name: "vacuum",
	}
	todo3 := &todo{
		Name: "learn go",
	}
	todos := [3]todo{*todo1, *todo2, *todo3}
	data, err := json.Marshal(todos)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
