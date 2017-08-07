package bootstrap

func BootstrapTemplate() string {
    return `package main

import (
    "github.com/gin-gonic/gin"
    "service"
)

func main() {
    serv := gin.Default()

    serv.GET("/status", func(c *gin.Context) {
        c.String(200, "OK")
    })

    g := serv.Group("/home")
    {
        slide := g.Group("/slide")
        {
            slide.POST("", service.AddSlide)
            slide.GET("", service.GetAllSlides)
            slide.PUT("/:id", service.UpdateSlide)
            slide.DELETE("/:id", service.DeleteSlide)
        }
        //module := g.Group("/module")
        //{
        //
        //}
        //theme := g.Group("/theme")
        //{
        //
        //}
    }

    serv.Run()
}
`
}