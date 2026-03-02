package handlers

import "net/http"

type HTTPServer struct {
	handlers *HTTPHandlers
	addres   string
}

func NewHTTPServer(handlers *HTTPHandlers, addres string) *HTTPServer {
	return &HTTPServer{
		handlers: handlers,
		addres:   addres,
	}
}

func (s *HTTPServer) StartServer() error {
	return http.ListenAndServe(s.addres, s.handlers.router)
}
