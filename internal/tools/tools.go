package tools

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

//Генерация уникального ID для БД
func GenerateUniqueID()(string,error){
	godotenv.Load()
	alphabet := os.Getenv("ALPHABET_GEN")
	length := os.Getenv("LENGTH_GEN")
	if alphabet == ""{
		alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	}
	if length == ""{
		length = "10"
	}
	l,err := strconv.Atoi(length)
	if err != nil{
		log.Print(err)
		return "",errors.New("failer convert str to int")
	}
	shortUrl,err := gonanoid.Generate(alphabet,l)
	if err != nil{
		log.Print(err)
		return "",errors.New("failer to create short link")
	}
	return shortUrl,nil
}
