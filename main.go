package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("hello world")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "hello world")
		// io.WriteString(w, r.Method)  //for startig first
		temp1 := template.Must(template.ParseFiles("index.html"))
		//temp1.Execute(w, nil) //second starting

		films := map[string][]film{
			"films": {
				{Title: "the godfather", Director: "harsh"},
				{Title: "blade runner", Director: "makadiya"},
				{Title: "the thing", Director: "patel"},
			},
		}
		temp1.Execute(w, films)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
