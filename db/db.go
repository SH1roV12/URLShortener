package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func InitDB() *sql.DB{
	godotenv.Load()
	user:=os.Getenv("DB_USER")
	pass:=os.Getenv("DB_PASS")
	host:=os.Getenv("DB_HOST")
	port:=os.Getenv("DB_PORT")
	name:=os.Getenv("DB_NAME")
	dbstring:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)
	fmt.Println(dbstring)
	db,err:=sql.Open("mysql",dbstring)
	if err != nil{
		panic(err)
	}
	if err := db.Ping(); err != nil{
		panic(err)
	}
	fmt.Println("DB has been started")
	return db
}