package entity
import (
	"encoding/json"

)
type UserTable struct {
	UserName string `xorm:"pk varchar(255) notnull "`
	Password string `xorm:"varchar(255) notnull"`
	Email    string `xorm:"varchar(255) notnull"`
	Phone    string `xorm:"varchar(255) notnull"`
}

type MeetingTable struct {
	Title     string     `xorm:"pk varchar(255) notnull "`
	Sponsor   string     `xorm:"varchar(255) notnull"`
	StartTime string 	 `xorm:"varchar(255) notnull"`
	EndTime   string 	 `xorm:"varchar(255) notnull"`
	Participators string `xorm:"varchar(255) notnull"`
}


// 得到所有用户数据
func FindAllUser() ([]User, error) {	
	//返回[]map[string]string和error
	alluser, err := engine.QueryString("SELECT * FROM user_table")

	uSlice := []User{}
	
	for _, t := range alluser {
		uSlice = append(uSlice, User{t["user_name"], t["password"], t["email"], t["phone"]})
	}

	return uSlice, err
}

// 得到所有用户数据
func FindUserByName(name string) User {	

	u := &UserTable{UserName:name}
	_, err := engine.Get(u)
	if err != nil {
		panic(err)
	}
	return User{u.UserName, u.Password, u.Email, u.Phone}
}

//对未存在的用户进行插入，若存在则返回错误
func CreateUser_DB(user *User) error {
	_, err := engine.Insert(&UserTable{user.Name, user.Password, user.Email, user.Phone})
	return err
}

//删除用户
func DeleteUser_DB(user *User) error {
	// 通过 Delete 方法删除记录
	_, err := engine.Delete(&UserTable{UserName:user.Name})
	
	return err
}
//更新用户
func UpdateUser_DB(user *User) error {
	a := &UserTable{UserName:user.Name}
	_ , err := engine.Get(a)
	// 方法 Update 接受的第一个参数必须是指针地址，指向需要更新的内容。
	a.Password = user.Password
	a.Email = user.Email
	a.Phone = user.Phone
  _, err = engine.Update(a)
	return err
}

// 得到所有会议数据
func FindAllMeeting() ([]Meeting, error) {
	allmeeting, err := engine.QueryString("SELECT * FROM meeting_table")

	mSlice := []Meeting{}

	for _, t := range allmeeting {
		var pa []string
		if err := json.Unmarshal([]byte(t["participators"]), &pa); err != nil {
			panic(err)

		}

		mSlice = append(mSlice, Meeting{t["sponsor"], t["title"], t["start_time"], t["end_time"], pa})
	}	
	return mSlice, err
}

//对未存在的会议进行插入，若存在则返回错误
func CreateMeeting_DB(meeting *Meeting) error {
	
	_, err := engine.Insert(&MeetingTable{meeting.Title, meeting.Sponsor, meeting.StartDate,meeting.EndDate, GetJson(meeting.Participators)})
	return err
}

//删除会议
func DeleteMeeting_DB(meeting *Meeting) error {
	// 通过 Delete 方法删除记录
	_, err := engine.Delete(&MeetingTable{Title:meeting.Title})
	return err
}

//更新会议

func UpdateMeeting_DB(meeting *Meeting) error {
	a := &MeetingTable{Title:meeting.Title}
	_, err := engine.Get(a)
	// 方法 Update 接受的第一个参数必须是指针地址，指向需要更新的内容。
	a.Participators = GetJson(meeting.Participators)
	_, _ = engine.Update(a)
	return err
}


