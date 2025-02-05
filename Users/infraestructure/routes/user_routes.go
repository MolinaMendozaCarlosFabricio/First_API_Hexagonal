package routes

import (
	"example2.com/m/Users/application"
	"example2.com/m/Users/infraestructure"
	"example2.com/m/Users/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine){
	ps := infraestructure.NewUserRepoMySQL()
	ru := application.NewRegisterUser(ps)
	ru_controller := controllers.NewRegisterUserController(*ru)
	gu := application.NewGetUser(ps)
	gu_controller := controllers.NewGetUserController(*gu)
	gau := application.NewGetAllUsers(ps)
	gau_controller := controllers.NewGetAllUsersController(*gau)
	eu := application.NewEditUser(ps)
	eu_controller := controllers.NewEditUserController(*eu)
	du := application.NewDeleteUser(ps)
	du_controller := controllers.NewDeleteUserController(*du)

	users := r.Group("/users")
	{
		users.POST("/", ru_controller.Execute)
		users.GET("/:id", gu_controller.Execute)
		users.GET("/", gau_controller.Execute)
		users.PUT("/:id", eu_controller.Execute)
		users.DELETE("/:id", du_controller.Execute)
	}
}