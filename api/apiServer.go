package api

import (
	"fmt"
	"net/http"
	"github.com/ProtoML/ProtoML-persist/persist"
) 

const (
	APIDatatype                        = "/datatype"
	APITransform                       = "/transform"
	APITransformAddChild               = "/transform/add/child"

	APIDatatypePrefix                  = "/datatype/"
	APITransformPrefix                 = "/transform/"
	APITransformDeletePrefix           = "/transform/delete/"
	APITransformUpdatePrefix           = "/transform/update/"

	APIDatatypePrefixLen               = len(APIDatatypePrefix)
	APITransformPrefixLen              = len(APITransformPrefix)
	APITransformDeletePrefixLen        = len(APITransformDeletePrefix)
	APITransformUpdatePrefixLen        = len(APITransformUpdatePrefix)

	APIDefaultPort                     = 8080
)

type serverState struct {
	store *persist.PersistStorage
}

func (server* serverState) unrecognizedCall(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Call Unrecognized")
}

func (server* serverState) handleDatatype(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expecting GET on all datatypes");
}

func (server* serverState)  handleDatatypePrefix(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[APIDatatypePrefixLen:]
	fmt.Fprintf(w, "Expecting GET on datatype with id %s", id)
}

func (server* serverState) handleTransform(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expecting GET on all transforms")
}

func (server* serverState) handleTransformPrefix(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[APITransformPrefixLen:]
	fmt.Fprintf(w, "Expecting GET on transform with id %s", id)
}

func (server* serverState) handleTransformAddRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expecting POST to add a root")
}

func (server* serverState) handleTransformAddChild(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expecting POST to add a child")
}

func (server* serverState) handleTransformDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[APITransformDeletePrefixLen:]
	fmt.Fprintf(w, "Expecting DELETE on transform with id %s", id)
}

func (server* serverState) handleTransformUpdate(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[APITransformUpdatePrefixLen:]
	fmt.Fprintf(w, "Expecting UPDATE on transform with id %s", id)
}

func APIServer(port int, store *persist.PersistStorage) (err error) {
	apiServerState := &serverState{store}

	http.HandleFunc("/", apiServerState.unrecognizedCall)
	http.HandleFunc(APIDatatype,                   apiServerState.handleDatatype)            // GET
	http.HandleFunc(APIDatatypePrefix,             apiServerState.handleDatatypePrefix)      // GET
	http.HandleFunc(APITransform,                  apiServerState.handleTransform)           // GET
	http.HandleFunc(APITransformPrefix,            apiServerState.handleTransformPrefix)     // GET
	http.HandleFunc(APITransformAddChild,          apiServerState.handleTransformAddChild)   // POST
	http.HandleFunc(APITransformDeletePrefix,      apiServerState.handleTransformDelete)     // DELETE
	http.HandleFunc(APITransformUpdatePrefix,      apiServerState.handleTransformUpdate)     // UPDATE
	err = http.ListenAndServe(":"+string(port), nil)
	return
} 

