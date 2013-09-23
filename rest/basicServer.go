package main

import (
	"fmt"
	"net/http"
)

const datatype                        = "/datatype"
const transform                       = "/transform"
const transformAddRoot                = "/transform/add/root"
const transformAddChild               = "/transform/add/child"

const datatypePrefix                  = "/datatype/"
const transformPrefix                 = "/transform/"
const transformDeletePrefix           = "/transform/delete/"
const transformUpdatePrefix           = "/transform/update/"

const datatypePrefixLen               = len(datatypePrefix)
const transformPrefixLen              = len(transformPrefix)
const transformDeletePrefixLen        = len(transformDeletePrefix)
const transformUpdatePrefixLen        = len(transformUpdatePrefix)

func unrecognizedCall(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Call Unrecognized")
}

func handleDatatype(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expecting GET on all datatypes");
}

func handleDatatypePrefix(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[datatypePrefixLen:]
	fmt.Fprintf(w, "Expecting GET on datatype with id %s", id)
}

func handleTransform(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expecting GET on all transforms")
}

func handleTransformPrefix(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[transformPrefixLen:]
	fmt.Fprintf(w, "Expecting GET on transform with id %s", id)
}

func handleTransformAddRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expecting POST to add a root")
}

func handleTransformAddChild(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expecting POST to add a child")
}

func handleTransformDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[transformDeletePrefixLen:]
	fmt.Fprintf(w, "Expecting DELETE on transform with id %s", id)
}

func handleTransformUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[transformUpdatePrefixLen:]
	fmt.Fprintf(w, "Expecting UPDATE on transform with id %s", id)
}

func main() {
	http.HandleFunc("/", unrecognizedCall)
	http.HandleFunc(datatype,                   handleDatatype)            // GET
	http.HandleFunc(datatypePrefix,             handleDatatypePrefix)      // GET
	http.HandleFunc(transform,                  handleTransform)           // GET
	http.HandleFunc(transformPrefix,            handleTransformPrefix)     // GET
	http.HandleFunc(transformAddRoot,           handleTransformAddRoot)    // POST
	http.HandleFunc(transformAddChild,          handleTransformAddChild)   // POST
	http.HandleFunc(transformDeletePrefix,      handleTransformDelete)     // DELETE
	http.HandleFunc(transformUpdatePrefix,      handleTransformUpdate)     // UPDATE
	http.ListenAndServe(":8080", nil)
}
