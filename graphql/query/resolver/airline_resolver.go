package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func GetAirlines(p graphql.ResolveParams) (i interface{},e error){
	airlines ,err:= models.GetAirlines()

	return airlines,err
}