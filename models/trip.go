package models

import (
	"fmt"
	"github.com/abenbyy/go-backend/database"
	"time"
)

type Trip struct{
	Id					int 		`gorm:primary_key`
	Train				Train		`gorm:"foreignkey:TrainRefer"`
	TrainRefer  		int
	From				Station		`gorm:"foreignkey:FromRefer"`
	FromRefer			int
	To					Station		`gorm:"foreignkey:ToRefer"`
	ToRefer				int
	Departure			time.Time
	Arrival				time.Time
	Duration			int
	Price				int
	Tax					int
	Service				int

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

	SeedTripData()

}

func GetAllTrips()([]Trip){

	db, err:= database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()

	var trips []Trip

	db.Find(&trips)

	for i,_ :=range trips{
		db.Model(trips[i]).Related(&trips[i].Train,"train_refer").Related(&trips[i].From,"from_refer").Related(&trips[i].To,"to_refer")
	}

	return trips
}
func GetTrip(id int)(Trip){
	db, err:= database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()

	var trip Trip

	db.Where("id = ?",id).First(&trip)

	db.Model(trip).Related(&trip.Train,"train_refer").Related(&trip.From,"from_refer").Related(&trip.To,"to_refer")

	return trip

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


func CreateTrip(t Trip){
	db, err:= database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	db.Create(&t)


}

func UpdateTrip(id int,t Trip){
	db, err:= database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	db.Model(&Trip{}).Where("id = ?",id).Updates(Trip{
		TrainRefer: t.TrainRefer,
		FromRefer:  t.FromRefer,
		ToRefer:    t.ToRefer,
		Duration:   t.Duration,
		Price:      t.Price,

	})

}
func DeleteTrip(id int){
	db, err:= database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	db.Where("id = ?", id).Delete(&Trip{})
}

func SeedTripData(){
	db, err:= database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

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
