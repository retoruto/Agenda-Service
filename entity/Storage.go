package entity

import (
	"os"
	"io"
	"fmt"
	"encoding/json"
)


var userlist []User
var meetinglist []Meeting
var CurrentUser User

type uFilter func (*User) bool
type uSwitcher func (*User) 
type mFilter func (*Meeting) bool
type mSwitcher func (*Meeting) 

// getJson
func GetJson(v interface{}) string {
	a, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(a)
}

// ReadFromDb .
func ReadFromDb() {
  var err1, err2 error
	userlist, err1 = FindAllUser()
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Fail to read users from sqlite3")
	}
	meetinglist, err2 = FindAllMeeting()
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "Fail to read meetings from sqlite3")
	}
	
}

func ReadCurrentUser()  {
	file1, err1 := os.Open("CurUser")
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Fail to open CurUser")
	}
	dec1 := json.NewDecoder(file1)
	err1 = dec1.Decode(&CurrentUser)	
	if err1 != io.EOF && err1 != nil {
		fmt.Fprintf(os.Stderr, "Fail to Decode")
	} 
	file1.Close()
}


func writeCurrentUser()  {
	file1, err1 := os.Create("CurUser");
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Fail to create CurUser")		
	}
	enc1 := json.NewEncoder(file1)	
	if err1 := enc1.Encode(&CurrentUser); err1 != nil {
		fmt.Fprintf(os.Stderr, "Fail to encode")
	}
	file1.Close()
}
func createUser(t_user User) {
	userlist = append(userlist,t_user)
	err := CreateUser_DB(&t_user)
	if err != nil {
		fmt.Print(err)
		fmt.Fprintf(os.Stderr, "Fail to create User")		
	}
}
func createMeeting(t_meeting Meeting) {
	meetinglist = append(meetinglist,t_meeting)
	err := CreateMeeting_DB(&t_meeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail to create Meeting")		
	}
}

func queryUser(filter uFilter) []User {
	var dy []User;
	for _, u := range userlist {
		if filter(&u) {
			dy = append(dy, u)
		}
	}
	return dy
}
/*
func updateUser(filter uFilter, switcher uSwitcher) int {
	n := 0
	for _, u := range userlist {
		if filter(&u) {
			switcher(&u)
			n++
		}
	}
	return n
}
*/

func deleteUser(filter uFilter) int {
	n := 0
	for i, u := range userlist {
		if filter(&u) {
		  err := DeleteUser_DB(&u)
		  if err != nil {
		    fmt.Fprintf(os.Stderr, "Fail to delete user")		
    	}
			userlist[i] = userlist[len(userlist) - 1 - n]			
			n++

		}
	}
	
	userlist = userlist[:len(userlist)  - n]
	return n
}
func queryMeeting(filter mFilter) []Meeting {
	var dy []Meeting;
	for _, m := range meetinglist {
		if filter(&m) {
			dy = append(dy, m)
		}
	}
	return dy
}

/*
func updateMeeting(filter mFilter, switcher mSwitcher) int {
	n := 0
	for _, m := range meetinglist {
		if filter(&m) {
			switcher(&m)
			n++
		}
	}
	return n
}
*/
func deleteMeeting(filter mFilter) int {
	n := 0
	for i, m := range meetinglist {
		if filter(&m) {
		  err := DeleteMeeting_DB(&m)
		  if err != nil {
		    fmt.Fprintf(os.Stderr, "Fail to delete Meeting")		
    	}
			meetinglist[i] = meetinglist[len(meetinglist) - 1 - n]
			n++
		}
	}
		
	meetinglist = meetinglist[:len(meetinglist) - n]
	return n
}



