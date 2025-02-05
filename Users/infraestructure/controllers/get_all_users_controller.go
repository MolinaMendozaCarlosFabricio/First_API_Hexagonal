package controllers

import (
	"net/http"

	"example2.com/m/Users/application"
	"github.com/gin-gonic/gin"
)

type GetAllUsersController struct {
	Service application.GetAllUsers
}

func NewGetAllUsersController(uc application.GetAllUsers)*GetAllUsersController{
	return&GetAllUsersController{Service: uc}
}

func(controller *GetAllUsersController)Execute(c *gin.Context){
	result, err := controller.Service.Execute()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Error al obtener todos los usuarios",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Message":"Usuarios obtenidos",
		"Users": result,
	})
}