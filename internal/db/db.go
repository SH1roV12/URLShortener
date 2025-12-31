package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct{
	*gorm.DB
}


// Инициализация БД
func InitDB(dsn string) *DB{
	
	
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	} )

	if err != nil{
		panic(err)
	}

	if err := goose.SetDialect("mysql"); err != nil{
		panic(err)
	}


	sqlDB,err:=db.DB()
	if err != nil{
		panic(err)
	}
	if err := goose.Up(sqlDB, "migrations"); err != nil{
		panic(err)
	}

	if err != nil{
		panic(err)
	}

	fmt.Println("DB has been started")
	return &DB{db}
}