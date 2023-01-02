package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "rlesjak.com/ha-scheduler/database/generated"
	models "rlesjak.com/ha-scheduler/model"
	"rlesjak.com/ha-scheduler/scheduler"
	"rlesjak.com/ha-scheduler/services"
)

func RegisterTriggersController(router *gin.RouterGroup) {
	elGroupsRouter := router.Group("/trigger")

	postStandaloneTrigger(elGroupsRouter)
	deleteTrigger(elGroupsRouter)
}

// POST /api/v1/trigger
func postStandaloneTrigger(router *gin.RouterGroup) {
	router.POST("", func(ctx *gin.Context) {
		var reqBody db.CreateStandaloneTriggerParams

		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			abortWithBadRequest(ctx, err)
			return
		}

		uid, err := models.Q.CreateStandaloneTrigger(ctx, reqBody)
		if err != nil {
			abortWithGenericError(ctx, err)
			return
		}

		nextRun, schedulerErr := services.CreateScheduledJob(
			reqBody.Condition,
			reqBody.Command,
			uid,
		)

		if schedulerErr != nil {
			abortWithGenericError(ctx, err)
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"uuid": uid, "nextRun": nextRun})
	})
}

// DELETE /api/v1/trigger
func deleteTrigger(router *gin.RouterGroup) {
	router.DELETE(":uuid", func(ctx *gin.Context) {
		uid, err := uuid.Parse(ctx.Param("uuid"))
		if err != nil {
			abortWithBadRequest(ctx, err)
			return
		}

		if err := models.Q.DeleteTrigger(ctx, uid); err != nil {
			abortWithGenericError(ctx, err)
			return
		}

		if err := scheduler.DeleteJob(uid); err != nil {
			abortWithGenericError(ctx, err)
			return
		}

		ctx.Status(http.StatusOK)
	})
}
