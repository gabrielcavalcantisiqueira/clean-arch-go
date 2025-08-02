package webserver

import (
	"net/http"
	"strings"

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

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	key := method + ":" + path
	s.Handlers[key] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for key, handler := range s.Handlers {
		parts := strings.SplitN(key, ":", 2)
		if len(parts) != 2 {
			continue // key inválida, pula
		}
		method := parts[0]
		path := parts[1]
		s.Router.MethodFunc(method, path, handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
