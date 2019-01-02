package server

import (
	"net/http"
	"github.com/qiongtubao/latte_go_lib"
)
type Server struct {
	Port int
	Gets map[string]func(http.ResponseWriter, *http.Request)
	Posts map[string]func(http.ResponseWriter, *http.Request)
	mux *http.ServeMux
}
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var method func(http.ResponseWriter, *http.Request)
	if r.Method == "GET" {
		method = s.Gets[r.URL.Path]
	} else {
		method = s.Posts[r.URL.Path]
	}
	if method != nil {
		method(w, r)
	} else {
		http.NotFound(w, r)
	}
}
func (s Server) Get(name string, method func(http.ResponseWriter, *http.Request)) {
    s.Gets[name] = method                           
}
func (s Server) Post(name string, method func(http.ResponseWriter, *http.Request)) {
	s.Posts[name] = method
}
func (s Server) Start() {
	http.ListenAndServe(":" + lib.IntToStr(s.Port), s)
}
func CreateServer() Server {
	gets := map[string]func(http.ResponseWriter, *http.Request){}
	posts := map[string]func(http.ResponseWriter, *http.Request){}
	server := Server{3000, gets, posts, http.NewServeMux()}
	return server
}