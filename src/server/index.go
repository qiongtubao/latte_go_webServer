package server

import (
	"net/http"
	"github.com/qiongtubao/latte_go_lib"
)
type Server struct {
	Port int
	Methods map[string]func(http.ResponseWriter, *http.Request)
	mux *http.ServeMux
}
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := s.Methods[r.URL.Path]
	if method != nil {
		method(w, r)
	} else {
		http.NotFound(w, r)
	}
}
func (s Server) AddHandle(name string, method func(http.ResponseWriter, *http.Request)) {
    s.Methods[name] = method                           
}
func (s Server) Start() {
	http.ListenAndServe(":" + lib.IntToStr(s.Port), s)
}
func CreateServer() Server {
	m := map[string]func(http.ResponseWriter, *http.Request){}
	server := Server{3000, m, http.NewServeMux()}
	return server
}