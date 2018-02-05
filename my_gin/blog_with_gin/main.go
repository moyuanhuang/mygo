package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
    router = gin.Default()
    router.LoadHTMLGlob("templates/*")

    initializeRoutes()
    setSecretKey()

    router.Run()
}

func homePageGET(c *gin.Context) {
    articles := getAllArticles()

    c.HTML(http.StatusOK, "index.html", gin.H{
        "title": "Home Page",
        "payload": articles,
        "is_logged_in": getIsLoggedIn(c),
    })
}
