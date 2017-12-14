**Agenda-Service**

-------------------
本人负责的主要工作为：</br>
1.设计API</br>
2.编写cmd_test.go</br>
3.辅助完成功能测试和mock测试</br>
4.编写README.md</br>

1.设计API完成的界面位于网址：</br>
https://retokani.docs.apiary.io/#</br>
与组长对应书写的agenda功能如下图：</br>
![这里写图片描述](http://img.blog.csdn.net/20171214082846013?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
![这里写图片描述](http://img.blog.csdn.net/20171214082909728?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
![这里写图片描述](http://img.blog.csdn.net/20171214082920542?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
![这里写图片描述](http://img.blog.csdn.net/20171214082929948?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvcXFfMzY4MTY5MTI=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

---------------------------------------------------------------------
</br>2.编写cmd_test.go.</br>
模块测试代码如下：</br>
测试UserRegister:</br>
```
func TestRegister(t *testing.T) {
	fmt.Println("=====> In TEST of UserRegister")
	UserRegisterCmd.Flags().Set("username", "Alice")
	UserRegisterCmd.Flags().Set("password", "123")
	UserRegisterCmd.Flags().Set("email", "Alice@163.com")
	UserRegisterCmd.Flags().Set("phone", "123")
	UserRegisterCmd.Run(UserRegisterCmd, nil)
}
```
测试UserLogin：
```
func TestLogin(t *testing.T) {
	fmt.Println("=====> In TEST of UserLogin")
	UserLoginCmd.Run(UserLoginCmd, nil)
	UserLoginCmd.Flags().Set("username", "Alice")
	UserLoginCmd.Flags().Set("password", "123")
	UserLoginCmd.Run(UserLoginCmd, nil)
}
```
测试users（列出所有的用户）：
```
func TestShowAllUsers(t *testing.T) {
	fmt.Println("=====> In TEST of ListAllUser")
	usersCmd.Run(usersCmd, nil)
}
```
测试MeetingCreate：
```
func TestCreateNewMeeting(t *testing.T) {
	fmt.Println("=====> In TEST of MeetingCreate")
	MeetingCreateCmd.Flags().Set("title", "testMeeting")
	MeetingCreateCmd.Flags().Set("members", "testUser0,testUser1")
	MeetingCreateCmd.Flags().Set("starttime", "2000/01/01/00:00")
	MeetingCreateCmd.Flags().Set("endtime", "2001/01/01/00:00")
	MeetingCreateCmd.Run(MeetingCreateCmd, nil)
}
```
测试meetings：(列出所有会议)
```
func TestShowAllMeetings(t *testing.T) {
	fmt.Println("=====> In TEST of ListAllMeeting")
	meetingsCmd.Run(meetingsCmd, nil)
}
```
测试UserDelete：
```
func TestUserDelete(t *testing.T) {
	fmt.Println("=====> In TEST of UserDelete")
	UserDeleteCmd.Run(UserDeleteCmd, nil)
	UserDeleteCmd.Run(usersCmd, nil)
}

```
测试方法：</br>
1.修改host为：http://localhost:8080完成本地测试。</br>
2.修改host为：https://private-c9f16-retokani.apiary-mock.com</br>
完成travis测试</br>

----------------------------------
3.辅助完成功能测试和mock测试：</br>
测试方法：</br>
1.修改host为：http://localhost:8080完成本地测试：</br>
```
$ go test
=====> In TEST of UserRegister
Register successfully.The account detail is :
{
  "Name": "Alice",
  "Password": "123",
  "Email": "Alice@163.com",
  "Phone": "123"
}

=====> In TEST of UserLogin
Using UserName and PassWord to login Agenda:

attention:If the PassWord is right,you can login Agenda and use it
If forget the PassWord,you must register another one User

Usage:
  agenda login -u [UserName] -p [PassWord] [flags]

Flags:
  -p, --password string   new user's password
  -u, --username string   new user's username

Global Flags:
      --config string   config file (default is $HOME/.cli.yaml)
Login successfully!
 username:Alice
=====> In TEST of ListAllUser
[
  {
    "Name": "Alice",
    "Password": "123",
    "Email": "Alice@163.com",
    "Phone": "123"
  }
]

=====> In TEST of MeetingCreate
To create a new meeting with:

[Title] the Title of the meeting
[Participator] the Participator of the meeting,the Participator can only attend one meeting during one meeting time
[StartTime] the StartTime of the meeting
[EndTime] the EndTime of the meeting

Usage:
  agenda create -t [Title] -p [Participator] -s [StartTime] -e [EndTime] [flags]

Flags:
  -e, --EndTime string         meeting's endTime
  -p, --Participator strings   meeting's participator
  -s, --StartTime string       meeting's startTime
  -t, --Title string           meeting title

Global Flags:
      --config string   config file (default is $HOME/.cli.yaml)
=====> In TEST of ListAllMeeting
[]

=====> In TEST of UserDelete
Delete successfully.Delete user failed.
PASS
ok  	Agenda-Service/cli/cmd	0.138s
```
2.修改host为：https://private-c9f16-retokani.apiary-mock.com</br>
完成travis测试:(测试细节在组长的测试网站)</br>
![这里写图片描述](https://camo.githubusercontent.com/3fbf21a5dcded84195243254682ee3c874c6a107/68747470733a2f2f7777772e7472617669732d63692e6f72672f4c65756e67436869486f2f4167656e64612d536572766963652e7376673f6272616e63683d6d6173746572)

----------------------------------
</br>4.编写README.md</br>
