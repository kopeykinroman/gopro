package handlers

import "net/http"

type HTTPServer struct {
	handlers *HTTPHandlers
}

func NewHTTPServer(handlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		handlers: handlers,
	}
}

func (s *HTTPServer) StartServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, s.handlers.HandleRouteMethod)

	return http.ListenAndServe(`:8080`, mux)
}
