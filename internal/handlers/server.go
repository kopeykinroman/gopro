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
	return http.ListenAndServe(`:8080`, s.handlers.router)
}
