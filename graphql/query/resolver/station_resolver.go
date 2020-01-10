package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func GetStations(p graphql.ResolveParams) (i interface{},e error){
	stations ,err:= models.GetStations()

	return stations ,err
}