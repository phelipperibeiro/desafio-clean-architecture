package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (webServer *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	webServer.Handlers[path] = handler
}

func (webServer *WebServer) Start() {

	webServer.Router.Use(middleware.Logger)

	for path, handler := range webServer.Handlers {
		webServer.Router.Handle(path, handler)
	}

	http.ListenAndServe(webServer.WebServerPort, webServer.Router)
}
