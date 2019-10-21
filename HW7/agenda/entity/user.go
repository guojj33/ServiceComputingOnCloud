package entity
type User struct {
	Username     string
	Password     string
	Mail         string
	Telephone    string
}

func (user *User) InitUser(username, password, email, telephone string) {
	user.Username = username
	user.Password = password
	user.Mail = email
	user.Telephone = telephone
}
