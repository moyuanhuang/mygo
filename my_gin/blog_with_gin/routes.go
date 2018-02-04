package main

func initializeRoutes() {
    router.GET("/", homePageGET)
    router.GET("/article/view/:id", displayArticle)

    router.GET("/user/login", displayLogin)
    router.POST("/user/login", handleLogin)
    router.GET("/user/logout", handleLogout)
}
