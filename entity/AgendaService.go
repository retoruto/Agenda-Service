package entity

import (
	"fmt"
)




func StartAgenda() bool {
	ReadFromDb()
	ReadCurrentUser()
	if CurrentUser.Name == "" {
		return false
	}
	return true
}
/*
func QuitAgenda() {
	writeToFile()
	writeCurrentUser()
}*/

/**
* check if the username match password
* @param userName the username want to login
* @param password the password user enter
* @return if success, true will be returned
*/
//登录命令不需要调用StartAgenda,但需要调用QuitAgenda来保存登录信息
func UserLogIn(userName string, password string) bool{
		ReadFromDb()
		ReadCurrentUser()
		if (CurrentUser.Name != "") {
			return false
		}
		filter := func(u *User) bool {
		  return u.getName() == userName && u.getPassword() == password
		}

		ulist := queryUser(filter)
		if (len(ulist) == 0) {
			return false
		} else {
			//当前用户信息
			CurrentUser = ulist[0]
			writeCurrentUser()
			return true
		}
	}

/**
 * regist a user
 * @param userName new user's username
 * @param password new user's password
 * @param email new user's email
 * @param phone new user's phone
 * @return if success, true will be returned
 */	
func UserRegister(userName, password, email, phone string) bool {
	filter := func(u *User) bool {
		return u.getName()== userName
	}
	ulist := queryUser(filter)

	if (len(ulist) == 0) {
		createUser(User{userName, password, email, phone})
		return true
	} else {
		return false
	}
}

/**
 * delete a user
 * @param userName user's username
 * @param password user's password
 * @return if success, true will be returned
 */	
func DeleteUser(userName string, password string) bool {
	uf := func(u *User) bool {
		return (u.getName()== userName) && (u.getPassword() == password)
	}
	mf := func(m *Meeting) bool {
		return m.getSponsor() == userName || m.isParticipator(userName)
	}
	if (deleteUser(uf) != 0) {
		deleteMeeting(mf)
		if (userName == CurrentUser.Name) {
			CurrentUser.InitUser("", "", "", "")
			writeCurrentUser()
		}
		return true
	} else {
		return false
	}
}

/**
 * list all users from storage
 * @return a user list result
 */
func ListAllUsers() []User {
	filter := func(u *User) bool {
		return true
	}
	return queryUser(filter)
}

/**
 * create a meeting
 * @param userName the sponsor's userName
 * @param title the meeting's title
 * @param participator the meeting's participator
 * @param startData the meeting's start date
 * @param endData the meeting's end date
 * @return if success, true will be returned
 */
