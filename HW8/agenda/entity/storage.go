package entity

import (
	"fmt"
	"encoding/json"
	"os"
)

var (
	curUser    string
	allUsers   []User
	userLib    string = "./data/user.json"
)

func Init() {
	ReadUserFile()
}

func ReadUserFile() {
	file, err := os.Open(userLib)
	if err != nil {
		return
	}
	state, _ := file.Stat()
	if state.Size() == 0 {
		return
	}
	buffer := make([]byte, state.Size())
	_, err = file.Read(buffer)
	if err != nil {
		return
	}
	buffer = []byte(os.ExpandEnv(string(buffer)))
	err = json.Unmarshal(buffer, &allUsers)
	if err != nil {
		return
	}
}

func UpdateLib() {
	WriteUserFile()
}

func WriteUserFile() {
	userRec, err := json.Marshal(allUsers)
	if err != nil {
		fmt.Println(err)
	}
	f, _ := os.Create(userLib)
	defer f.Close()
	f.WriteString(string(userRec))
}

func isUserUnique(user User) bool {
	for _, useri := range allUsers {
		if useri.Username == user.Username {
			return false		
		}
	}
	return true
}

func CreateUser(name, password, mail, telephone string) bool {
	if name == "" || password == "" || mail == "" || telephone == "" {
		return false	
	}
	user := User{
		Username:  name,
		Password:  password,
		Mail:      mail,
		Telephone: telephone,
	}
	if isUserUnique(user) {
		allUsers = append(allUsers, user)
		UpdateLib()
		return true	
	} else {
		return false
	}
}

func DeleteUser(username, password string) bool {
	for i, useri := range allUsers {
		if useri.Username == username && useri.Password == password {
			allUsers[i] = allUsers[len(allUsers)-1]
			allUsers    = allUsers[0:len(allUsers)-1]
			UpdateLib()
			return true
		}
	}
	return false
}
