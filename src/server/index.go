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
func (s Server) AddHandle(name string, method func(http.ResponseWriter, *http.Request)) {
    s.Methods[name] = method                           
}
func (s Server) Start() {
	for name := range s.Methods {
		// fmt.Println(country, "首都是", s.mux [ name ])
		s.mux.HandleFunc(name, s.Methods[name])
	}
	http.ListenAndServe(":" + lib.IntToStr(s.Port), s.mux)
}
func CreateServer() Server {
	m := map[string]func(http.ResponseWriter, *http.Request){}
	server := Server{3000, m, http.NewServeMux()}
	return server
}