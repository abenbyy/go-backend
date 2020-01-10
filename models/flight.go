package models

import (
	"github.com/abenbyy/go-backend/database"
	"time"
)

type Flight struct {
	Id				int			`gorm:primary_key`
	Airline 		Airline		`gorm:"foreignkey:AirlineRefer"`
	AirlineRefer	uint
	Routes			[]Route
	From			Airport		`gorm:"foreignkey:FromRefer"`
	FromRefer		uint
	To				Airport		`gorm:"foreignkey:ToRefer"`
	ToRefer			uint
	Departure		time.Time
	Arrival 		time.Time
	Duration		uint
	Price			int				`gorm:"type:int"`
	Tax				int				`gorm:"type:int"`
	ServiceCharge	int				`gorm:"type:int"`
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

	db.DropTableIfExists(&Flight{})

	db.AutoMigrate(&Flight{}).AddForeignKey("airline_refer","airlines(Id)","CASCADE","CASCADE").AddForeignKey("from_refer","airports(Id)","CASCADE","CASCADE").AddForeignKey("to_refer","airports(Id)","CASCADE","CASCADE")


	db.Create(&Flight{

		AirlineRefer:  2,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,5,30,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,8,0,0,0,time.Local),
		Duration:      150,
		Price:         790000,
		Tax:           10000,
		ServiceCharge: 0,
	})
	db.Create(&Flight{

		AirlineRefer:  2,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,6,30,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,9,30,0,0,time.Local),
		Duration:      180,
		Price:         1200000,
		Tax:           0,
		ServiceCharge: 0,
	})

	db.Create(&Flight{

		AirlineRefer:  2,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,9,30,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,12,0,0,0,time.Local),
		Duration:      150,
		Price:         990000,
		Tax:           0,
		ServiceCharge: 10000,
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,10,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,12,45,0,0,time.Local),
		Duration:      165,
		Price:         1500000,
		Tax:           0,
		ServiceCharge: 0,
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,21,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,23,0,0,0,time.Local),
		Duration:      120,
		Price:         2000000,
		Tax:           0,
		ServiceCharge: 0,
	})

	//var flight Flight
	//
	//var flights []Flight
	//
	//
	//db.First(&flight).Related(&flight.Airline,"airline_refer")
	//
	//
	//
	//
	//
	//
	//db.Find(&flights)
	//for i, _ := range flights {
	//	db.Model(flights[i]).Related(&flights[i].Airline,"airline_refer")
	//}
	//
	//
	//fmt.Println(flight.Airline)
	//fmt.Println(flights)

	//res,err2 := GetFlights("Jakarta","Singapore")
	//if err2 != nil{
	//	panic(err2)
	//}
	//
	//fmt.Println(res)

}

func GetFlights(source string, destination string)([]Flight,error){
	db, err := database.Connect()
	if err != nil{
		panic(err)
	}

	defer db.Close()

	var flights []Flight

	db.Where("from_refer IN (?) AND to_refer IN (?)", db.Table("airports").Select("Id").Where("city = ?",source).SubQuery(), db.Table("airports").Select("Id").Where("city = ?",destination).SubQuery()).Find(&flights)

	for i,_ := range flights{
		db.Model(flights[i]).Related(&flights[i].Airline,"airline_refer").Related(&flights[i].From,"from_refer").Related(&flights[i].To,"to_refer")
	}
	
	return flights,nil

}