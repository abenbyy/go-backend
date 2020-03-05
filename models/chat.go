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

func GetAllChats(id int)([]Chat, error){

	db, err:= database.Connect()
	if err!=nil{
		panic(err)
	}

	defer db.Close()

	_, err = ValidateKey(middleware.ApiKey)

	if err!=nil{
		return nil, err
	}

	var chats []Chat

	db.Where("user1 = ? OR user2 = ?",id,id).Find(&chats)

	return chats, nil

}

func GetChatById(id int)([]Chat, error){
	db, err:= database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()
	_, err = ValidateKey(middleware.ApiKey)

	if err!=nil{
		return nil, err
	}

	var chats []Chat
	db.Where("id = ?", id).Find(&chats)

	return chats, nil
}

func InsertChat(user1 int, user2 int){
	db, err:= database.Connect()
	if err!=nil{
		panic(err)
	}

	defer db.Close()

	_,err = ValidateKey(middleware.ApiKey)

	if err!=nil{
		return
	}

	db.Create(&Chat{
		User1:     user1,
		User2:     user2,
		Content: "",
	})
}

func UpdateChat(id int, content string){
	db, err:= database.Connect()
	if err!=nil{
		panic(err)
	}

	defer db.Close()

	_,err = ValidateKey(middleware.ApiKey)

	if err!=nil{
		return
	}

	db.Model(&Chat{}).Where("id = ?",id).Updates(&Chat{
		Content:   content,
	})

	return
}

