package models

import (
	"github.com/abenbyy/go-backend/database"
	"github.com/abenbyy/go-backend/middleware"
	"time"
)

type Station struct {
	Id			int			`gorm:primary_key`
	Code 		string		`gorm:"type:varchar(100);not null"`
	Name		string 		`gorm:"type:varchar(100);not null"`
	City		string 		`gorm:"type:varchar(100);not null"`
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

	//db.DropTableIfExists(&Station{})
	db.AutoMigrate(&Station{})

	////1
	//db.Create(&Station{
	//	Code:      "JAKK",
	//	Name:      "Jakarta Kota",
	//	City:      "Jakarta",
	//})
	//
	////2
	//db.Create(&Station{
	//	Code:      "SMT",
	//	Name:      "Semarangtawang",
	//	City:      "Semarang",
	//})
	//
	////3
	//db.Create(&Station{
	//	Code:      "PSE",
	//	Name:      "Pasar Senen",
	//	City:      "Jakarta",
	//})
	//
	////4
	//db.Create(&Station{
	//	Code:      "GMR",
	//	Name:      "Gambir",
	//	City:      "Jakarta",
	//})
	//
	//
	////5
	//db.Create(&Station{
	//	Code:      "CMI",
	//	Name:      "Cimahi",
	//	City:      "Bandung",
	//})
	//
	////6
	//db.Create(&Station{
	//	Code:      "KAC",
	//	Name:      "Kiaracondong",
	//	City:      "Bandung",
	//})
	//
	////7
	//db.Create(&Station{
	//	Code:      "BD",
	//	Name:      "Bandung",
	//	City:      "Bandung",
	//})

}

func GetStations()([]Station,error){
	db, err := database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	_, err = ValidateKey(middleware.ApiKey)

	if err != nil{
		return nil, err
	}

	var stations []Station

	db.Find(&stations)

	return stations, nil

}