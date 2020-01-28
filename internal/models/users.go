package userdata

import (
	"../db"
)

type Users struct {
	ID       int
	Username string
	Email 	 string
	password string
}

func GetUserById(id int) Users {
	//connect to db
	db.Init();

	libraryDb := db.GetDB();

	defer db.Close()

	libraryDb.SingularTable(true)

	var user Users

	libraryDb.First(&user, "id = ?", id)

	return user
}
