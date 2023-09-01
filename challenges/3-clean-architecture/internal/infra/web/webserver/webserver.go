package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServerHandler struct {
	Handler http.HandlerFunc
	Method  string
}
type WebServer struct {
	Router        chi.Router
	Handlers      map[string][]WebServerHandler
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string][]WebServerHandler),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path, method string, handler http.HandlerFunc) {
	if _, ok := s.Handlers[path]; !ok {
		s.Handlers[path] = make([]WebServerHandler, 0)
	}
	s.Handlers[path] = append(s.Handlers[path], WebServerHandler{
		Handler: handler,
		Method:  method,
	})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, config := range s.Handlers {
		for _, handler := range config {
			s.Router.Method(handler.Method, path, handler.Handler)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
