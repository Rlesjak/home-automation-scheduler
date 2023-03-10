package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

	{
		// /api/v1/elGroup/child**
		childRouter := elGroupsRouter.Group("/child")
		getChildElementGroups(childRouter)
		postChildElementGroups(childRouter)
	}
}

// DELETE /api/v1/elGroup
func deleteElementGroups(router *gin.RouterGroup) {
	router.DELETE("/:uuid", func(ctx *gin.Context) {

		uid, err := uuid.Parse(ctx.Param("uuid"))
		if err != nil {
			abortWithBadRequest(ctx, err)
			return
		}

		if err := models.Q.DeleteElementsGroup(ctx, uid); err != nil {
			abortWithGenericError(ctx, err)
			return
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

		uuid, err := models.Q.CreateMasterElementsGroup(ctx, reqBody)

		if err != nil {
			abortWithGenericError(ctx, err)
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"uuid": uuid})
	})
}

// GET /api/v1/elGroup/child/:uuid
func getChildElementGroups(router *gin.RouterGroup) {
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

		childGroups, err := models.Q.GetChildElementGroupsOf(ctx, uid)
		if err != nil {
			abortWithGenericError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, childGroups)
	})
}

// POST /api/v1/elGroup/child
func postChildElementGroups(router *gin.RouterGroup) {
	router.POST("", func(ctx *gin.Context) {

		reqBody := db.CreateChildElementsGroupParams{}

		binding.EnableDecoderDisallowUnknownFields = true
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			abortWithBadRequest(ctx, err)
			return
		}

		if reqBody.Name == "" {
			abortWithMessage(ctx, http.StatusBadRequest, "Missing key 'name'")
			return
		}

		uuid, err := models.Q.CreateChildElementsGroup(ctx, reqBody)
		if err != nil {
			abortWithGenericError(ctx, err)
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"uuid": uuid})
	})
}
