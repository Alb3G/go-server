package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	age := r.FormValue("age")
	project := r.FormValue("project")

	fmt.Fprintf(w, "name = %v\n", name)
	fmt.Fprintf(w, "age = %v\n", age)
	fmt.Fprintf(w, "Project in course -> %v", project)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "post" {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	file_server := http.FileServer(http.Dir("./static"))
	http.Handle("/", file_server)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
