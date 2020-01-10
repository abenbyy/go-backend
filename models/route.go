package models

import (
	"github.com/abenbyy/go-backend/database"
	"time"
)

type Route struct {
	Id				int				`gorm:primary_key`
	From			string			`gorm:varchar(100);not null`
	To				string			`gorm:varchar(100);not null`
	//FlightId
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		*time.Time		`sql:index`
}

func init(){
	db, err := database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&Route{})
}