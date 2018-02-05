package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

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
            "is_logged_in": getIsLoggedIn(c),
        })
    }
}
