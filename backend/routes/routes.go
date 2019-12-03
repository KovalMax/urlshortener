package routes

import (
    "os"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

    "github.com/KovalMax/urlshortener/controller"
)

func Handle() {
    port := os.Getenv("APP_PORT")

    app := echo.New()
    app.Use(middleware.Logger())

    app.GET("/:alias", controller.GetLink)
    app.POST("/links", controller.CreateLink)

    app.Logger.Fatal(app.Start(port))
}
