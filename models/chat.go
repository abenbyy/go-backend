package models

import (
	"github.com/abenbyy/go-backend/database"
	"github.com/abenbyy/go-backend/middleware"
	"time"
)

type Chat struct{
	Id			int		`gorm:primary_key`
	User1		int
	User2		int
	Content		string	`gorm:"type:text;"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time	`sql:index`
}

func init(){
	db, err:= database.Connect()

	if err!=nil{
		panic(err)
	}

	db.AutoMigrate(&Chat{})

}

func GetAllChats()([]Chat, error){

	db, err:= database.Connect()
	if err!=nil{
		panic(err)
	}

	defer db.Close()

	_, err = ValidateKey(middleware.ApiKey)

	if err!=nil{
		return nil, err
	}

	return nil, nil

}