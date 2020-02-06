package _type

import "github.com/graphql-go/graphql"

var flightRouteType *graphql.Object

func GetFlightRouteType() *graphql.Object{
	if flightRouteType == nil{
		flightRouteType = graphql.NewObject(graphql.ObjectConfig{
			Name:"FlightRouteType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"from": &graphql.Field{
					Type:graphql.String,
				},
				"fromcode": &graphql.Field{
					Type:graphql.String,
				},
				"to": &graphql.Field{
					Type:graphql.String,
				},
				"tocode": &graphql.Field{
					Type:graphql.String,
				},
				"flightduration": &graphql.Field{
					Type:graphql.Int,
				},
				"transitduration": &graphql.Field{
					Type:graphql.Int,
				},
				"aeroplanetype": &graphql.Field{
					Type:graphql.String,
				},
				"aeroplanename": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}

	return flightRouteType
}