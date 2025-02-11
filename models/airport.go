package models

import (
	_"fmt"
	"github.com/abenbyy/go-backend/database"
	"github.com/abenbyy/go-backend/middleware"
	"time"
)

type Airport struct{
	Id			int			`gorm:primary_key`
	Code		string		`gorm: "type:char(3);not null"`
	Name		string		`gorm: "type:varchar(100);not null"`
	City		string		`gorm: "type:varchar(100);not null"`
	CityCode	string		`gorm: "type:char(4);not null"`
	Province	string		`gorm: "type:varchar(100);not null"`
	Country		string		`gorm: "type:varchar(100);not null"`
	CreatedAt	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	*time.Time	`sql:index`

}

func init(){
	db,err := database.Connect()

	if err !=nil{
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Airport{})
	//db.Create(&Airport{
	//	Id:        1,
	//	Code:      "CGK",
	//	Name:      "Soekarno-Hatta",
	//	City:      "Jakarta",
	//	CityCode:  "JKTC",
	//	Province:  "DKI Jakarta",
	//	Country:   "Indonesia",
	//
	//})
	//
	//db.Create(&Airport{
	//	Id:        2,
	//	Code:      "CHG",
	//	Name:      "Changi",
	//	City:      "Singapore",
	//	CityCode:  "SGKC",
	//	Province:  "Singapore",
	//	Country:   "Singapore",
	//
	//})
	//db.Create(&Airport{
	//	Id:        3,
	//	Code:      "PLM",
	//	Name:      "Palembang",
	//	City:      "Palembang",
	//	CityCode:  "PLM",
	//	Province:  "Sumatera Selatan",
	//	Country:   "Indonesia",
	//
	//})
}

func GetAirports()([]Airport, error){
	db, err := database.Connect()

	if err !=nil{
		panic(err)
	}

	defer db.Close()

	var airports []Airport

	_, err = ValidateKey(middleware.ApiKey)

	if err != nil{
		return nil, err
	}

	db.Find(&airports)

	return airports,nil
}
