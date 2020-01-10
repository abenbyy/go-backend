package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func GetFlights(p graphql.ResolveParams) (i interface{},e error){
	source :=p.Args["source"].(string)
	destination :=p.Args["destination"].(string)

	flights ,err:= models.GetFlights(source,destination)

	return flights,err
}