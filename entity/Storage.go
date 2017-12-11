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

func ReadCurrentUser()  {
	file1, err1 := os.Open("CurUser");
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

func ReadFromFile() {
	//读user
	file1, err1 := os.Open("UserInfo");
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Fail to open UserInfo")
	}
	dec1 := json.NewDecoder(file1)
	err1 = dec1.Decode(&userlist)	
	if err1 != io.EOF && err1 != nil {
		fmt.Fprintf(os.Stderr, "Fail to Decode")
	}
	file1.Close()

	//读Meeting
	file2, err2 := os.Open("MeetingInfo");
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "Fail to open MeetingInfo")
	}
	dec2 := json.NewDecoder(file2)
	err2 = dec2.Decode(&meetinglist)	
	if err2 != io.EOF && err2 != nil{
		fmt.Fprintf(os.Stderr, "Fail to Decode")
	} 
	file2.Close()
}

func writeToFile()  {
	//写User
	file1, err1 := os.Create("UserInfo");
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Fail to create UserInfo")		
	}
	enc1 := json.NewEncoder(file1)	
	if err1 := enc1.Encode(&userlist); err1 != nil {
		fmt.Fprintf(os.Stderr, "Fail to encode")
	}
	file1.Close()
	//写Meeting
	file2, err2 := os.Create("MeetingInfo");
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "Fail to create MeetingInfo")		
	}
	enc2 := json.NewEncoder(file2)	
	if err2 := enc2.Encode(&meetinglist); err2 != nil {
		fmt.Fprintf(os.Stderr, "Fail to encode")
	}
	file2.Close()
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
	
}
func createMeeting(t_meeting Meeting) {
	meetinglist = append(meetinglist,t_meeting)
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

func deleteUser(filter uFilter) int {
	n := 0
	for i, u := range userlist {
		if filter(&u) {
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

func deleteMeeting(filter mFilter) int {
	n := 0
	for i, m := range meetinglist {
		if filter(&m) {
			meetinglist[i] = meetinglist[len(meetinglist) - 1 - n]
			n++
		}
	}
	meetinglist = meetinglist[:len(meetinglist) - n]
	return n
}



