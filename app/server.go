package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	handler *Handler
}

func NewServer(handler *Handler) *Server {
	return &Server{
		handler: handler,
	}
}
func (s *Server) Start() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods("GET").HandlerFunc(s.handler.HandleGetAllTask)
	router.Path("/tasks").Methods("GET").Queries("completed", "false").HandlerFunc(s.handler.HandleGetUncomleteTask)
	router.Path("/tasks").Methods("POST").HandlerFunc(s.handler.HandleCreateTask)
	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.handler.HandleGetTask)
	router.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.handler.HandleComleteTask)
	router.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(s.handler.HandleDeleteTask)

	return http.ListenAndServe(":9091", router)
}
