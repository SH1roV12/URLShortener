package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


type Config struct{
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct{
	User string
	Password string
	Host string
	Name string
	Port string
}

func LoadConfig()*Config{
	godotenv.Load()
	user:=os.Getenv("DB_USER")
	pass:=os.Getenv("DB_PASS")
	host:=os.Getenv("DB_HOST")
	port:=os.Getenv("DB_PORT")
	name:=os.Getenv("DB_NAME")

	return &Config{
		DatabaseConfig: DatabaseConfig{
			User: user,
			Password: pass,
			Host: host,
			Port: port,
			Name: name,
		},
	}
}

func (db *DatabaseConfig) DSN() string{
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db.User, db.Password, db.Host, db.Port, db.Name)
}