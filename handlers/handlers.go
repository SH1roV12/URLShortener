package handlers

import (
	"database/sql"
	"net/http"
	"urlshortener/tools"

	"github.com/gin-gonic/gin"
)


type URLRequest struct{
	URL string `json:"url"`
}
func CreateLink(c *gin.Context, db *sql.DB){
	var req URLRequest
	if err := c.ShouldBindJSON(&req); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}

	var counter uint64
	err:=db.QueryRow("SELECT COUNT(*) FROM urls").Scan(&counter)
	if err != nil{
		panic(err)
	}

	shortId:=tools.GenerateUniqueID(counter+1)
	
	_,err=db.Exec("INSERT INTO urls (short_url,original_url) VALUES(?,?)",shortId,req.URL)
	if err != nil{
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"original_url":req.URL,"short_url":shortId})
}