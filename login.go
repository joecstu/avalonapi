package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	var request struct {
		//*model.Example
	}
	var response struct {
		Status        string `json:",omitempty"` //"success | error | inactive"
		StatusMessage string `json:",omitempty"`
		//Example      // *model.Example
	}
	err := context.BindJSON(&request)
	if err != nil {
		response.Status = "error"
		response.StatusMessage = err.Error()
		context.JSON(http.StatusInternalServerError, response)
		return
	}
//	example, err := ds.Mongo.InsertExample(request.Example)
	if err != nil {
		response.Status = "error"
		response.StatusMessage = err.Error()
		context.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Status = "success"
	response.StatusMessage = "Insert example"
	//response.Example = example
	context.JSON(http.StatusOK, response)
}
