package main

import (
    "net/http"
    "log"

	"github.com/gin-gonic/gin"
)

func displayLogin(c *gin.Context) {
    c.HTML(http.StatusOK, "login.html", gin.H{})
}

func handleLogin(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    if isUserValid(username, password) {
        // generate token
        signedToken, err := generateSignedJWTToken(username)

        if err != nil {
            c.HTML(http.StatusInternalServerError, "login.html", gin.H{
                "title": "Login",
                "Error": gin.H{
                    "Title": "Login Failed",
                    "Message": "Internal server error, please wait and try again.",
                },
            })
            log.Println(err.Error())
            return
        }

        // func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
        c.SetCookie("AccessToken", signedToken, 60, "", "", false, true)

        c.HTML(http.StatusOK, "login-successful.html", gin.H{
            "is_logged_in": true,
            "title": "Success",
        })

    } else {
        // redirect to home page
        c.HTML(http.StatusBadRequest, "login.html", gin.H{
            "title": "Login",
            "Error": gin.H{
                "Title": "Login Failed",
                "Message": "Invalid credentials provided",
            },
        })
    }
}

func handleLogout(c *gin.Context) {
}
