package models

import (
	"github.com/abenbyy/go-backend/database"
	"time"
)

type Airline struct{
	Id			int				`gorm:primary_key`
	//Image
	Name		string			`gorm:varchar(100);not null`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time		`sql:index`
}

func init(){
	db, err := database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&Airline{})

}

func GetAirlines()([]Airline, error){
	db, err := database.Connect()

	if err !=nil{
		panic(err)
	}

	defer db.Close()

	var airlines []Airline

	db.Find(&airlines)

	return airlines, nil
}