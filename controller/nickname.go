package controller

import (
	"avalonapi/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangeNickName(context *gin.Context) {
	var request struct {
		Key      string
		Nickname string
	}
	var response struct {
		Status        string `json:",omitempty"` //"success | error | inactive"
		StatusMessage string `json:",omitempty"`
		Nickname      string
	}
	err := context.BindJSON(&request)

	if err != nil {
		response.Status = "ส่งข้อมูลมาผิดพลาด"
		response.StatusMessage = err.Error()
		context.JSON(http.StatusInternalServerError, response)
		return
	}

	err, status := data.ChangeNickName(request.Nickname, request.Key)
	if status == 0 {
		response.Status = "ChangeNickNameFailed"
		context.JSON(http.StatusOK, response)
	} else {
		response.Status = "ChangeNickNameSuccessful"
		response.Nickname = request.Nickname
		context.JSON(http.StatusOK, response)
	}
}
