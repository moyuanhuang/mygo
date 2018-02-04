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
    })
}

func displayArticle(c *gin.Context) {
    article, err := getArticleByID(c.Param("id"))
    if err != nil {
        // difference: err will be printed to console
        c.AbortWithError(http.StatusNotFound, err)
        // c.AbortWithStatus(http.StatusNotFound)
    } else {
        c.HTML(http.StatusOK, "article.html", gin.H{
            "title": article.Title,
            "Content": article.Content,
        })
    }
}
