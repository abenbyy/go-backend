package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func GetCars(p graphql.ResolveParams) (i interface{},e error){
	city :=p.Args["city"].(string)


	cars ,err:= models.GetCars(city)

	return cars,err
}
