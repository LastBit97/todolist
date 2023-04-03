package controller

import (
	"net/http"
	"strconv"

	"github.com/LastBit97/todolist/ent"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"

	"github.com/LastBit97/todolist/service"
)

func UserGetAllController(ctx *gin.Context) {
	transaction := ctx.MustGet("transaction").(*sentry.Span)

	span := transaction.StartChild("db")
	users, err := service.NewUserOps(ctx).UsersGetAll()
	span.Finish()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": users})
	transaction.Status = 1
}

func UserGetByIDController(ctx *gin.Context) {
	transaction := ctx.MustGet("transaction").(*sentry.Span)

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	span := transaction.StartChild("db")
	user, err := service.NewUserOps(ctx).UserGetByID(id)
	span.Finish()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
	transaction.Status = 1
}

func UserCreateController(ctx *gin.Context) {
	transaction := ctx.MustGet("transaction").(*sentry.Span)

	var newUser ent.User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	span := transaction.StartChild("db")
	user, err := service.NewUserOps(ctx).UserCreate(newUser)
	span.Finish()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": user})
	transaction.Status = 1
}

func UserUpdateController(ctx *gin.Context) {
	transaction := ctx.MustGet("transaction").(*sentry.Span)

	var newUserData ent.User
	if err := ctx.ShouldBindJSON(&newUserData); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	newUserData.ID = id

	span := transaction.StartChild("db")
	updatedUser, err := service.NewUserOps(ctx).UserUpdate(newUserData)
	span.Finish()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedUser})
	transaction.Status = 1
}

func UserDeleteController(ctx *gin.Context) {
	transaction := ctx.MustGet("transaction").(*sentry.Span)

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	span := transaction.StartChild("db")
	deletedID, err := service.NewUserOps(ctx).UserDelete(id)
	span.Finish()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, deletedID)
	transaction.Status = 1
}
