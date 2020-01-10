package _type

import "github.com/graphql-go/graphql"

var tripType *graphql.Object

func GetTripType() *graphql.Object{
	if tripType == nil{
		tripType = graphql.NewObject(graphql.ObjectConfig{
			Name:"TripType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"train": &graphql.Field{
					Type:GetTrainType(),
				},
				"from": &graphql.Field{
					Type:GetStationType(),
				},
				"to": &graphql.Field{
					Type:GetStationType(),
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
				"service": &graphql.Field{
					Type:graphql.Int,
				},
			},
		})
	}

	return tripType
}