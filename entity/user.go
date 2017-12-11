package entity

type User struct {
	Name     string
	Password string
	Email    string
	Phone    string
}

func (m *User) InitUser(t_userName, t_userPassword, t_userEmail, t_userPhone string) {
	m.Name = t_userName
	m.Password = t_userPassword
	m.Email = t_userEmail
	m.Phone = t_userPhone
}

func (m *User) copyUser(t User) {
	m.Name = t.Name
	m.Password = t.Password
	m.Email = t.Email
	m.Phone = t.Phone
}

func (m User) getName() string {
	return m.Name
}

func (m *User) setName(n string) {
	m.Name = n
}

func (m *User) getPassword() string {
	return m.Password
}

func (m *User) setPassword(p string) {
	m.Password = p
}
func (m User) getEmail() string {
	return m.Email
}

func (m *User) setEmail(e string) {
	m.Email = e
}
func (m User) getPhone() string {
	return m.Phone
}

func (m *User) setPhone(p string) {
	m.Phone = p
}
