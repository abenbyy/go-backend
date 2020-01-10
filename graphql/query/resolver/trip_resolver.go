package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func GetTrips(p graphql.ResolveParams) (i interface{},e error){
	source :=p.Args["source"].(string)
	destination :=p.Args["destination"].(string)

	trips ,err:= models.GetTrips(source,destination)

	return trips,err
}