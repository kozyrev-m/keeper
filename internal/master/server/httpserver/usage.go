package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
)

// respond - helper for easy rendering of the response.
func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)

	if data == nil {
		data = map[string]string{}
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("can't encode data to json: %s", err.Error())
	}
}

// error use to simplify rendering of the response (based on 'func (s *server) respond(wr) {...}').
func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}
