package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func getUrlSubmissionPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}

func processUrl(c *gin.Context) {
	shortPath := generateRandomPath()
	shortURL := "http://" + c.Request.Host + "/r/" + shortPath

	fullURL := c.PostForm("url")
	if !strings.HasPrefix(fullURL, "http") {
		fullURL = "https://" + fullURL
	}

	addURL(shortPath, fullURL)
	c.HTML(http.StatusOK, "created.tmpl", gin.H{
		"title": "Generated URL",
		"url":   template.URL(shortURL),
	})
}

func urlRedirect(c *gin.Context) {
	location := c.Request.URL.Path
	c.Redirect(http.StatusFound, retrieveURL(location[len("/r/"):]))
}

// returns a random string of letters between 4 and 10 characters long to use as part of the shortened URL
func generateRandomPath() string {
	var retString string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rand.Intn(6)+4; i++ {
		retString += string(byte(rand.Intn(26) + 97))
	}
	return retString
}
