package models

import (
	"fmt"
	"github.com/abenbyy/go-backend/database"
	"time"
)

type Trip struct{
	Id					int 		`gorm:primary_key`
	Train				Train		`gorm:"foreignkey:TrainRefer"`
	TrainRefer  		uint
	From				Station		`gorm:"foreignkey:FromRefer"`
	FromRefer			uint
	To					Station		`gorm:"foreignkey:ToRefer"`
	ToRefer				uint
	Departure			time.Time
	Arrival				time.Time
	Duration			uint
	Price				uint
	Tax					uint
	Service				uint

	CreatedAt			time.Time
	UpdatedAt			time.Time
	DeletedAt			*time.Time	`sql:index`
}


func init(){
	db, err := database.Connect()
	if err!=nil{
		panic(err)
	}

	defer db.Close()

	db.DropTableIfExists(&Trip{})

	db.AutoMigrate(&Trip{}).AddForeignKey("train_refer","trains(Id)","CASCADE","CASCADE").AddForeignKey("from_refer","stations(Id)","CASCADE","CASCADE").AddForeignKey("to_refer","stations(Id)","CASCADE","CASCADE")


	db.Create(&Trip{
		TrainRefer: 1,
		FromRefer:  4,
		ToRefer:    8,
		Departure:  time.Date(2020,1,21,9,40,00,0,time.Local),
		Arrival:    time.Date(2020,1,21,12,55,00,0,time.Local),
		Duration:	195,
		Price:      110000,
		Tax:        10000,
		Service:    0,
	})

	db.Create(&Trip{
		TrainRefer: 6,
		FromRefer:  3,
		ToRefer:    6,
		Departure:  time.Date(2020,1,21,19,00,00,0,time.Local),
		Arrival:    time.Date(2020,1,21,22,35,00,0,time.Local),
		Duration:	215,
		Price:      120000,
		Tax:        0,
		Service:    0,
	})

	db.Create(&Trip{
		TrainRefer: 8,
		FromRefer:  3,
		ToRefer:    6,
		Departure:  time.Date(2020,1,21,10,40,00,0,time.Local),
		Arrival:    time.Date(2020,1,21,13,40,00,0,time.Local),
		Duration:	300,
		Price:      66000,
		Tax:        10000,
		Service:    0,
	})
}

func GetTrips(source string, destination string)([]Trip, error){
	db, err := database.Connect()
	if err!=nil{
		panic(err)
	}

	defer db.Close()

	var trips []Trip

	//fmt.Println(db.Table("stations").Select("Id").Where("city = ?",source))
	//fmt.Println(db.Table("stations").Select("Id").Where("city = ?",destination))

	db.Where("from_refer IN (?) AND to_refer IN (?)", db.Table("stations").Select("Id").Where("city = ?",source).SubQuery(), db.Table("stations").Select("Id").Where("city = ?",destination).SubQuery()).Find(&trips)


	for i,_ :=range trips{
		db.Model(trips[i]).Related(&trips[i].Train,"train_refer").Related(&trips[i].From,"from_refer").Related(&trips[i].To,"to_refer")
	}

	fmt.Println(trips)

	return trips, nil
}
