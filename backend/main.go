package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	db "rlesjak.com/ha-scheduler/database/generated"
)

func main() {

	if err := godotenv.Load(".env.local"); err != nil {
		log.Println("FAILED TO LOAD ENV FILE !! \n", err)
		return
	}

	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dataSourceName := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable", host, username, databaseName, port)

	log.Println("DATASOURCE NAME: ")
	log.Println(dataSourceName)
	log.Println("---")

	mainContext := context.Background()

	database, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Println("ERROR CONNECTING TO DATABASE \n", err)
		return
	}

	if err := database.Ping(); err != nil {
		log.Println("ERROR PINGING DB \n", err)
		return
	}

	query := db.New(database)

	r := gin.Default()
	r.GET("hello", func(ctx *gin.Context) {
		// uid, err := uuid.Parse(ctx.Param("uuid"))

		// if err != nil {
		// 	ctx.JSON(500, err)
		// 	return
		// }

		groups, err := query.GetMasterElementGroups(mainContext)

		if err != nil {
			ctx.JSON(500, err)
			return
		}
		ctx.JSON(200, groups)
	})

	r.POST("create", func(ctx *gin.Context) {
		var reqBody db.CreateChildElementsGroupParams
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			ctx.JSON(500, err)
			return
		}

		if err := query.CreateChildElementsGroup(mainContext, reqBody); err != nil {
			ctx.JSON(500, err)
			return
		}
	})

	r.DELETE("delete/:uuid", func(ctx *gin.Context) {

		uid, err := uuid.Parse(ctx.Param("uuid"))
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		if err := query.DeleteElementsGroup(mainContext, uid); err != nil {
			ctx.JSON(500, err)
			return
		}
	})

	r.Run("0.0.0.0:9090")
}
