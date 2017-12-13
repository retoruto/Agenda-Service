package entity
import (
    "testing"
    "reflect"
    "fmt"
)

func init() {
  userlist = nil
  meetinglist = nil
}
func TestCreateUser_DB(t *testing.T) {
    userlist, _ = FindAllUser()
    cases := []struct {
        in User
        want []User
    }{
        {t_users[0], t_users[:1]},
        {t_users[1], t_users[:2]},
        {t_users[2], t_users[:3]},
    }
    for _, c := range cases {
        CreateUser_DB(&c.in)
        userlist, _ = FindAllUser()
        fmt.Printf("CreateUser_DB userlist: %v\n", userlist)
        if got := userlist; !reflect.DeepEqual(got, c.want)  {
            t.Errorf("CreateUser_DB(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}

func TestCreateMeeting_DB(t *testing.T) {
    meetinglist, _ = FindAllMeeting()
    cases := []struct {
        in Meeting
        want []Meeting
    }{
        {t_meetings[0], t_meetings[:1]},
        {t_meetings[1], t_meetings[:2]},
        {t_meetings[2], t_meetings[:3]},
    }
    for _, c := range cases {
        CreateMeeting_DB(&c.in)
        meetinglist, _ := FindAllMeeting()
        fmt.Printf("CreateMeeting_DB meetinglist: %v\n", meetinglist)
        if got := meetinglist; !reflect.DeepEqual(got, c.want)  {
            t.Errorf("CreateMeeting(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
/*
func TestQueryMeeting(t *testing.T) {
    cases := []struct {
        in mFilter
        want []Meeting
    }{
        {func (m *Meeting) bool { return m.getSponsor() == "a"}, []Meeting{t_meetings[0]}},
        {func (m *Meeting) bool { return m.getSponsor() == "b"}, []Meeting{t_meetings[1]}},
        {func (m *Meeting) bool { return m.getSponsor() == "c"}, []Meeting{t_meetings[2]}},
    }
    for _, c := range cases {
        got := queryMeeting(c.in)
        if !reflect.DeepEqual(got, c.want)  {
            t.Errorf("queryMeeting(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
func TestQueryUser(t *testing.T) {
    cases := []struct {
        in uFilter
        want []User
    }{
        {func (m *User) bool { return m.getName() == "a"}, []User{t_users[0]}},
        {func (m *User) bool { return m.getName() == "b"}, []User{t_users[1]}},
        {func (m *User) bool { return m.getName() == "c"}, []User{t_users[2]}},
    }
    for _, c := range cases {
        got := queryUser(c.in)
        if !reflect.DeepEqual(got, c.want)  {
            t.Errorf("queryMeeting(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
*/
/*
func TestUpdateUser(t *testing.T) {
    cases := []struct {
        inf uFilter
        ins uSwitcher
        want []User
    }{
        {func (m *User) bool { return m.getName() == "a"}, 
        func (m *User) {m.setName("aaa")}, t_users[1:4]},
        
        {func (m *User) bool { return m.getName() == "b"},
        func (m *User) {m.setName("bbb")}, t_users[2:5]},
        {func (m *User) bool { return m.getName() == "c"},
        func (m *User) {m.setName("ccc")}, t_users[3:6]},
    }
    for _, c := range cases {
        i := updateUser(c.inf, c.ins)
        //userlist[i].setName("aaaa")
        fmt.Printf("userlist: %v %d\n", userlist, i)
        if got := userlist; !reflect.DeepEqual(got, c.want)  {
            t.Errorf("UpdateUser(%q) == %q, want %q", c.inf, got, c.want)
        }
    }
}*/

func TestDeleteUser_DB(t *testing.T) {
    userlist, _ = FindAllUser()
    cases := []struct {
        in User
        want []User
    }{
        {t_users[0], t_users[1:3]},
        {t_users[1], t_users[2:3]},
        {t_users[2], []User{}},
    }
    for _, c := range cases {
        DeleteUser_DB(&c.in)
        userlist, _ = FindAllUser()
        fmt.Printf("DeleteUser_DB userlist: %v  want: %v\n", userlist, c.want)
        if got := userlist; !reflect.DeepEqual(got, c.want)  {
            t.Errorf("DeleteUser_DB(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}


func TestDeleteMeeting_DB(t *testing.T) {
    meetinglist, _ = FindAllMeeting()
    cases := []struct {
        in Meeting
        want []Meeting
    }{
        {t_meetings[0], t_meetings[1:3]},
        {t_meetings[1], t_meetings[2:3]},
        {t_meetings[2], t_meetings[3:3]},
    }
    for _, c := range cases {
        DeleteMeeting_DB(&c.in)
        meetinglist, _ := FindAllMeeting()
        fmt.Printf("DeleteMeeting_DB meetinglist: %v\n", meetinglist)
        if got := meetinglist; !reflect.DeepEqual(got, c.want)  {
            t.Errorf("DeleteMeeting_DB(%q) == %q, want %q", c.in, got, c.want)
        }
    }

}