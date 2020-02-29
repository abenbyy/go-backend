package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)


func GetAllFlights(p graphql.ResolveParams)(i interface{}, e error){
	flights, err := models.GetAllFlights()

	return flights, err


}

func GetFlights(p graphql.ResolveParams) (i interface{},e error){
	source :=p.Args["source"].(string)
	destination :=p.Args["destination"].(string)

	flights ,err:= models.GetFlights(source,destination)

	return flights,err
}

func FilterFlights(p graphql.ResolveParams) (i interface{}, e error){

	airlines := p.Args["airlines"]
	facilities := p.Args["facilities"]
	departures:=p.Args["departures"]
	arrivals:=p.Args["arrivals"]
	duration:=p.Args["duration"].(int)


	flights, err := models.FilterFlights(airlines, facilities,departures,arrivals,duration)

	return flights,err
}