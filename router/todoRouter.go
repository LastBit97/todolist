package router

import (
	"github.com/LastBit97/todolist/controller"
	"github.com/gin-gonic/gin"
)

func registerTodoRouter(rg *gin.RouterGroup) {

	router := rg.Group("/todo")
	router.GET("/:id", controller.TodoGetByIDController)
	router.POST("/", controller.TodoCreateController)
	router.PUT("/:id", controller.TodoUpdateController)
	router.DELETE("/:id", controller.TodoDeleteController)
}
