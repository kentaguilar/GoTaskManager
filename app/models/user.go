package model

import (
	"../helpers"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name string
	Username string
	Password string
}

func (givenUser *User) SaveUser() string{
    db := helper.OpenDb()

	var user string
    message := "Success"

    db.QueryRow("select username from users where username=?", givenUser.Username).Scan(&user)    

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(givenUser.Password), bcrypt.DefaultCost)
    if err != nil{
        message = "Server error, unable to create your account"
    }

    _, err = db.Exec("insert into users(name, username, password) values(?, ?, ?)", givenUser.Name, givenUser.Username, hashedPassword)
    if err != nil{
        message = "Server error, unable to create your account"
    }

    helper.CloseDb(db)

    return message
}