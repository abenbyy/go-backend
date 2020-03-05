package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func CreateUser(p graphql.ResolveParams) (i interface{}, e error){
	firstname:=p.Args["firstname"].(string)
	lastname:=p.Args["lastname"].(string)
	password:=p.Args["password"].(string)
	email:=p.Args["email"].(string)
	phone:=p.Args["phone"].(string)

	user, err := models.CreateUser(firstname,lastname,password,email,phone);

	return user,err;
}

func UpdateUser(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)
	fname:= p.Args["firstname"].(string)
	lname:= p.Args["lastname"].(string)
	email:= p.Args["email"].(string)
	phone:= p.Args["phone"].(string)
	city:= p.Args["city"].(string)
	address:= p.Args["address"].(string)
	postcode:= p.Args["postcode"].(string)
	lang:= p.Args["language"].(string)

	u:= models.User{
		FirstName: fname,
		LastName:  lname,
		Email:     email,
		Phone:     phone,
		City:      city,
		Address:   address,
		PostCode:  postcode,
		Language:  lang,
		Currency:  "IDR",
	}

	user:= models.UpdateDetail(id,u)

	return user,nil
}