package service
import (
	"net/http"
    "Agenda-Service/entity"
	"github.com/unrolled/render"
	"fmt"
	"io/ioutil"
	"encoding/json"
)
func ListAllUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		ulist, _ := entity.FindAllUser()
		fmt.Println(ulist)
		
		formatter.JSON(w, http.StatusOK, ulist)
	}
}
func UserRegisterHandler(formatter *render.Render) http.HandlerFunc {
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
		entity.UserRegister(u.Name, u.Password, u.Email, u.Phone)
		formatter.JSON(w, http.StatusOK, u)
	}
}

func DeleteUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		fmt.Println(232131)
		/*
		result, _ := ioutil.ReadAll(req.Body)
		req.Body.Close()

		u := entity.User{}
		err := json.Unmarshal(result, &u)
		if err != nil {
			panic(err)
		}
		*/
		if !entity.DeleteUser(u.Name, u.Password) {
			formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"delete failed!!!"})
			return
		}
		formatter.JSON(w, http.StatusNoContent, struct{ info string }{"Deleted"})
	}
}
