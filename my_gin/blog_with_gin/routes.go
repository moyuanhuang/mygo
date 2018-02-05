package main

func initializeRoutes() {
    router.Use(setLoginStatus)

    router.GET("/", homePageGET)

    userRouter := router.Group("/user")
    {
        userRouter.GET("/login", ensureNotLoggedIn, displayLogin)
        userRouter.POST("/login",  handleLogin)
        userRouter.GET("/logout", ensureLoggedIn, handleLogout)
    }

    articleRouter := router.Group("/article/")
    {
        articleRouter.GET("/view/:id", displayArticle)
    }
}
