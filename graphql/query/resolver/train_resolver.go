package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func GetTrains(p graphql.ResolveParams) (i interface{},e error){
	trains ,err:= models.GetTrains()

	return trains,err
}