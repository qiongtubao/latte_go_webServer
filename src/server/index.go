package server

import (
	"net/http"
	"github.com/qiongtubao/latte_go_lib"
)
type Server struct {
	Port int
	Gets map[string][]func(http.ResponseWriter, *http.Request) bool
	Posts map[string][]func(http.ResponseWriter, *http.Request) bool
	mux *http.ServeMux
}
func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var methods []func(http.ResponseWriter, *http.Request) bool
	if r.Method == "GET" {
		methods = s.Gets[r.URL.Path]
	} else {
		methods = s.Posts[r.URL.Path]
	}
	if methods != nil {
		arrLen := len(methods)
		var i int
		var result bool
		for i = 0; i < arrLen; i++ {
			result = methods[i](w, r)
			if result == false {
				return 
			}
		}
	} else {
		http.NotFound(w, r)
	}
}
func (s Server) Get(name string, method ...func(http.ResponseWriter, *http.Request) bool) {
    s.Gets[name] = method                           
}
func (s Server) Post(name string, method ...func(http.ResponseWriter, *http.Request) bool ) {
	s.Posts[name] = method
}
func (s Server) Start() {
	http.ListenAndServe(":" + lib.IntToStr(s.Port), s)
}
func CreateServer() Server {
	gets := map[string][]func(http.ResponseWriter, *http.Request) bool {}
	posts := map[string][]func(http.ResponseWriter, *http.Request) bool {}
	server := Server{3000, gets, posts, http.NewServeMux()}
	return server
}