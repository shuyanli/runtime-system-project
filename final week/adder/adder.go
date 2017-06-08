package main

import (
	//"bytes"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type add struct {
	Sum int
}

var templates = template.Must(template.ParseFiles("template.html"))

func handleStructAdd(w http.ResponseWriter, r *http.Request) {
	//get rid of the intermediate objects.
	//instead of using a buffer, we pass the IO.write directly to templates.Execute()


	//var html bytes.Buffer
	first, second := r.FormValue("first"), r.FormValue("second")
	one, err := strconv.Atoi(first)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	two, err := strconv.Atoi(second)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	m := struct{ a, b int }{one, two}
	structSum := add{Sum: m.a + m.b}

	//err = templates.Execute(&html, structSum)
	err = templates.Execute(w, structSum)
	
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	//get rid of intermediate objects
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//w.Write([]byte(html.String()))
}

func main() {

	http.HandleFunc("/struct", handleStructAdd)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}