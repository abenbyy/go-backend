package models

import (
	"github.com/abenbyy/go-backend/database"
	"time"
)

type Room struct{
	Id			int				`gorm:primary_key`
	Name		string			`gorm:varchar(100);not null`
	Image		string
	//Service
	RoomSize	int	 			`gorm:int`
	Price		int				`gorm:int`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time		`sql:index`
}

func init(){
	db,err := database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	db.DropTableIfExists(&Room{})
	//db.AutoMigrate(&Room{})
}