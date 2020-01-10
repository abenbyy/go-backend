package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func GetUsers(p graphql.ResolveParams) (i interface{},e error){
	users,err:= models.GetAllUser()

	return users,err
}

func GetUserByEmail(p graphql.ResolveParams) (i interface{}, e error){
	email:=p.Args["email"].(string)
	user,err:= models.GetUserByEmail(email)

	//if err != nil{
	//	return nil,err
	//}

	return user,err

}

func GetUserByPhone(p graphql.ResolveParams) (i interface{}, e error){
	phone:=p.Args["phone"].(string)
	user,err:= models.GetUserByPhone(phone)

	//if err != nil{
	//	return nil,err
	//}
	return user,err

}

func GetUserByPhoneOrEmail(p graphql.ResolveParams) (i interface{},e error){
	arg:= p.Args["arg"].(string)

	user, err:= models.GetUserByPhoneOrEmail(arg)

	return user, err
}