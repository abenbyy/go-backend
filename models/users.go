package models

import (
	"fmt"
	"github.com/abenbyy/go-backend/database"
	"time"
)

type User struct{
	Id 			int			`gorm: primary_key`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	*time.Time 	`sql:"index"`
	FirstName 	string		`gorm: "type:varchar(100);not null"`
	LastName 	string		`gorm: "type:varchar(100);not null"`
	Password	string		`gorm: "type:varchar(100);not null"`
	Email		string		`gorm: "type:varchar(100);not null"`
	Phone		string		`gorm: "type:varchar(100);not null"`
	Title		string		`gorm: "type:varchar(100);"`
	City		string		`gorm: "type:varchar(100);"`
	Address		string		`gorm: "type:varchar(100);"`
	PostCode	string		`gorm: "type:varchar(100);"`

}

func init(){
	db,err:= database.Connect()
	if err != nil{
		panic(err)
	}
	defer db.Close()
	//db.DropTableIfExists(&User{})
	db.AutoMigrate(&User{})
	//CreateUser("Andree","Benaya","asd123","benayaabyatar@gmail.com","+6281286588074")

}

func GetAllUser() ([]User , error){
	db, err:= database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var users []User

	db.Find(&users)

	return users,nil;
}

func GetUserByEmail(email string)([]User, error){
	db, err:= database.Connect()
	if err!=nil{
		return nil,err
	}
	defer db.Close()

	var user []User

	if db.Where("email = ?",email).Find(&user).RecordNotFound(){
		return nil,nil
	}

	fmt.Println(user)

	return user,err
}

func GetUserByPhone(phone string)([]User, error){
	db, err:= database.Connect()
	if err!=nil{
		return nil,err
	}
	defer db.Close()

	var user []User

	if db.Where("phone = ?",phone).Find(&user).RecordNotFound(){
		return nil,nil
	}

	fmt.Println(user)

	return user,err
}

func GetUserByPhoneOrEmail(arg string) ([]User, error){
	db, err:= database.Connect()

	if err != nil{
		return nil,err
	}
	defer db.Close()

	var user []User

	if db.Where("email = ?",arg).Or("phone = ?", arg).Find(&user).RecordNotFound(){
		return nil,nil
	}

	fmt.Println(user)

	return user,nil
}

func CreateUser(firstname string, lastname string, password string, email string, phone string)(*User, error){
	db, err:= database.Connect()

	if err !=nil{
		return nil,err
	}
	defer db.Close()
	var user = User{

		FirstName: firstname,
		LastName: lastname,
		Password: password,
		Email:    email,
		Phone:    phone,
	}
	if db.NewRecord(user){
		db.Create(&user)
	}

	return &user,nil



}