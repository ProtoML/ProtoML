package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Call Unrecognized")
}

func transform(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expecting GET on all transforms")
}

func transformID(w http.ResponseWriter, r *http.Request) {
	const transPath = len("/transform/")
	id := r.URL.Path[transPath:]
	fmt.Fprintf(w, "Expecting GET on transform with id %s", id)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/transform", transform)
	http.HandleFunc("/transform/", transformID)
	http.ListenAndServe(":8080", nil)
}

/* SDD Notes

// SDD Homework?: Mediator

/transform                              (GET)
/transform/:id                          (GET)
/transform/add/root                     (POST)
/transform/add/child                    (POST)
/transform/delete/:id                   (DELETE)
/transform/update/:id                   (UPDATE)
/datatype                               (GET)
/datatype/:name                         (GET)
*/
