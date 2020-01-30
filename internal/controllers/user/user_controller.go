package user

import (
	"../../models/users"
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (pc Controller) GetUserInfoById(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(422, "Please input id to get user info.")
	} else {
		userInfo := userdata.GetUserById(id)
		c.JSON(200, userInfo)
	}
}

func (pc Controller) ShowIndex(c *gin.Context) {
	c.JSON(200, "Hello API")
}
