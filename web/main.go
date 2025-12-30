package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/Wsine/feishu2md/utils"
	"github.com/gin-gonic/gin"
)

//go:embed templ/*
var f embed.FS

func main() {
	if mode := os.Getenv("GIN_MODE"); mode != "release" {
		utils.LoadEnv()
	}

	router := gin.New()
	templ := template.Must(template.New("").ParseFS(f, "templ/*.templ.html"))
	router.SetHTMLTemplate(templ)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.templ.html", nil)
	})
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	router.GET("/download", downloadHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run("0.0.0.0:" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
