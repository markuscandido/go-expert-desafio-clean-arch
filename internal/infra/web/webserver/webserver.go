package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebHandler struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Handlers      []WebHandler
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      []WebHandler{},
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, WebHandler{
		Method:  method,
		Path:    path,
		Handler: handler,
	})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, h := range s.Handlers {
		switch h.Method {
		case http.MethodGet:
			s.Router.Get(h.Path, h.Handler)
		case http.MethodPost:
			s.Router.Post(h.Path, h.Handler)
		case http.MethodPut:
			s.Router.Put(h.Path, h.Handler)
		case http.MethodDelete:
			s.Router.Delete(h.Path, h.Handler)
		case http.MethodPatch:
			s.Router.Patch(h.Path, h.Handler)
		default:
			s.Router.Handle(h.Path, h.Handler)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
