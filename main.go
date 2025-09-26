package main

import (
	"urlshortener/db"
	"urlshortener/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	//Запуск БД
	db:=db.InitDB()
	r:=gin.Default()
	api:=r.Group("api")
	//Ручка создания ссылки
	api.POST("/create",func(c *gin.Context){
		handlers.CreateLink(c,db)
	})
	//Ручка получения всех ссылок 
	api.GET("/getallurl", func(c* gin.Context){
		handlers.GetLinks(c,db)
	})
	//Редирект на оригинальную ссылку
	r.GET("/:short",func(c *gin.Context){
		handlers.RedirectFromShortURL(c,db)
	})
	r.Run(":8080")
}