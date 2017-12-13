package service

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/v1/login", LoginHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/key", UserRegisterHandler(formatter)).Methods("POST")
	mx.HandleFunc("/v1/ListAllUser", ListAllUserHandler(formatter)).Methods("GET")
	mx.HandleFunc("/v1/ListAllUser", ListAllUserHandler(formatter)).Methods("GET")
	//mx.HandleFunc("/service/userinfo/{username}", DeleteUserHandler(formatter)).Methods("DELETE")
	//mx.HandleFunc("/v1/users", testHandler(formatter)).Methods("GET")
}


