package handlers

import (
	"database/sql"
	"net/http"
	"urlshortener/tools"

	"github.com/gin-gonic/gin"
)

// Структура входящей ссылки
type URLRequest struct{
	URL string `json:"url"`
}

//Структура данных о ссылке из БД
type URL struct{
	ShortURL string `json:"shorturl"`
	OriginalURL string `json:"originalurl"`
	Date string `json:"date"`
}

//Создание короткой ссылки из большой входящей
func CreateLink(c *gin.Context, db *sql.DB){
	var req URLRequest
	//Обработка входящей ссылки в формате JSON
	if err := c.ShouldBindJSON(&req); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}

	var counter uint64
	//Подсчет сколько строк уже есть в БД для создания уникального ID
	err:=db.QueryRow("SELECT COUNT(*) FROM urls").Scan(&counter)
	if err != nil{
		panic(err)
	}
	
	//Генерация короткой ссылки
	shortId:=tools.GenerateUniqueID(counter+1)
	
	//Вставка данных в БД(короткая ссылка и оригинальная входящая)
	_,err=db.Exec("INSERT INTO urls (short_url,original_url) VALUES(?,?)",shortId,req.URL)
	if err != nil{
		panic(err)
	}
	// Ответ пользователю о добавлении ссылки
	c.JSON(http.StatusOK, gin.H{"original_url":req.URL,"short_url":shortId})
}


//Получение всех данных о сохраненных ссылках 
func GetLinks(c *gin.Context, db *sql.DB){
	
	// Получение всех строк из БД в формате sql.Rows
	rows,err:=db.Query("SELECT original_url,short_url,data FROM urls")
	if err != nil{
		panic(err)
	}
	var data []URL
	for rows.Next(){
		var u URL
		// Обработка столбца в переменную
		err:=rows.Scan(&u.OriginalURL,&u.ShortURL,&u.Date)
		if err != nil{
			panic(err)
		}
		//Добавление ссылки в переменную для вывода пользователю
		data = append(data, u)
	}

	c.JSON(200,data)
}


//Редирект на ссылку по короткой ссылке
func RedirectFromShortURL(c *gin.Context, db *sql.DB){
	//Поиск по какому ID искать оригинальную ссылку в БД
	shortUrl := c.Param("short")
	var originalURL string
	//Поиск оригинальной ссылки по короткой ссылке
	err:=db.QueryRow("SELECT original_url FROM urls WHERE short_url = ?",shortUrl).Scan(&originalURL)
	if err != nil{
		panic(err)
	}
	//Редирект на оригинальную ссылку
	c.Redirect(http.StatusFound,originalURL)
}