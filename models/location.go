package models

import (
	"github.com/abenbyy/go-backend/database"
	"time"
)

type Location struct{
	Id			int			`gorm:primary_key`
	Longitude	float32
	Latitude	float32
	City		string
	Province	string

	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time	`sql:index`
}

func init(){
	db, err := database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&Location{})

}