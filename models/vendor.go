package models

import (
	"github.com/abenbyy/go-backend/database"
	"time"
)

type Vendor struct{
	Id			int			`gorm:primary_key`
	Region		string		`gorm:"type:varchar(100)"`
	Name		string		`gorm:"type:varchar(100)"`
	//RentLocations

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

	db.AutoMigrate(&Vendor{})

}
