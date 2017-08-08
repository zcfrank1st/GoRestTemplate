package service

func SimpleServiceTemplate() string{
    return `package service

import "github.com/gin-gonic/gin"

func GetAll{{ . | Title }}s(context *gin.Context) {

}

func Get{{ . | Title }}(context *gin.Context) {

}
func Add{{ . | Title }}(context *gin.Context) {

}
func Update{{ . | Title }}(context *gin.Context) {

}

func Delete{{ . | Title }}(context *gin.Context) {

}
`
}