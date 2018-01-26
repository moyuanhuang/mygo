package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func main() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        // write the given string in the context response body
        // String(code int, format string, values ...[]interface) { }
        myName := "Moyuan Huang"
        c.String(http.StatusOK, "This is the response from %s", myName);
    })

    router.Run(":8080")
}

// TODO: https://www.jianshu.com/p/a31e4ee25305
// TODO: https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin
