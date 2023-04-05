package controller

import (
	"net/http"
	"strconv"

	"github.com/LastBit97/todolist/ent"
	"github.com/gin-gonic/gin"

	"github.com/LastBit97/todolist/service"
)

func TodoGetByIDController(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	todo, err := service.NewTodoOps(ctx).TodoGetByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": todo})
}

func TodoCreateController(ctx *gin.Context) {
	var newTodo ent.Todo

	if err := ctx.ShouldBindJSON(&newTodo); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	createdTodo, err := service.NewTodoOps(ctx).TodoCreate(newTodo)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": createdTodo})
}

func TodoUpdateController(ctx *gin.Context) {
	var newTodoData ent.Todo
	if err := ctx.ShouldBindJSON(&newTodoData); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	todoIdStr := ctx.Param("id")
	todoId, err := strconv.Atoi(todoIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	newTodoData.ID = todoId

	updatedTodo, err := service.NewTodoOps(ctx).TodoUpdate(newTodoData)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedTodo})
}

func TodoDeleteController(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	deletedID, err := service.NewTodoOps(ctx).TodoDelete(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, deletedID)
}
