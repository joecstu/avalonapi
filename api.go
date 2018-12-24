package main

import (
	"github.com/gin-gonic/gin"

	"avalonapi/model"
	"net/http"
	"avalonapi/data"
)


type ID struct {
	Username string
	Password string
}


func main() {
	r := gin.Default()
	//r.GET("/ping",Login )
	r.POST("/register",func (context *gin.Context) {
		var request struct {
			*model.UserRegis
		}
		var response struct {
			Status        string `json:",omitempty"` //"success | error | inactive"
			StatusMessage string `json:",omitempty"`
		}
		err := context.BindJSON(&request)
		if err != nil {
			response.Status = "ส่งข้อมูลมาผิดพลาด"
			response.StatusMessage = err.Error()
			context.JSON(http.StatusInternalServerError, response)
			return
		}
		//example, err := ds.Mongo.InsertExample(request.Example)
		err = data.CreateUser(request.UserRegis)

		if err != nil {
			response.Status = "Email Used"
			response.StatusMessage = err.Error()
			context.JSON(http.StatusOK, response)
			return
		}else{
			response.Status = "RegisterSuccessful"
			response.StatusMessage = "Insert example"
			context.JSON(http.StatusOK, response)
		}

	})
	//r.POST("/register",Register)
	r.Run(":1312") // listen and serve on 0.0.0.0:8080
}