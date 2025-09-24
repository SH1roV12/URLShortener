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
type URL struct{
	ShortURL string `json:"shorturl"`
	OriginalURL string `json:"originalurl"`
	Date string `json:"date"`
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


func GetLinks(c *gin.Context, db *sql.DB){
	
	rows,err:=db.Query("SELECT original_url,short_url,data FROM urls")
	if err != nil{
		panic(err)
	}
	var data []URL
	for rows.Next(){
		var u URL
		err:=rows.Scan(&u.OriginalURL,&u.ShortURL,&u.Date)
		if err != nil{
			panic(err)
		}
		data = append(data, u)
	}

	c.JSON(200,data)
}

func RedirectFromShortURL(c *gin.Context, db *sql.DB){
	shortUrl := c.Param("short")
	var originalURL string
	err:=db.QueryRow("SELECT original_url FROM urls WHERE short_url = ?",shortUrl).Scan(&originalURL)
	if err != nil{
		panic(err)
	}
	c.Redirect(http.StatusFound,originalURL)
}