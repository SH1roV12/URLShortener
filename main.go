package main

import (
	"urlshortener/db"
	"urlshortener/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db:=db.InitDB()
	r:=gin.Default()
	api:=r.Group("api")
	api.POST("/create",func(c *gin.Context){
		handlers.CreateLink(c,db)
	})
	api.GET("/getallurl", func(c* gin.Context){
		handlers.GetLinks(c,db)
	})
	r.GET("/:short",func(c *gin.Context){
		handlers.RedirectFromShortURL(c,db)
	})
	r.Run(":8080")
}