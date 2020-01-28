package handler

import (
	user_v1 "../controllers/v1"
	user_v2 "../controllers/v2"
	"github.com/gin-gonic/gin"
)

func Init() {
	router := StartRouter()
	router.Run(":3000")

}

func StartRouter() *gin.Engine {
	// release mode will not output debug information
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	ctrlV1 := user_v1.Controller{}
	ctrlV2 := user_v2.Controller{}

	v1Group := router.Group("/v1")
	{
		v1Group.GET("/userInfo", ctrlV1.ShowUserInfo)
		v1Group.GET("/", ctrlV1.ShowIndex)
	}

	v2Group := router.Group("/v2")
	{
		v2Group.GET("/userInfo", ctrlV2.ShowUserInfo)
		v2Group.GET("/", ctrlV2.ShowIndex)
	}

	return router
}
