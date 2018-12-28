package controller

import (
	"avalonapi/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout (context *gin.Context) {
	var request struct {
		Key string
	}
	var response struct {
		Status        string `json:",omitempty"` //"success | error | inactive"
		StatusMessage string `json:",omitempty"`
		Nickname string
	}
	err := context.BindJSON(&request)
	if err != nil {
		response.Status = "ส่งข้อมูลมาผิดพลาด"
		response.StatusMessage = err.Error()
		context.JSON(http.StatusInternalServerError, response)
		return
	}

	err = data.Logout(request.Key)
	if err != nil {
		response.Status = "LogoutFailed"
		context.JSON(http.StatusOK, response)
	}else{
		response.Status = "LogoutSuccessful"
		context.JSON(http.StatusOK, response)
	}



}