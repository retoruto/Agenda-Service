package service
import (
	"net/http"
    "Agenda-Service/entity"
	"github.com/unrolled/render"
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
)
func LoginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("login")
		entity.StartAgenda()
		req.ParseForm()
		if (entity.CurrentUser.Name != "") {
			w.WriteHeader(http.StatusForbidden)
		} else {	
			if entity.UserLogIn(req.FormValue("username"), req.FormValue("password")) {
				entity.CurrentUser = entity.FindUserByName(req.FormValue("username"))
				fmt.Print(entity.CurrentUser)
				formatter.JSON(w, http.StatusOK, entity.CurrentUser)
			}
			
		}
	}
}
/*
func LogoutHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("logout")
		entity.StartAgenda()
		req.ParseForm()
		if (entity.CurrentUser.Name != "") {
			w.WriteHeader(http.StatusForbidden)
		} else {	
			if entity.UserLogIn(req.FormValue("username"), req.FormValue("password")) {
				entity.CurrentUser = entity.FindUserByName(req.FormValue("username"))
				fmt.Print(entity.CurrentUser)
				formatter.JSON(w, http.StatusOK, entity.CurrentUser)
			}
			
		}
	}
}
*/
func ListAllUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("ListAllUser")
		entity.StartAgenda()
		if (entity.CurrentUser.Name == "") {
			w.WriteHeader(http.StatusForbidden)
		} else {
			uList, err := entity.FindAllUser()
			fmt.Println(uList)
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
	  entity.StartAgenda()
	  fmt.Println(entity.CurrentUser)
		if (entity.CurrentUser.Name == "") {
			w.WriteHeader(http.StatusForbidden)
		} else {
			if entity.DeleteUser(entity.CurrentUser.Name, entity.CurrentUser.Password) {
				w.WriteHeader(http.StatusOK)
				fmt.Println("delete user successfully!")
			} else {
				w.WriteHeader(http.StatusForbidden)
				fmt.Println("Fail to delete user!")
			}
		}
		
	}
}
func DeleteMeetingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		title := vars["title"]
		if (entity.CurrentUser.Name == "") {
			w.WriteHeader(http.StatusForbidden)
		} else {
			if entity.DeleteMeeting(entity.CurrentUser.Name, title) {
				w.WriteHeader(http.StatusOK)
				fmt.Print("delete meeting successfully!")
			} else {
				w.WriteHeader(http.StatusForbidden)
				fmt.Print("Fail to delete user!")
			}
		}
		
	}
}
func ListAllMeetingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
	  entity.StartAgenda()
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

func CreateMeetingHandler(formatter *render.Render) http.HandlerFunc {

		return func(w http.ResponseWriter, req *http.Request) {
		  entity.StartAgenda()
		  fmt.Println(entity.CurrentUser.Name)
			if (entity.CurrentUser.Name == "") {
				w.WriteHeader(http.StatusForbidden)
			} else {
				decoder := json.NewDecoder(req.Body)
				var m entity.Meeting
				err := decoder.Decode(&m)
				m.Sponsor = entity.CurrentUser.Name
				fmt.Println(err)
				if err != nil {
					panic(err)
				}
				if !entity.CreateMeeting(m.Sponsor, m.Title, m.StartDate, m.EndDate, m.Participators) {
					w.WriteHeader(http.StatusForbidden)
				} else {
					
					formatter.JSON(w, http.StatusCreated, m)
				}		
			}
		}
}
