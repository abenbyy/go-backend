package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)


func GetAllTrips(p graphql.ResolveParams)(i interface{}, e error){


	trips:=models.GetAllTrips()

	return trips, nil
}

func GetTrip(p graphql.ResolveParams)(i interface{}, e error){

	id:= p.Args["id"].(int)

	trip:=models.GetTrip(id)

	return trip, nil
}
func GetTrips(p graphql.ResolveParams) (i interface{},e error){
	source :=p.Args["source"].(string)
	destination :=p.Args["destination"].(string)

	trips ,err:= models.GetTrips(source,destination)

	return trips,err
}