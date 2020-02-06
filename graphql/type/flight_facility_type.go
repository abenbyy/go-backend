package _type

import "github.com/graphql-go/graphql"

var flightFacilityType *graphql.Object

func GetFlightFacilityType() *graphql.Object{
	if flightFacilityType == nil{
		flightFacilityType = graphql.NewObject(graphql.ObjectConfig{
			Name:"FlightFacilityType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"name": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}

	return flightFacilityType
}