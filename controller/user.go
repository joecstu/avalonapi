package controller

import (
	"avalonapi/data"
	"avalonapi/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Useronline(context *gin.Context) {

	var response struct {
		Status        string `json:",omitempty"` //"success | error | inactive"
		StatusMessage string `json:",omitempty"`
		Useronline    []model.User
	}

	//example, err := ds.Mongo.InsertExample(request.Example)
	err :=http.ErrAbortHandler
	err,response.Useronline = data.Useronline()

	if err != nil {
		response.Status = "GetUserOnlineSuccessful"
		context.JSON(http.StatusOK, response)
	}else{
		response.Status = "GetUserOnlineSuccessful"
		context.JSON(http.StatusOK, response)
	}

}
