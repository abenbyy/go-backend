package models

import (
	_"fmt"
	"github.com/abenbyy/go-backend/database"
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
}

func GetAirports()([]Airport, error){
	db, err := database.Connect()

	if err !=nil{
		panic(err)
	}

	defer db.Close()

	var airports []Airport

	db.Find(&airports)

	return airports,nil
}
