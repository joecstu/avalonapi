package main

import (
	"avalonapi/controller"
	"github.com/gin-gonic/gin"
)

//เราจะทำสำเร็จไปด้วยกันนะ


type ID struct {
	Username string
	Password string
}


func main()  {
	r := gin.Default()
	//r.GET("/ping",Login )
	r.POST("/register",controller.Register)
	r.POST("/login",controller.Login)
	r.POST("/logout",controller.Logout)
	r.POST("/createnickname",controller.CreateNickName)

	r.Run(":1312") // listen and serve on 0.0.0.0:8080
}

