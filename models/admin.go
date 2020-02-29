package models

import (
	"github.com/abenbyy/go-backend/database"
	"time"
)

type Admin struct {
	Id		int		`gorm:primary_key`
	Username string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

}

func init(){
	db, err := database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()

	db.DropTableIfExists(&Admin{})
	db.AutoMigrate(&Admin{})

	db.Create(&Admin{
		Username:  "abenbyy",
		Password:  "asd123",
	})
}

func GetAdmin(username string, password string)(Admin, error){

	db, err:= database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	var admin Admin

	db.Where("username = ? AND password = ?",username, password).First(&admin)

	return admin, nil
}