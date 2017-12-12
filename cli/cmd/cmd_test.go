package cmd

import (
	"fmt"
	"os"
	"testing"
)

func TestRegister(t *testing.T) {
	fmt.Println("=====> In TEST of Register")
	registerCmd.Flags().Set("username", "r0beRT")
	registerCmd.Flags().Set("password", "passw0rd")
	registerCmd.Flags().Set("email", "dr.paper@live.com")
	registerCmd.Flags().Set("phone", "17665310114")
	registerCmd.Run(registerCmd, nil)
}

func TestLogin(t *testing.T) {
	fmt.Println("=====> In TEST of Login")
	logoutCmd.Run(logoutCmd, nil)
	loginCmd.Flags().Set("username", "r0beRT")
	loginCmd.Flags().Set("password", "passw0rd")
	loginCmd.Run(loginCmd, nil)
}

func TestShowAllUsers(t *testing.T) {
	fmt.Println("=====> In TEST of Showing all users")
	usersCmd.Run(usersCmd, nil)
}

func TestCreateNewMeeting(t *testing.T) {
	fmt.Println("=====> In TEST of Creating a new meeting")
	createmeetingCmd.Flags().Set("title", "testMeeting")
	createmeetingCmd.Flags().Set("members", "testUser0,testUser1")
	createmeetingCmd.Flags().Set("starttime", "2017/12/25/13:00")
	createmeetingCmd.Flags().Set("endtime", "2017/12/25/17:00")
	createmeetingCmd.Run(createmeetingCmd, nil)
}

func TestShowAllMeetings(t *testing.T) {
	fmt.Println("=====> In TEST of Showing all meetings")
	meetingsCmd.Run(meetingsCmd, nil)
}

func TestDestroy(t *testing.T) {
	fmt.Println("=====> In TEST of destroying account")
	destroyCmd.Run(destroyCmd, nil)
	usersCmd.Run(usersCmd, nil)
	os.Remove(".key")
}
