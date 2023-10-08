package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-web/constants"
	"github.com/go-web/endpoint"
	"github.com/go-web/middleware"
	"github.com/go-web/pkg/model"
)

type Server struct {
	endpoints map[string]endpoint.Handler
}

func NewServer() Server {
	server := Server{}
	server.endpoints = make(map[string]endpoint.Handler, 0)
	return server
}

func (s *Server) Start() {

	handler := middleware.AuthMiddleware(s.Handle())

	http.Handle("/", handler)
}

func (s *Server) AddHandler(e endpoint.Endpoint) {

	key := s.CreateKey(e.Path, e.Method)

	s.endpoints[key] = e.Handler
}

func (s *Server) CreateKey(path string, method string) string {
	return fmt.Sprintf("%s~%s", path, method)
}

func (s *Server) Handle() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		log.Println("Path : ", r.URL.Path)
		log.Println("Method : ", r.Method)

		// TODO
		// Handle cors
		if r.Method == "OPTIONS" {
			b, _ := json.Marshal(model.Error{
				Message: "Invalid route",
				Status:  200,
			})
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			rw.Header().Set("Access-Control-Allow-Methods", "*")
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(b))
			return
		}

		key := s.CreateKey(r.URL.Path, r.Method)

		handler, ok := s.endpoints[key]
		log.Println("ok : ", ok)
		if ok {
			response, err := handler(r)

			if err != nil {
				s.WriteError(err, rw)
				return
			}

			response.Write(rw)
		} else {
			b, _ := json.Marshal(model.Error{
				Message: "Invalid route",
				Status:  404,
			})
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte(b))
		}
	})
}

func (s *Server) WriteError(err error, rw http.ResponseWriter) {
	e := err.(model.Error)

	b, _ := json.Marshal(model.Error{
		Message: e.Message,
		Status:  e.Status,
	})

	status := http.StatusInternalServerError

	if e.Status != 0 {
		status = e.Status
	}

	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	rw.Header().Set("Access-Control-Allow-Methods", "*")
	rw.Header().Set(constants.HEADER_CONTENT_TYPE_KEY, constants.HEADER_APPLICATION_JSON)
	rw.WriteHeader(status)
	rw.Write([]byte(b))
}
