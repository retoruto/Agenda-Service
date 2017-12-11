package entity
type UserTable struct {
	UserName string `xorm:"varchar(255) notnull unique"`
	Password string `xorm:"varchar(255) notnull"`
	Email    string `xorm:"varchar(255) notnull"`
	Phone    string `xorm:"varchar(255) notnull"`
}

type MeetingTable struct {
	Title     string     `xorm:"pk varchar(255) notnull unique"`
	Sponsor   string     `xorm:"varchar(255) notnull"`
	StartTime string 	 `xorm:"varchar(255) notnull"`
	EndTime   string 	 `xorm:"varchar(255) notnull"`
	Participators string `xorm:"varchar(255) notnull"`
}
