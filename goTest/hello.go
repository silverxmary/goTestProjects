package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))

	fmt.Println("Hola Mundo!!!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola aplicacion http!!! %s ", r.URL.Path)
	})

	http.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {

		tmpl.Execute(w, struct{ Saludo string }{"Adios Mary"})

	})

	http.ListenAndServe(":8080", nil)

}
