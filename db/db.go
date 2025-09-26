package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Инициализация БД
func InitDB() *sql.DB{
	// Подгрузка .env переменных
	godotenv.Load()
	user:=os.Getenv("DB_USER")
	pass:=os.Getenv("DB_PASS")
	host:=os.Getenv("DB_HOST")
	port:=os.Getenv("DB_PORT")
	name:=os.Getenv("DB_NAME")
	//Сборка данных для БД
	dbstring:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)
	fmt.Println(dbstring)
	// Открытие БД
	db,err:=sql.Open("mysql",dbstring)
	if err != nil{
		panic(err)
	}
	// Проверка на подключение к БД
	if err := db.Ping(); err != nil{
		panic(err)
	}
	fmt.Println("DB has been started")
	return db
}