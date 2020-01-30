package userdata

import (
	"../../db"
)

type Users struct {
	ID       int
	Username string
	Email 	 string
	Password string
}

var user Users

func GetUserById(id string) Users {
	//connect to db
	db.Init()

	libraryDb := db.GetDB()

	defer db.Close()

	libraryDb.SingularTable(true)
	libraryDb.Select("id, username, email").First(&user, "id = ?", id)

	return user
}

func Login(email string, password string) (bool, bool) {
	db.Init()
	defer db.Close()

	libraryDb := db.GetDB()

	libraryDb.First(&user, "email = ?", email)
	if user.ID == 0 {
		return false, false
	}
	if user.Password != password {
		return true, false
	}

	return true, true
}
