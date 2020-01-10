package _type

import "github.com/graphql-go/graphql"

var airportType *graphql.Object

func GetAirportType() *graphql.Object{
	if airportType == nil{
		airportType = graphql.NewObject(graphql.ObjectConfig{
			Name:"AirportType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"name": &graphql.Field{
					Type:graphql.String,
				},
				"code": &graphql.Field{
					Type:graphql.String,
				},
				"city": &graphql.Field{
					Type:graphql.String,
				},
				"citycode": &graphql.Field{
					Type:graphql.String,
				},
				"province": &graphql.Field{
					Type:graphql.String,
				},
				"country": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}

	return airportType
}