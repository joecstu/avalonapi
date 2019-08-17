package controller

import (
	"avalonapi/data"
	"avalonapi/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(context *gin.Context) {
	var request struct {
		*model.User
	}
	var response struct {
		Status        string `json:",omitempty"` //"success | error | inactive"
		StatusMessage string `json:",omitempty"`
		Key           string
		Nickname      string
	}
	err := context.BindJSON(&request)
	if err != nil {
		response.Status = "ส่งข้อมูลมาผิดพลาด"
		response.StatusMessage = err.Error()
		context.JSON(http.StatusInternalServerError, response)
		return
	}
	//example, err := ds.Mongo.InsertExample(request.Example)
	err, nickname, key := data.Login(request.User)

	if err != nil {
		response.Status = "LoginFailed"
		response.StatusMessage = err.Error()
		context.JSON(http.StatusOK, response)
		return
	} else {
		response.Status = "LoginSuccessful"
		response.Key = key
		response.Nickname = nickname
		context.JSON(http.StatusOK, response)
	}

}
