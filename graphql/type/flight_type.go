package _type

import "github.com/graphql-go/graphql"

var flightType *graphql.Object

func GetFlightType() *graphql.Object{
	if flightType == nil{
		flightType = graphql.NewObject(graphql.ObjectConfig{
			Name:"FlightType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"airline": &graphql.Field{
					Type:GetAirlineType(),
				},
				"from": &graphql.Field{
					Type:GetAirportType(),
				},
				"to": &graphql.Field{
					Type:GetAirportType(),
				},
				"departure": &graphql.Field{
					Type:graphql.DateTime,
				},
				"arrival": &graphql.Field{
					Type:graphql.DateTime,
				},
				"duration": &graphql.Field{
					Type:graphql.Int,
				},
				"price": &graphql.Field{
					Type:graphql.Int,
				},
				"tax": &graphql.Field{
					Type:graphql.Int,
				},
				"servicecharge": &graphql.Field{
					Type:graphql.Int,
				},
				"facilities": &graphql.Field{
					Type: graphql.NewList(GetFlightFacilityType()),
				},
				"routes": &graphql.Field{
					Type: graphql.NewList(GetFlightRouteType()),
				},
			},
		})
	}

	return flightType
}