package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"avalonapi/model"
	"avalonapi/data"
)

func Register(con *gin.Context) {
	var request struct {
		*model.UserRegis
	}
	var response struct {
		Status        string `json:",omitempty"` //"success | error | inactive"
		StatusMessage string `json:",omitempty"`
	}
	err := con.BindJSON(&request)
	if err != nil {
		response.Status = "ส่งข้อมูลมาผิดพลาด"
		response.StatusMessage = err.Error()
		con.JSON(http.StatusInternalServerError, response)
		return
	}
	//example, err := ds.Mongo.InsertExample(request.Example)
	err = data.CreateUser(request.UserRegis)

	if err != nil {
		response.Status = "Email Used"
		response.StatusMessage = err.Error()
		con.JSON(http.StatusOK, response)
		return
	}else{
		response.Status = "RegisterSuccessful"
		response.StatusMessage = "Insert example"
		con.JSON(http.StatusOK, response)
	}

}
