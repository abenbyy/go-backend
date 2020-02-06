package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func GetHotels(p graphql.ResolveParams) (i interface{},e error){
	city :=p.Args["city"].(string)


	hotels ,err:= models.GetHotels(city)

	return hotels,err
}