func CreateMeeting(userName, title, startDate, endDate string, participator []string) bool {
	var sd Date
	var ed Date
	sd = stringToDate(startDate)
	ed = stringToDate(endDate)

	if ((!sd.isValid()) || (!ed.isValid())) {
		fmt.Println("日期不合法")
		return false
	}
	if (sd.GreaterOrEqual(ed)) {
		fmt.Println("开始日期不可大于结束日期")
		return false
	}
	
	/*-------------------1-------------------------*/
	uf := func(u *User) bool {
		return userName == u.getName()
	}

	ulist := queryUser(uf)
	if (len(ulist) == 0) {
		fmt.Println("发起人未注册")
		return false
	}
	
	/*-------------------2-------------------------*/
	uf2 := func(u *User) bool {
		for _, p := range participator {
			if p == u.getName() {
				return true
			}
		}
		return false
	}
	ulist2 := queryUser(uf2)
	if (len(ulist2) != len(participator)) {
		fmt.Println("存在参与者未注册")
		return false
	}
	
	/*-------------------3-------------------------*/
	uf3 := func(m *Meeting) bool {
	return title == m.getTitle()
	}
	ulist3 := queryMeeting(uf3)
	if (len(ulist3) != 0) {
		fmt.Println("会议主题已存在")
		return false
	}
	
	/*-------------------4--------------------------*/
	uf4 := func(m *Meeting) bool {
		if (!(userName == m.getSponsor() || m.isParticipator(userName))) {
			return false
		}
		if ((userName == m.getSponsor() || m.isParticipator(userName)) &&
			(sd.GreaterOrEqual(stringToDate(m.getEndDate())) ||
			ed.SmallerOrEqual(stringToDate(m.getStartDate())))) {
			return false
		} else {
			return true
		}
	}
	ulist4 := queryMeeting(uf4)
	if (len(ulist4) != 0) {
		fmt.Println("与发起人或者参与者其他会议冲突")
		return false
	}
	
	/*-------------------5--------------------------*/
	uf5 := func(m *Meeting) bool {
	for _, p := range participator {
		if ( !(p == m.getSponsor() || m.isParticipator(p)) ) {
				return false
			}
			if ((p == m.getSponsor() || m.isParticipator(p)) &&
			(sd.GreaterOrEqual(stringToDate(m.getEndDate())) ||
			ed.SmallerOrEqual(stringToDate(m.getStartDate())))) {
				return false
			} else {
				return true
			}
		}
		return false
	}
	ulist5 := queryMeeting(uf5)
	if (len(ulist5) != 0) {
		fmt.Println("与发起人其他会议冲突")
		return false
	}
	
	/*-------------------6--------------------------*/

	for i := 0; i < len(participator); i++ {
		for j := i + 1; j < len(participator); j++ {
			if (participator[i] == participator[j]) {
				fmt.Println("参与者不能重复")
				return false
			}
		}
	}
	
	/*-------------------７--------------------------*/
	for _, p := range participator {
		if (userName == p) {
			fmt.Println("参与者不能有发起者")
			return false
		}
	}

	/*-------------------8--------------------------*/
	if (len(participator) == 0) {
		fmt.Println("参与者不能为空")
		return false
	}

	createMeeting(Meeting{userName, title, startDate, endDate, participator})
	return true
}

/**
* search a meeting by username, time interval (user as sponsor or participator)
* @param uesrName the user's userName
* @param startDate time interval's start date
* @param endDate time interval's end date
* @return a meeting list result
*/
func MeetingQuery(sponsor, startDate, endDate string) []Meeting {
	var ttt []Meeting
	sd := stringToDate(startDate)
	ed := stringToDate(endDate)
	if (sd.isMoreThan(ed) || !sd.isValid() || !ed.isValid()) {
		fmt.Println("日期不合法")
		return ttt //此时a为空
	}

	filter := func(a *Meeting) bool {
		if ((a.Sponsor == sponsor || a.isParticipator(sponsor)) &&
		 (stringToDate(a.getEndDate()).GreaterOrEqual(sd)&&stringToDate(a.getStartDate()).SmallerOrEqual(sd))) {
			return true
		}
		if ((a.Sponsor == sponsor || a.isParticipator(sponsor)) && 
		(stringToDate(a.getStartDate()).SmallerOrEqual(ed)) && stringToDate(a.getStartDate()).GreaterOrEqual(sd)) {
			return true
		}
		return false
	}
	return queryMeeting(filter)
}
/**
* list all meetings the user take part in
* @param userName user's username
* @return a meeting list result
*/
func ListAllMeetings(name string) []Meeting {
	filter := func(a *Meeting) bool {
		return a.Sponsor == name || a.isParticipator(name)
	}
	return queryMeeting(filter)
}

/**
* delete a meeting by title and its sponsor
* @param userName sponsor's username
* @param title meeting's title
* @return if success, true will be returned
*/
func ListAllSponsorMeetings(name string) []Meeting {
	filter := func(a *Meeting) bool {
		return  name == a.Sponsor
	}
	return queryMeeting(filter)
}

/**
* list all meetings the user take part in and sponsor by other
* @param userName user's username
* @return a meeting list result
*/
func ListAllParticipateMeetings(name string) []Meeting {
	filter := func(a *Meeting) bool {
		return a.isParticipator(name)
	}
	return queryMeeting(filter)
}

