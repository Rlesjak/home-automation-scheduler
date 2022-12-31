package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "rlesjak.com/ha-scheduler/database/generated"
	models "rlesjak.com/ha-scheduler/model"
)

func RegisterElementGroupsController(router *gin.RouterGroup) {
	elGroupsRouter := router.Group("/elGroup")

	deleteElementGroups(elGroupsRouter)

	{
		// /api/v1/elGroup/master**
		mastersRouter := elGroupsRouter.Group("/master")
		getMasterElementGroups(mastersRouter)
		postMasterElementGroups(mastersRouter)
	}
}

// DELETE /api/v1/elGroup
func deleteElementGroups(router *gin.RouterGroup) {
	router.DELETE("/:uuid", func(ctx *gin.Context) {

		uid, err := uuid.Parse(ctx.Param("uuid"))
		if err != nil {
			abortWithBadRequest(ctx, err)
		}

		if err := models.Q.DeleteElementsGroup(ctx, uid); err != nil {
			abortWithGenericError(ctx, err)
		}

		ctx.Status(http.StatusOK)
	})
}

// GET /api/v1/elGroup/master
func getMasterElementGroups(router *gin.RouterGroup) {
	router.GET("", func(ctx *gin.Context) {

		masterGroups, err := models.Q.GetMasterElementGroups(ctx)
		if err != nil {
			abortWithGenericError(ctx, err)
			return
		}

		ctx.JSON(200, masterGroups)
	})
}

// POST /api/v1/elGroup/master
func postMasterElementGroups(router *gin.RouterGroup) {
	router.POST("", func(ctx *gin.Context) {

		var reqBody db.CreateMasterElementsGroupParams

		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			abortWithBadRequest(ctx, err)
			return
		}

		if err := models.Q.CreateMasterElementsGroup(ctx, reqBody); err != nil {
			abortWithGenericError(ctx, err)
		}

		ctx.Status(http.StatusCreated)
	})
}
