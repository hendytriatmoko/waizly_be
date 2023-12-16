package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task_microservices/common"
	"task_microservices/daos"
	"task_microservices/models"
)

type Merk struct {
	daos daos.Merk
}

func (u *Merk) TaskCreate(c *gin.Context) {

	params := models.CreateTask{}

	response := models.Response{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.ApiMessage = "validation " + mess
		c.JSON(400, response)
	} else {

		data, err := u.daos.TaskCreate(params)

		if err != nil {
			response.ApiStatus = 0
			response.ApiMessage = err.Error()
			c.JSON(500, response)
		} else {
			response.ApiStatus = 1
			response.Data = data
			response.ApiMessage = common.StatusSukses
			c.JSON(http.StatusOK, response)

		}

	}

}

func (u *Merk) GetDataTask(c *gin.Context) {

	response := models.Response{}
	params := models.GetTask{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.ApiMessage = "validation " + mess
		c.JSON(400, response)
	} else {

		data, err := u.daos.TaskGet(params)

		if err != nil {
			response.ApiStatus = 0
			response.ApiMessage = err.Error()
			c.JSON(500, response)
		} else {
			response.ApiStatus = 1
			response.Data = data
			response.ApiMessage = common.StatusSukses
			c.JSON(http.StatusOK, response)

		}

	}

}

func (u *Merk) TaskUpdate(c *gin.Context) {

	params := models.UpdateTask{}

	response := models.Response{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.ApiMessage = "validation " + mess
		c.JSON(400, response)
	} else {

		data, err := u.daos.TaskUpdate(params)

		if err != nil {
			response.ApiStatus = 0
			response.ApiMessage = err.Error()
			c.JSON(500, response)
		} else {
			response.ApiStatus = 1
			response.Data = data
			response.ApiMessage = common.StatusSukses
			c.JSON(http.StatusOK, response)

		}

	}

}

//func (u *Merk) MerkDelete(c *gin.Context) {
//
//	params := models.DeleteMerk{}
//
//	response := models.Response{}
//
//	err := c.ShouldBind(&params)
//
//	if err != nil {
//		var mess string
//		if err != nil {
//			mess = mess + err.Error()
//		}
//
//		response.ApiMessage = "validation " + mess
//		c.JSON(400, response)
//	} else {
//
//		data, err := u.daos.MerkDelete(params)
//
//		if err != nil {
//			response.ApiStatus = 0
//			response.ApiMessage = err.Error()
//			c.JSON(500, response)
//		} else {
//			response.ApiStatus = 1
//			response.Data = data
//			response.ApiMessage = common.StatusSukses
//			c.JSON(http.StatusOK, response)
//
//		}
//
//	}
//
//}