/**
* delete a meeting by title and its sponsor
* @param userName sponsor's username
* @param title meeting's title
* @return if success, true will be returned
*/
func DeleteMeeting(name, title string) bool {
	filter := func(a *Meeting) bool {
		return a.Title == title && a.Sponsor == name
	}
	return deleteMeeting(filter) > 0
}
/**
* delete all meetings by sponsor
* @param userName sponsor's username
* @return if success, true will be returned
*/
func DeleteAllMeetings(name string) bool {
	filter := func(a *Meeting) bool {
		return  a.Sponsor == name
	}
	return deleteMeeting(filter) > 0
}

/**
* add participator into the meetings created by sponsor
* @param title meeting's title
* @return if success, true will be returned
*/
func Addparticipator(title string, participator []string) bool {
	/*---------------------1------------------------*/
	if len(participator) == 0 {
		fmt.Println("必须至少添加一个参与者")
		return false
	} 
	/*---------------------2------------------------*/
	filter1 := func(m *Meeting) bool {
		return  m.Sponsor == CurrentUser.Name && m.Title == title
	}
	//找到创建人的一条会议
	mlist := queryMeeting(filter1)
	if (len(mlist) == 0) {
		fmt.Println("找不到当前用户创建的该会议")
		return false
	}
	/*---------------------3------------------------*/
	for i := 0; i < len(participator); i++ {
		for j := i + 1; j < len(participator); j++ {
			if (participator[i] == participator[j]) {
				fmt.Println("添加的参与者不能重复")
				return false
			}
		}
	}
	/*----------------------4-----------------------*/
	for _, p := range mlist[0].Participators {
		for _, pp := range participator{
			if (pp == p) {
				fmt.Println("存在添加的参与者已参与该会议")
				return false
			}
		}
	}
	/*----------------------5-----------------------*/
	for _, p := range participator {
		if (CurrentUser.Name == p) {
			fmt.Println("参与者不能有发起者")
			return false
		}
	}
	/*-------------------6--------------------------*/
	sd := stringToDate(mlist[0].StartDate)
	ed := stringToDate(mlist[0].EndDate)
	filter2 := func(m *Meeting) bool {
		for _, p := range participator {
			if (!(p == m.getSponsor() || m.isParticipator(p))) {
				return false
			}
			if ((p == m.getSponsor() || m.isParticipator(p)) &&
				(sd.GreaterOrEqual(stringToDate(m.getEndDate())) ||
				ed.SmallerOrEqual(stringToDate(m.getStartDate())))) {
				return false
			} else {
				return true
			}
		}
		return true
	}
	mlist1 := queryMeeting(filter2)
	if (len(mlist1) != 0) {
		fmt.Println("添加的参与者与其他会议冲突")
		return false
	}
	
	/*-------------------7--------------------------*/
	filter3 := func(m *Meeting) bool {
		for _, p := range participator {
			if (!(p == m.getSponsor() || m.isParticipator(p))) {
				return false
			}
			if ((p == m.getSponsor() || m.isParticipator(p)) &&
				(sd.GreaterOrEqual(stringToDate(m.getEndDate())) ||
				ed.SmallerOrEqual(stringToDate(m.getStartDate())))) {
				return false
			} else {
				return true
			}
		}
		return true
	}
	mlist2 := queryMeeting(filter3)
	if (len(mlist2) != 0) {
		fmt.Println("添加的参与者与其他会议冲突")
		return false
	}
	/*---------------------------------------------*/
	//add
	for _, p := range participator {
		mlist[0].Participators = append(mlist[0].Participators, p)
	}

	//写回原会议
	for i, m := range meetinglist {
		if m.Title ==  mlist[0].Title {
			meetinglist[i] = mlist[0]
		}
	}
	
	return true
}

