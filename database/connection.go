package database

import (
	//"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	DB = Connect()
	return DB
}

func Connect() *gorm.DB {

	// retrieve the url
	//dbURL := os.Getenv("DATABASE_URL")
	url := "postgres://duqqerkdxsmzkb:fa98b952bcaab61ba7fed7f12a4fb1448a8668ccb3a09e8efb3741bd28cbf438@ec2-52-72-99-110.compute-1.amazonaws.com:5432/dahahunspaa3kt"

	DB, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		panic("Se perdió la conexión a la base de datos")

	}
	println("Conexión a la base de datos establecida")

	return DB
	//connection.AutoMigrate(&models.User{})
}
