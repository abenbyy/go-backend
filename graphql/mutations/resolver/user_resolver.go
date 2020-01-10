package resolver

import(
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