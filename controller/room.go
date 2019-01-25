package controller

import (
	"avalonapi/data"
	"avalonapi/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRoom (context *gin.Context) {
	var request struct {
		Key string
		Nickname string
	}
	var response struct {
		Status        string `json:",omitempty"` //"success | error | inactive"
		StatusMessage string `json:",omitempty"`
		Room model.Room
	}
	err := context.BindJSON(&request)

	if err != nil {
		response.Status = "ส่งข้อมูลมาผิดพลาด"
		response.StatusMessage = err.Error()
		context.JSON(http.StatusInternalServerError, response)
		return
	}
	status:=1
	response.Room,err,status = data.CreateRoom(request.Nickname,request.Key)
	if status==0 {
		response.Status = "CreateRoomFailed"
		response.Room=model.Room{}
		context.JSON(http.StatusOK, response)
	}else{
		response.Status = "CreateRoomSuccessful"
		context.JSON(http.StatusOK, response)
	}
}
func Joinroom (context *gin.Context) {
	var request struct {
		Key string
		Nickname string
		Roomid int
	}
	var response struct {
		Status        string `json:",omitempty"` //"success | error | inactive"
		StatusMessage string `json:",omitempty"`
		Room model.Room
	}
	err := context.BindJSON(&request)

	if err != nil {
		response.Status = "ส่งข้อมูลมาผิดพลาด"
		response.StatusMessage = err.Error()
		context.JSON(http.StatusInternalServerError, response)
		return
	}
	status:=1
	response.Room,err,status = data.Joinroom(request.Nickname,request.Key,request.Roomid)
	if status==0 {
		response.Status = "JoinRoomFailed"
		response.Room=model.Room{}
		context.JSON(http.StatusOK, response)
	}else{
		response.Status = "JoinRoomSuccessful"
		context.JSON(http.StatusOK, response)
	}
}


func Getroom (context *gin.Context) {
	var request struct {
		Roomid int
	}
	var response struct {
		Status        string `json:",omitempty"` //"success | error | inactive"
		StatusMessage string `json:",omitempty"`
		Room model.Room
	}
	err := context.BindJSON(&request)

	if err != nil {
		response.Status = "ส่งข้อมูลมาผิดพลาด"
		response.StatusMessage = err.Error()
		context.JSON(http.StatusInternalServerError, response)
		return
	}
	status:=1
	response.Room,err,status = data.Getroom(request.Roomid)
	if status==0 {
		response.Status = "GetRoomFailed"
		response.Room=model.Room{}
		context.JSON(http.StatusOK, response)
	}else{
		response.Status = "GetRoomSuccessful"
		context.JSON(http.StatusOK, response)
	}
}