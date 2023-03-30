package router

import (
	"github.com/LastBit97/todolist/controller"
	"github.com/gin-gonic/gin"
)

func registerUserRouter(rg *gin.RouterGroup) {
	router := rg.Group("/user")
	router.GET("/", controller.UserGetAllController)
	router.GET("/{id}", controller.UserGetByIDController)
	router.POST("/", controller.UserCreateController)
	router.PUT("/{id}", controller.UserUpdateController)
	router.DELETE("/{id}", controller.UserDeleteController)
}