/**
* remove participator from the meetings created by sponsor
* @param title meeting's title
* @return if success, true will be returned
*/
func Removeparticipator(title string, participator []string) bool {
	/*---------------------1------------------------*/
	if len(participator) == 0 {
		fmt.Println("必须至少删除一个参与者")
		return false
	} 
	/*---------------------2------------------------*/
	filter1 := func(m *Meeting) bool {
		return  m.Sponsor == CurrentUser.Name && m.Title == title
	}
	//找到创建人的一条会议
	mlist := queryMeeting(filter1)
	if (len(mlist) == 0) {
		fmt.Println("找不到当前用户创建的该会议")
		return false
	}
	/*----------------------3-----------------------*/
	for _, p := range participator {
		if (CurrentUser.Name == p) {
			fmt.Println("不能删除发起者,你可以退出会议（quit）")
			return false
		}
	}
	/*----------------------4-----------------------*/
	if (len(participator) > len(mlist[0].Participators)) {
		fmt.Println("删除人数不能大于原有人数")
		return false
	}
	/*----------------------5-----------------------*/
	for _, p := range participator {
		if (mlist[0].isParticipator(p) == false) {
			fmt.Println("存在待删除者没有参与该会议")
			return false
		}	
	}
	/*----------------------6-----------------------*/
	//delete
	for _, p := range participator {
		n := 0
		for i, pp := range mlist[0].Participators {
			if p == pp {
				mlist[0].Participators[i] = mlist[0].Participators[len(mlist[0].Participators) - 1 - n]			
				n++
			}
		}
		mlist[0].Participators = mlist[0].Participators[:len(mlist[0].Participators)  - n]
	}
	/*----------------------7-----------------------*/

	//重新写回meetinglist
	for i, m := range meetinglist {
		if m.Title ==  mlist[0].Title {
			meetinglist[i] = mlist[0]
		}
	}
	
	filter2 := func(m *Meeting) bool {
		if len(m.Participators) == 0 {
			return true
		}
		return false
	}
	if deleteMeeting(filter2) == 1 {
		fmt.Println("该会议的参与者已清空，会议已删除")
		return true
	}

	return true
}


/**
* CurrentUser quits meeting in which the currentUser participate 
* @param title meeting's title
* @return if success, true will be returned
*/
func QuitMeeting(title string) bool {
	/*---------------------1------------------------*/
	filter1 := func(m *Meeting) bool {
		return  (m.Sponsor == CurrentUser.Name ||  m.isParticipator(CurrentUser.Name))&& m.Title == title
	}
	//找到当前用户参加的一条会议，当前用户为参与者或发起者
	mlist := queryMeeting(filter1)
	if (len(mlist) == 0) {
		fmt.Println("找不到当前用户参加的该会议")
		return false
	}
	/*----------------------2-----------------------*/
	//当前用户为该会议的发起者，则删除该会议
	if (CurrentUser.Name == mlist[0].Sponsor) {
		if DeleteMeeting(CurrentUser.Name, title) {
			fmt.Println("当前用户为该会议的发起者，该会议已被删除")
			return true
		}
	}
	
	/*----------------------3-----------------------*/
	//delete
	n := 0
	for i, pp := range mlist[0].Participators {
		if CurrentUser.Name == pp {
			mlist[0].Participators[i] = mlist[0].Participators[len(mlist[0].Participators) - 1 - n]			
			n++
		}
	}
	mlist[0].Participators = mlist[0].Participators[:len(mlist[0].Participators)  - n]
	/*----------------------7-----------------------*/

	//重新写回meetinglist
	for i, m := range meetinglist {
		if m.Title ==  mlist[0].Title {
			meetinglist[i] = mlist[0]
		}
	}
	
	filter2 := func(m *Meeting) bool {
		if len(m.Participators) == 0 {
			return true
		}
		return false
	}
	if deleteMeeting(filter2) == 1 {
		fmt.Println("该会议的参与者已清空，会议已删除")
		return true
	}

	return true
}


