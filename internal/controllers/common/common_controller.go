package common

import (
	"github.com/gin-gonic/gin"
	"github.com/personal-library-back/internal/models/users"
)

type Controller struct{}

type LoginInfo struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (pc Controller) Login(c *gin.Context) {
	var loginInfo LoginInfo

	c.BindJSON(&loginInfo)

	isEmailValid, isPasswordValid := userdata.Login(loginInfo.Email, loginInfo.Password)
	if isEmailValid && isPasswordValid {
		c.JSON(200, "login success")
	} else if !isEmailValid {
		c.JSON(403, "User does not exist")
	} else {
		c.JSON(403, "Password wrong")
	}
}

func (pc Controller) ShowIndex(c *gin.Context) {
	c.JSON(200, "Welcome to personal library project")
}
