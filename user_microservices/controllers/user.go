package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user_microservices/common"
	"user_microservices/daos"
	"user_microservices/models"
)

type User struct {
	daos daos.User
}

func (u *User) UserCreate(c *gin.Context) {

	params := models.CreateUser{}

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

		data, err := u.daos.UserCreate(params)

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

func (u *User) GetDataUser(c *gin.Context) {

	response := models.Response{}
	params := models.GetUser{}

	err := c.ShouldBind(&params)

	if err != nil {
		var mess string
		if err != nil {
			mess = mess + err.Error()
		}

		response.ApiMessage = "validation " + mess
		c.JSON(400, response)
	} else {

		data, err := u.daos.UserGet(params)

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

func (u *User) Signin(c *gin.Context) {

	params := models.UserToken{}
	response := models.Response{}

	err := c.ShouldBind(&params)

	if err != nil {
		response.ApiStatus = 0
		response.ApiMessage = err.Error()
		c.JSON(400, response)
	} else if params.Email == "" && params.Password == "" {
		response.ApiStatus = 0
		response.ApiMessage = "Email Tidak boleh Kosong / Password Tidak boleh Kosong"
		c.JSON(400, response)
	} else {
		dataUser, token, err := u.daos.Signin(params)
		if err != nil {
			response.ApiStatus = 0
			response.ApiMessage = err.Error()
			c.JSON(400, response)
		} else {
			response.ApiStatus = 1
			response.Data = dataUser
			response.ApiMessage = common.StatusSukses
			if token != "" {
				response.Token = token
			}

			c.JSON(http.StatusOK, response)
		}
	}
}
