package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	db "rlesjak.com/ha-scheduler/database/generated"
	models "rlesjak.com/ha-scheduler/model"
)

func RegisterTriggerGroupsController(router *gin.RouterGroup) {
	trigGroupsRouter := router.Group("/trigGroup")

	deleteTriggerGroups(trigGroupsRouter)

	{
		// /api/v1/trigGroup/master**
		mastersRouter := trigGroupsRouter.Group("/master")
		getMasterTriggerGroups(mastersRouter)
		postMasterTriggerGroups(mastersRouter)
	}

	{
		// /api/v1/trigGroup/child**
		childRouter := trigGroupsRouter.Group("/child")
		getChildTriggerGroups(childRouter)
		postChildTriggerGroups(childRouter)
	}
}

// DELETE /api/v1/trigGroup
func deleteTriggerGroups(router *gin.RouterGroup) {
	router.DELETE("/:uuid", func(ctx *gin.Context) {

		uid, err := uuid.Parse(ctx.Param("uuid"))
		if err != nil {
			abortWithBadRequest(ctx, err)
			return
		}

		if err := models.Q.DeleteTriggersGroup(ctx, uid); err != nil {
			abortWithGenericError(ctx, err)
			return
		}

		ctx.Status(http.StatusOK)
	})
}

// GET /api/v1/trigGroup/master
func getMasterTriggerGroups(router *gin.RouterGroup) {
	router.GET("", func(ctx *gin.Context) {

		masterGroups, err := models.Q.GetMasterTriggerGroups(ctx)
		if err != nil {
			abortWithGenericError(ctx, err)
			return
		}

		ctx.JSON(200, masterGroups)
	})
}

// POST /api/v1/trigGroup/master
func postMasterTriggerGroups(router *gin.RouterGroup) {
	router.POST("", func(ctx *gin.Context) {

		var reqBody db.CreateMasterTriggersGroupParams

		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			abortWithBadRequest(ctx, err)
			return
		}

		uuid, err := models.Q.CreateMasterTriggersGroup(ctx, reqBody)

		if err != nil {
			abortWithGenericError(ctx, err)
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"uuid": uuid})
	})
}

// GET /api/v1/trigGroup/child/:uuid
func getChildTriggerGroups(router *gin.RouterGroup) {
	router.GET("", func(ctx *gin.Context) {
		abortWithMessage(ctx, http.StatusBadRequest,
			"Missing uuid of master group. Use: .../child/:uuid",
		)
	})

	router.GET(":uuid", func(ctx *gin.Context) {

		uid, err := uuid.Parse(ctx.Param("uuid"))
		if err != nil {
			abortWithBadRequest(ctx, err)
			return
		}

		childGroups, err := models.Q.GetChildTriggerGroupsOf(ctx, uid)
		if err != nil {
			abortWithGenericError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, childGroups)
	})
}

// POST /api/v1/trigGroup/child
func postChildTriggerGroups(router *gin.RouterGroup) {
	router.POST("", func(ctx *gin.Context) {

		reqBody := db.CreateChildTriggersGroupParams{}

		binding.EnableDecoderDisallowUnknownFields = true
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			abortWithBadRequest(ctx, err)
			return
		}

		if reqBody.Name == "" {
			abortWithMessage(ctx, http.StatusBadRequest, "Missing key 'name'")
			return
		}

		uuid, err := models.Q.CreateChildTriggersGroup(ctx, reqBody)
		if err != nil {
			abortWithGenericError(ctx, err)
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"uuid": uuid})
	})
}
