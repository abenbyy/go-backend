package models

import (
	"github.com/abenbyy/go-backend/database"
)

type Train struct {
	Id					int 		`gorm:primary_key`
	Name				string 		`gorm:"type:varchar(100);not null"`
	Class				string		`gorm:"type:varchar(100);not null"`
	Subclass			string		`gorm:"type:varchar(100);not null"`
	Code				string		`gorm:"type:varchar(100);not null"`
	//Image
}

func init(){
	db, err := database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()

	//db.DropTableIfExists(&Train{})

	db.AutoMigrate(&Train{})

	//1
	//db.Create(&Train{
	//
	//	Name:     "Argo Parahyangan",
	//	Class:    "Executive",
	//	Subclass: "J",
	//	Code:	"47",
	//})
	//
	////2
	//db.Create(&Train{
	//
	//	Name:     "Pangandaran",
	//	Class:    "Executive",
	//	Subclass: "A",
	//	Code:	"175",
	//})
	//
	////3
	//db.Create(&Train{
	//
	//	Name:     "Argo Parahyangan",
	//	Class:    "Executive",
	//	Subclass: "J",
	//	Code:	"49",
	//})
	//
	////4
	//db.Create(&Train{
	//
	//	Name:     "Argo Willis",
	//	Class:    "Executive",
	//	Subclass: "H",
	//	Code:	"1",
	//})
	//
	////5
	//db.Create(&Train{
	//
	//	Name:     "Argo Parahyangan Priority",
	//	Class:    "Executive",
	//	Subclass: "I",
	//	Code:	"49P",
	//})
	//
	////6
	//db.Create(&Train{
	//
	//	Name:     "Serayu",
	//	Class:    "Economy",
	//	Subclass: "C",
	//	Code:	"325",
	//})
	//
	////7
	//db.Create(&Train{
	//
	//	Name:     "Majapahit",
	//	Class:    "Economy",
	//	Subclass: "J",
	//	Code:	"252",
	//})
	//
	////8
	//db.Create(&Train{
	//
	//	Name:     "Argo Sindoro",
	//	Class:    "Economy",
	//	Subclass: "P",
	//	Code:	"12",
	//})
	//
	////9
	//db.Create(&Train{
	//
	//	Name:     "Menoreh",
	//	Class:    "Business",
	//	Subclass: "S",
	//	Code:	"264",
	//})
	//
	////10
	//db.Create(&Train{
	//
	//	Name:     "Argo Bromo",
	//	Class:    "Business",
	//	Subclass: "H",
	//	Code:	"6L",
	//})

}

func GetTrains()([]Train, error){
	db, err := database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	var trains []Train

	db.Find(&trains)

	return trains, nil
}

