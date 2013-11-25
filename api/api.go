package api

import (
	"fmt"
	//"github.com/ProtoML/ProtoML/logger"
	"github.com/ant0ine/go-json-rest"
	"github.com/ProtoML/ProtoML-persist/persist/persistparsers"
	"github.com/ProtoML/ProtoML/types"
	"net/http"
)

const (
	APILOGTAG = "API"
	GRAPHROOT = "/graph"
	TRANSFORMROOT = "/transform"
	DATASETROOT = "/dataset"
)

type success struct {
	Sucess string
}

func (server *APIServerState) APIHandleFuncs() (routes []rest.Route) {
	routes = append(routes,
		rest.Route{"GET", GRAPHROOT, server.APIHandleGetGraph},
		rest.Route{"POST", TRANSFORMROOT, server.APIHandleNewTransform},
		rest.Route{"PUT", TRANSFORMROOT+"/:id", server.APIHandleUpdateTransform},
		rest.Route{"POST", DATASETROOT, server.APIHandleNewDataset},
	)
	return
}

func (server *APIServerState) APIHandleGetGraph(w *rest.ResponseWriter, req *rest.Request) {
	graph, err := server.Store.GetGraph()
	if err != nil {
		rest.Error(w, fmt.Sprintf("Error retrieving graph: %s", err), http.StatusBadRequest)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteJson(graph)
	return
}


func (server *APIServerState) APIHandleNewTransform(w *rest.ResponseWriter, req *rest.Request) {
	var itransform types.InducedTransform
	// decode request
	err := req.DecodeJsonPayload(&itransform)
	if err != nil {
		rest.Error(w, fmt.Sprintf("Could not parse input json: %s", err), http.StatusBadRequest)
		return
	}
	// validated induced transform
	err = persistparsers.ValidateInducedTransform(itransform)
	if err != nil {
		rest.Error(w, fmt.Sprintf("Induced transform could not be validated: %s", err), http.StatusBadRequest)
		return
	} 
	// add it to persist storage
	id, err := server.Store.AddInducedTransform(itransform)
	if err != nil {
		rest.Error(w, fmt.Sprintf("Could not add induced transform: %s", err), http.StatusBadRequest)
		return
	}
	s := success{id}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteJson(s)
	return
}

type itransformUpdate struct {
	Id string
	Itransform types.InducedTransform
}

func (server *APIServerState) APIHandleUpdateTransform(w *rest.ResponseWriter, req *rest.Request) {
	var itu itransformUpdate
	// decode request
	err := req.DecodeJsonPayload(&itu)
	if err != nil {
		rest.Error(w, fmt.Sprintf("Could not parse input json: %s", err), http.StatusBadRequest)
		return
	}
	// validated induced transform
	err = persistparsers.ValidateInducedTransform(itu.Itransform)
	if err != nil {
		rest.Error(w, fmt.Sprintf("Induced transform could not be validated: %s", err), http.StatusBadRequest)
		return
	} 
	// update the induced transform in the persist storage
	err = server.Store.UpdateInducedTransform(itu.Id, itu.Itransform)
	if err != nil {
		rest.Error(w, fmt.Sprintf("Could not add induced transform: %s", err), http.StatusBadRequest)
		return
	}
	s := success{itu.Id}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteJson(s)
	return
}

func (server *APIServerState) APIHandleNewDataset(w *rest.ResponseWriter, req *rest.Request) {
	return
}
