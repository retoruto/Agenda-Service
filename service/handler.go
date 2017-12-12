package service
import (
	"net/http"
    "Agenda-Service/entity"
	"github.com/unrolled/render"
	"fmt"
	"encoding/json"
)
func ListAllUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if (entity.CurrentUser.Name == "") {
			w.WriteHeader(http.StatusForbidden)
		} else {
			uList, err := entity.FindAllUser()
			if err != nil {
				panic(err)
			}
			formatter.JSON(w, http.StatusOK, uList)
		}
	}
}
func UserRegisterHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var u entity.User
		err := decoder.Decode(&u)
		if err != nil {
			panic(err)
		}

		if !entity.UserRegister(u.Name, u.Password, u.Email, u.Phone) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		formatter.JSON(w, http.StatusCreated, u)	
	}
}

func DeleteUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if (entity.CurrentUser.Name == "") {
			w.WriteHeader(http.StatusForbidden)
		} else {
			if entity.DeleteUser(entity.CurrentUser.Name, entity.CurrentUser.Password) {
				w.WriteHeader(http.StatusNoContent)
				fmt.Print("delete user successfully!")
			} else {
				w.WriteHeader(http.StatusForbidden)
				fmt.Print("Fail to delete user!")
			}
		}
		
	}
}

func ListAllMeetingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if (entity.CurrentUser.Name == "") {
			w.WriteHeader(http.StatusForbidden)
		} else {
			uMeeting, err := entity.FindAllMeeting()
			if err != nil {
				panic(err)
			}
			formatter.JSON(w, http.StatusOK, uMeeting)
		}
	}
}