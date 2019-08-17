package main

import (
	"avalonapi/controller"

	"github.com/gin-gonic/gin"
)

//เราจะทำสำเร็จไปด้วยกันนะ

func main() {
	r := gin.Default()
	//r.GET("/ping",Login )
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/logout", controller.Logout)
	r.POST("/changenickname", controller.ChangeNickName)
	r.GET("/useronline", controller.Useronline)
	r.GET("/logoutall", controller.LogoutAll)
	r.POST("/createroom", controller.CreateRoom)
	r.POST("/joinroom", controller.Joinroom)
	r.POST("/getroom", controller.Getroom)
	r.GET("/getallroom", controller.GetAllroom)
	r.GET("/deleteallroom", controller.DeleteAllroom)
	r.POST("leftroom", controller.LeftRoom)

	r.Run(":1312") // listen and serve on 0.0.0.0:8080
}
