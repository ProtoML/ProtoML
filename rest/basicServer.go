package main

import (
	"fmt"
	"net/http"
)

func unrecognizedCall(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Call Unrecognized")
}

func datatype(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expecting GET on all datatypes");
}

func datatypeID(w http.ResponseWriter, r *http.Request) {
	const pathLen = len("/datatype/");
	id := r.URL.Path[pathLen:]
	fmt.Fprintf(w, "Expecting GET on datatype with id %s", id)
}

func transform(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expecting GET on all transforms")
}

func transformID(w http.ResponseWriter, r *http.Request) {
	const pathLen = len("/transform/")
	id := r.URL.Path[pathLen:]
	fmt.Fprintf(w, "Expecting GET on transform with id %s", id)
}

func transformAddRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expecting POST to add a root")
}

func transformAddChild(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expecting POST to add a child")
}

func transformDelete(w http.ResponseWriter, r *http.Request) {
	const pathLen = len("/transform/delete/")
	id := r.URL.Path[pathLen:]
	fmt.Fprintf(w, "Expecting DELETE on transform with id %s", id)
}

func transformUpdate(w http.ResponseWriter, r *http.Request) {
	const pathLen = len("/transform/update/")
	id := r.URL.Path[pathLen:]
	fmt.Fprintf(w, "Expecting UPDATE on transform with id %s", id)
}

func main() {
	http.HandleFunc("/", unrecognizedCall)
	http.HandleFunc("/datatype",                datatype)            // GET
	http.HandleFunc("/datatype/",               datatypeID)          // GET
	http.HandleFunc("/transform",               transform)           // GET
	http.HandleFunc("/transform/",              transformID)         // GET
	http.HandleFunc("/transform/add/root",      transformAddRoot)    // POST
	http.HandleFunc("/transform/add/child",     transformAddChild)   // POST
	http.HandleFunc("/transform/delete/",       transformDelete)     // DELETE
	http.HandleFunc("/transform/update/",       transformUpdate)     // UPDATE
	http.ListenAndServe(":8080", nil)
}
