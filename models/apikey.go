package models

import (
	"errors"
	"github.com/abenbyy/go-backend/database"
	"time"
)

type ApiKey struct {
	Id		int		`gorm:primary_key"`
	Key		string	`gorm:"type:text;unique_index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time	`sql:"index"`
}

func init(){
	db, err:= database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	db.DropTableIfExists(&ApiKey{})
	db.AutoMigrate(&ApiKey{})
	SeedApiKey()
}

func ValidateKey(key string)([]ApiKey, error){
	db, err:= database.Connect()
	if err!= nil{
		panic(err)
	}

	defer db.Close()

	var keys []ApiKey
	db.Where("key = ?", key).Find(&keys)

	if len(keys) > 0{
		return keys, nil
	}else {
		return nil, errors.New("Key is not valid")
	}

}

func SeedApiKey(){
	db, err:= database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	//hashbabyhash
	db.Create(&ApiKey{
		Key:       "c91c62f9c5318bb784acdd455a4b41b1",
	})
}