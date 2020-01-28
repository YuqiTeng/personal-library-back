package user 

import (
	"../../models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Controller struct{}

func (pc Controller) ShowUserInfo(c *gin.Context) {
	id := c.Query("id")

	if len(id) == 0 {
		c.JSON(422, "v2: Please input easyid to query user info.")
	} else {
		id, _ := strconv.Atoi(c.Query("id"))
		userInfo := userdata.GetUserById(id)
		c.JSON(200, userInfo)
	}
}

func (pc Controller) ShowIndex(c *gin.Context) {
	c.JSON(200, "Hello API")
}
