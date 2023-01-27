package ui

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/config/v2"
)

var (
	//go:embed dist
	dist embed.FS
)

func RegisterEmbeddedUiRoutes(router *gin.Engine, config config.Config) error {
	fs, err := getHttpFileSystem(dist, "dist")
	if err != nil {
		return err
	}
	router.StaticFS("app", fs)

	return nil
}

// Enter a subdirectory in filesystem and return it as http filesystem
func getHttpFileSystem(embeddedFS embed.FS, subd string) (http.FileSystem, error) {
	dir, err := fs.Sub(embeddedFS, subd)
	if err != nil {
		return nil, err
	}

	return http.FS(dir), nil
}
