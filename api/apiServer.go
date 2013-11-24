package api

import (
	"fmt"
	"net/http"
	//"github.com/ProtoML/ProtoML-persist/persist/elastic"
	"github.com/ProtoML/ProtoML/logger"
	"github.com/ProtoML/ProtoML-persist/persist"
	"html"
	"github.com/ant0ine/go-json-rest"
) 

const (
	LOGTAG                             = "API-Server"
	DEFAULT_API_PORT                   = 8080
	SERVER_CLOSE_PANIC_ERROR           = "Closing webserver, ignore panic"
)

type APIServerState struct {
	Port int
	poisonPill chan bool
	errChan chan error
	Store persist.PersistStorage
	ResourceHandler rest.ResourceHandler
}

func New(port int, store persist.PersistStorage)(*APIServerState) {
	return &APIServerState{Port:port,Store:store,ResourceHandler:rest.ResourceHandler{}}
}

// Used to halt server due to non-trivial gracefull server stopping
func (server *APIServerState) Close() {
	server.poisonPill <- true // pop poison pill to panic then recover
}

// Start api server with ability to shutdown
func (server *APIServerState) Start() (errChan chan error) {
	routes := []rest.Route{rest.Route{"GET", "/", server.index}}
	routes = append(routes, server.APIHandleFuncs()...)
	server.ResourceHandler.SetRoutes(routes...)

	server.errChan = make(chan error)
	// main server
	HTTP := func() {
		logger.LogInfo(LOGTAG,"Starting API HTTP Server On Port :%d",server.Port)
		err := http.ListenAndServe(fmt.Sprintf(":%d",server.Port), &server.ResourceHandler)
		server.errChan <- err
	}
	// closing halt bomb
	server.poisonPill = make(chan bool)
	Panic := func(PoisonPill chan bool) {
		<-PoisonPill 
		panic(SERVER_CLOSE_PANIC_ERROR)
	}
	// stopping bomb and shutdown api server
	Recover := func(){
		recover()
		close(server.errChan)
	}
	// run in goroutine to allow for return
	go func() {
		defer server.Store.Close()
		defer Recover()
		go HTTP()
		Panic(server.poisonPill)
	}()
	
	return server.errChan
} 

func (server* APIServerState) unrecognizedCall(w *rest.ResponseWriter, req *rest.Request) {
	//http.Redirect(w, req, "/error", http.StatusNotFound)
	return
}

func (server* APIServerState) index(w *rest.ResponseWriter, req *rest.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))
}

