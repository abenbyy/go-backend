package models

import (
	"github.com/abenbyy/go-backend/database"
	"time"
)

type Car struct {
	Id			int				`gorm:primary_key`
	Brand		string
	Model		string
	Passenger	int
	Luggage		int
	Vendors		[]CarVendor		`gorm:"foreignkey:CarVendorRefer"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time		`sql:index`
}

type CarVendor struct{
	Id 						int		`gorm:primary_key`
	CarVendorRefer			int
	Name					string
	City					string
	Price					int
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time		`sql:index`
}


func init(){
	db,err := database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()

	db.DropTableIfExists(&CarVendor{})
	db.DropTableIfExists(&Car{})

	db.AutoMigrate(&Car{})
	db.AutoMigrate(&CarVendor{})

	db.Model(&CarVendor{}).AddForeignKey("car_vendor_refer","cars(id)","CASCADE","CASCADE")


	SeedCarData()


}

func GetCars(city string)([]Car,error){

	db,err := database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()


	var res []Car

	db.Preload("Vendors").Find(&res)


	var cars []Car

	for i := range(res){
		for j := range (res[i].Vendors){
			if res[i].Vendors[j].City == city{
				cars = append(cars, res[i])
				break
			}
		}
	}

	return cars, nil
}


func SeedCarData(){

	db, err := database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()

	db.Create(&Car{
		Brand:     "Daihatsu",
		Model:     "Ayla",
		Passenger: 4,
		Luggage:   1,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 1,
		Name:           "JALADIPA TRANS",
		City:           "Jakarta",
		Price:          208000,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 1,
		Name:           "AMJ FAST RENT",
		City:           "Jakarta",
		Price:          227000,
	})


	db.Create(&Car{
		Brand:     "Toyota",
		Model:     "All New Avanza",
		Passenger: 6,
		Luggage:   2,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 2,
		Name:           "SHEYCO TOUR",
		City:           "Jakarta",
		Price:          210000,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 2,
		Name:           "JALADIPA TRANS",
		City:           "Jakarta",
		Price:          243000,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 2,
		Name:           "Abisatya Rent Car",
		City:           "Jakarta",
		Price:          250000,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 2,
		Name:           "AMJ FAST RENT",
		City:           "Jakarta",
		Price:          255000,
	})


	db.Create(&Car{
		Brand:     "Daihatsu",
		Model:     "Sigra",
		Passenger: 6,
		Luggage:   2,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 3,
		Name:           "SHEYCO TOUR",
		City:           "Jakarta",
		Price:          210000,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 3,
		Name:           "AMJ FAST RENT",
		City:           "Jakarta",
		Price:          239000,
	})



	db.Create(&Car{
		Brand:     "Daihatsu",
		Model:     "All New Xenia",
		Passenger: 6,
		Luggage:   2,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 4,
		Name:           "SHEYCO TOUR",
		City:           "Jakarta",
		Price:          220000,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 4,
		Name:           "JALADIPA TRANS",
		City:           "Jakarta",
		Price:          243000,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 4,
		Name:           "Abisatya Rent Car",
		City:           "Jakarta",
		Price:          250000,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 4,
		Name:           "AMJ FAST RENT",
		City:           "Jakarta",
		Price:          255000,
	})

	db.Create(&Car{
		Brand:     "Toyota",
		Model:     "Calya",
		Passenger: 5,
		Luggage:   2,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 5,
		Name:           "AMJ FAST RENT",
		City:           "Jakarta",
		Price:          239000,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 5,
		Name:           "UNICAR",
		City:           "Jakarta",
		Price:          300000,
	})


	db.Create(&Car{
		Brand:     "Nissan",
		Model:     "Grand Livina",
		Passenger: 5,
		Luggage:   2,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 6,
		Name:           "JALADIPA TRANS",
		City:           "Jakarta",
		Price:          268000,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 6,
		Name:           "SHEYCO TOUR",
		City:           "Jakarta",
		Price:          300000,
	})

	db.Create(&Car{
		Brand:     "Suzuki",
		Model:     "Ertiga",
		Passenger: 5,
		Luggage:   2,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 7,
		Name:           "JALADIPA TRANS",
		City:           "Jakarta",
		Price:          268000,
	})

	db.Create(&Car{
		Brand:     "Mitsubishi",
		Model:     "Xpander",
		Passenger: 5,
		Luggage:   2,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 8,
		Name:           "JALADIPA TRANS",
		City:           "Jakarta",
		Price:          388000,
	})

	db.Create(&Car{
		Brand:     "Toyota",
		Model:     "Sienta",
		Passenger: 6,
		Luggage:   2,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 9,
		Name:           "UNICAR",
		City:           "Jakarta",
		Price:          388000,
	})


	db.Create(&Car{
		Brand:     "Toyota",
		Model:     "Grand Innova",
		Passenger: 6,
		Luggage:   2,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 10,
		Name:           "UNICAR",
		City:           "Jakarta",
		Price:          500000,
	})

	db.Create(&Car{
		Brand:     "Toyota",
		Model:     "New Innova Reborn",
		Passenger: 6,
		Luggage:   2,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 11,
		Name:           "UNICAR",
		City:           "Jakarta",
		Price:          700000,
	})

	db.Create(&Car{
		Brand:     "Toyota",
		Model:     "Hiace Commuter",
		Passenger: 16,
		Luggage:   4,
	})

	db.Create(&CarVendor{

		CarVendorRefer: 12,
		Name:           "Obet Rent Car",
		City:           "Jakarta",
		Price:          1200000,
	})




}