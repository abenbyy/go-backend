package models

import (
	"github.com/abenbyy/go-backend/database"
	"time"
)

type Hotel struct{
	Id 				int				`gorm: primary_key`
	Name 			string			`gorm:"type:varchar(100);not null"`
	Location		Location		`gorm:"foreignkey:LocationRefer"`
	LocationRefer	uint
	Address			string			`gorm:"type:varchar(100);not null"`
	Rating			int
	Room			Room			`gorm:"foreignkey:RoomRefer;not null"`
	RoomRefer		uint
	//Facilities
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		*time.Time		`sql:index`
}

func init(){
	db, err:= database.Connect()

	if err != nil{
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Hotel{})

}


