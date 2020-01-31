package user

import (
	"github.com/personal-library-back/internal/models/users"
	"github.com/personal-library-back/internal/resource/users"
	"github.com/gin-gonic/gin"
	"fmt"
)

type Controller struct{}

func (pc Controller) GetUserInfoById(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		msg := "Please input id to get user info"
		res := usersresource.SetLoginResponse("Error", "", "", msg)
		fmt.Println(res)
		c.JSON(422, res)
	} else {
		userInfo := userdata.GetUserById(id)
		if userInfo.ID == 0 {
			msg := fmt.Sprintf("User not found with id: " + id)
			res := usersresource.SetLoginResponse("Success", "", "", msg)
			c.JSON(200, res)	
		} else {
			res := usersresource.SetLoginResponse("Success", userInfo.Username, userInfo.Email, "")
			c.JSON(200, res)
		}
	}
}

func (pc Controller) ShowIndex(c *gin.Context) {
	c.JSON(200, "Hello API")
}
