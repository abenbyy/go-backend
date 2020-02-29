package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func GetAdmin(p graphql.ResolveParams) (i interface{},e error){
	username := p.Args["username"].(string)
	password := p.Args["password"].(string)

	admin, err := models.GetAdmin(username,password)

	return admin,err
}