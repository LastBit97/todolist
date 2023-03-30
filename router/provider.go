package router

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(rg *gin.RouterGroup) {
	registerTodoRouter(rg)
	registerUserRouter(rg)
}
