package bootstrap

func SimpleBootstrapTemplate() string {
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

    g := serv.Group("/{{.Project | ToLower}}")
    {
        {{ range $index, $value := .Services }}
        {{ $value | ToLower }} := g.Group("/{{ $value | ToLower }}")
        {
            {{ $value | ToLower }}.GET("", service.GetAll{{ $value | Title }}s)
            {{ $value | ToLower }}.GET("/:id", service.Get{{ $value | Title }})
            {{ $value | ToLower }}.POST("", service.Add{{ $value | Title }})
            {{ $value | ToLower }}.PUT("/:id", service.Update{{ $value | Title }})
            {{ $value | ToLower }}.DELETE("/:id", service.Delete{{ $value | Title }})
        }
        {{ end }}
    }

    serv.Run()
}
`
}