package controller

import (
	"net/http"
	"strconv"

	"github.com/LastBit97/todolist/ent"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"

	"github.com/LastBit97/todolist/service"
)

func TodoGetByIDController(ctx *gin.Context) {
	transaction := ctx.MustGet("transaction").(*sentry.Span)

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	span := transaction.StartChild("db")
	todo, err := service.NewTodoOps(ctx).TodoGetByID(id)
	span.Finish()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": todo})
	transaction.Status = 1
}

func TodoCreateController(ctx *gin.Context) {
	transaction := ctx.MustGet("transaction").(*sentry.Span)

	var newTodo ent.Todo

	if err := ctx.ShouldBindJSON(&newTodo); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	span := transaction.StartChild("db")
	createdTodo, err := service.NewTodoOps(ctx).TodoCreate(newTodo)
	span.Finish()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": createdTodo})
	transaction.Status = 1
}

func TodoUpdateController(ctx *gin.Context) {
	transaction := ctx.MustGet("transaction").(*sentry.Span)

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

	span := transaction.StartChild("db")
	updatedTodo, err := service.NewTodoOps(ctx).TodoUpdate(newTodoData)
	span.Finish()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedTodo})
	transaction.Status = 1
}

func TodoDeleteController(ctx *gin.Context) {
	transaction := ctx.MustGet("transaction").(*sentry.Span)

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	span := transaction.StartChild("db")
	deletedID, err := service.NewTodoOps(ctx).TodoDelete(id)
	span.Finish()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, deletedID)
	transaction.Status = 1
}
