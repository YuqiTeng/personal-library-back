package userdata

import (
	"github.com/personal-library-back/internal/models/db"
	"strconv"
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

	intId, _ := strconv.Atoi(id)
	libraryDb.SingularTable(true)
	libraryDb.Select("id, username, email").First(&user, "id = ?", intId)

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
