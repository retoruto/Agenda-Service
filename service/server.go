package service

import (
	"net/http"
  "fmt"
  "Agenda-Service/entity"
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
	//mx.HandleFunc("/hello/{id}", testHandler(formatter)).Methods("GET")
	mx.HandleFunc("/service/userinfo", postUserInfoHandler(formatter)).Methods("POST")
	mx.HandleFunc("/service/userinfo", getUserInfoHandler(formatter)).Methods("GET")
	//mx.HandleFunc("/v1/users", testHandler(formatter)).Methods("GET")
}

/*func testHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		//id := vars["id"]
		//formatter.JSON(w, http.StatusOK, entity.FindAllUser())
	}
}*/
func getUserInfoHandler(formatter *render.Render) http.HandlerFunc {
   fmt.Println("get")
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
	/*	if len(req.Form["userid"][0]) != 0 {
			i, _ := strconv.ParseInt(req.Form["userid"][0], 10, 32)

			u := entities.UserInfoService.FindByID(int(i))
			formatter.JSON(w, http.StatusBadRequest, u)
			return
		}*/
		ulist, _ := entity.FindAllUser()
		fmt.Println(ulist)
		formatter.JSON(w, http.StatusOK, ulist)
	}
}
func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {
   fmt.Println("post")
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if len(req.Form["username"][0]) == 0 {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
			return
		}
		u := entity.User{Name: req.Form["username"][0]}
		u.Password = req.Form["password"][0]
		u.Email = req.Form["email"][0]
		u.Phone = req.Form["phone"][0]
		//entity.UserRegister(u.Name, u.Password, u.Email, u.Phone)
		formatter.JSON(w, http.StatusOK, u)
	}
}
