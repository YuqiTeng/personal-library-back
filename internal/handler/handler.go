package handler

import (
	"../controllers/common"
	"../controllers/user"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := StartRouter()
	router.Run(":3000")

}

func StartRouter() *gin.Engine {
	// release mode will not output debug information
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	userController := user.Controller{}
	commonController := common.Controller{}

	router.GET("/", commonController.ShowIndex)
	router.GET("/userInfo/:id", userController.GetUserInfoById)
	router.POST("/login", commonController.Login)

	return router
}
