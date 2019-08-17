package controller

import (
	"avalonapi/data"
	"avalonapi/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(context *gin.Context) {
	var request struct {
		*model.User
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
	err = data.CreateUser(request.User)

	if err != nil {
		response.Status = "Email Used"
		response.StatusMessage = err.Error()
		context.JSON(http.StatusOK, response)
		return
	} else {
		response.Status = "RegisterSuccessful"
		response.StatusMessage = "Insert example"
		context.JSON(http.StatusOK, response)
	}

}
