package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func setLoginStatus(c *gin.Context) {
    token, err := c.Cookie("AccessToken")
    if err == nil && isValidToken(token) {
        c.Set("is_logged_in", true)
    } else {
        c.Set("is_logged_in", false)
    }
}

func ensureLoggedIn(c *gin.Context) {
    if !getIsLoggedIn(c) {
        c.AbortWithStatus(http.StatusUnauthorized)
    }
}

func ensureNotLoggedIn(c *gin.Context) {
    if getIsLoggedIn(c) {
        c.AbortWithStatus(http.StatusUnauthorized)
    }
}

func isValidToken(token string) bool{
    return true
}

func getIsLoggedIn(c *gin.Context) bool {
    isLoggedInInterface, _ := c.Get("is_logged_in")
    isLoggedIn := isLoggedInInterface.(bool)
    return isLoggedIn
}
