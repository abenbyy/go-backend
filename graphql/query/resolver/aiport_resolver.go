package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func GetAirports(p graphql.ResolveParams) (i interface{},e error){
	airports ,err:= models.GetAirports()

	return airports,err
}