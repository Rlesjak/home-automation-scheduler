package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "rlesjak.com/ha-scheduler/database/generated"
	models "rlesjak.com/ha-scheduler/model"
	"rlesjak.com/ha-scheduler/scheduler"
	"rlesjak.com/ha-scheduler/services"
	"rlesjak.com/ha-scheduler/services/parsers"
)

func RegisterTriggersController(router *gin.RouterGroup) {
	elGroupsRouter := router.Group("/trigger")

	postStandaloneTrigger(elGroupsRouter)
	putTriggerActivity(elGroupsRouter)
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

		// Start database transaction
		tx, txErr := models.DB.Begin()
		if txErr != nil {
			abortWithGenericError(ctx, txErr)
		}

		// Get models inside transaction
		duringTransaction := models.Q.WithTx(tx)

		// Create trigger record, and get uuid of trigger
		uid, err := duringTransaction.CreateStandaloneTrigger(ctx, reqBody)
		if err != nil {
			abortWithGenericError(ctx, err)
			return
		}

		var nextRun string

		if reqBody.Active {
			// Create scheduled job
			var schedulerErr error
			nextRun, schedulerErr = services.CreateScheduledJob(
				reqBody.Condition.String,
				reqBody.Command.String,
				uid,
			)

			// If there was an error scheduling a trigger
			// rollback the transaction
			if schedulerErr != nil {
				tx.Rollback()
				abortWithGenericError(ctx, schedulerErr)
				return
			}
		} else {
			// If the trigger is being deactivated by default
			// Only parse condition and command to check
			// for correctnes
			_, _, parseErr := parsers.ParseJob(reqBody.Condition.String, reqBody.Command.String)
			if parseErr != nil {
				tx.Rollback()
				abortWithBadRequest(ctx, parseErr)
				return
			}
		}

		// Commit transaction
		tx.Commit()
		ctx.JSON(http.StatusCreated, gin.H{"uuid": uid, "nextRun": nextRun})
	})
}

// PATCH /api/v1/trigger/activity
func putTriggerActivity(router *gin.RouterGroup) {
	router.PATCH("/activity", func(ctx *gin.Context) {

		var reqBody db.UpdateTriggerActiveParams

		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			abortWithBadRequest(ctx, err)
			return
		}

		var nextRun string

		if reqBody.Active {
			// Activate trigger (create job)
			// Create scheduled job
			var createJobErr error
			nextRun, createJobErr = services.CreateScheduledJobFromTriggerUuid(ctx, reqBody.Uuid)
			if createJobErr != nil {
				abortWithGenericError(ctx, createJobErr)
				return
			}

			if err := models.Q.UpdateTriggerActive(ctx, reqBody); err != nil {
				abortWithGenericError(ctx, err)
				return
			}

		} else {
			// Deactivate trigger (delete job)
			// Error does not have to be handlede here
			// if deletion failed it means job already does not exist
			scheduler.DeleteJob(reqBody.Uuid)
			if err := models.Q.UpdateTriggerActive(ctx, reqBody); err != nil {
				abortWithGenericError(ctx, err)
				return
			}
		}

		ctx.JSON(http.StatusCreated, gin.H{"nextRun": nextRun})
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
