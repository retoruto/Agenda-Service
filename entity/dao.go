package entity

// 得到所有用户数据
func FindAllUser() []User {	
	//返回[]map[string]string和error
	alluser, err := engine.QueryString("SELECT * FROM UserTable")
	CheckErr(err)
	
	var uSlice []User
	for _, t := range alluser {
		uSlice = append(uSlice, User{t["Name"], t["Password"], t["Email"], t["Phone"]})
	}
	return uSlice
}

//对未存在的用户进行插入，若存在则返回错误
func CreateUser_DB(user *User) error {
	_, err := engine.Insert(user)
	return err
}

//删除用户
func DeleteUser_DB(user *User) error {
	// 通过 Delete 方法删除记录
	_, err := engine.Delete(user)
	return err
}

//更新用户
func UpdateUser_DB(user *User) error {
	a := &User{Name:user.Name}
	_ , _ := engine.Get(a)
	// 方法 Update 接受的第一个参数必须是指针地址，指向需要更新的内容。
	_, err := engine.Update(a)
	return err
}

// 得到所有会议数据
func FindAllMeeting() []Meeting {
	allmeeting, err := engine.QueryString("SELECT * FROM MeetingTable")
	CheckErr(err)

	var mSlice []Meeting

	for _, t := range allmeeting {
		var pa []string
		if err := json.Unmarshal([]byte(t["Participators"]), &pa); err != nil {
			panic(err)

		}

		var sd string
		if err := json.Unmarshal([]byte(t["StartDate"]), &st); err != nil {
			panic(err)
		}

		var ed string
		if err := json.Unmarshal([]byte(t["EndDate"]), &et); err != nil {
			panic(err)
		}

		mSlice = append(mSlice, Meeting{t["Sponsor"], t["Title"], sd, ed, pa})
	}	
	return mSlice
}

//对未存在的会议进行插入，若存在则返回错误
func CreateMeeting_DB(meeting *Meeting) error {
	_, err := engine.Insert(meeting)
	return err
}

//删除会议
func DeleteMeeting_DB(meeting *Meeting) error {
	// 通过 Delete 方法删除记录
	_, err := engine.Delete(meeting)
	return err
}

//更新会议
/*
func UpdateMeeting_DB(meeting *Meeting) error {
	a := &Meeting{Name:user.Name}
	_ , _ := engine.Get(a)
	// 方法 Update 接受的第一个参数必须是指针地址，指向需要更新的内容。
	_, err := engine.Update(a)
	return err
}
*/

