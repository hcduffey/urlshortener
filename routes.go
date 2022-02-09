package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func initializeRoutes() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.SetTrustedProxies(nil)

	router.GET("/", getUrlSubmissionPage)
	router.POST("/", processUrl)

	router.GET("/r/:url", urlRedirect)
}
