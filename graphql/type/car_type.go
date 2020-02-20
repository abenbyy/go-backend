package _type

import "github.com/graphql-go/graphql"

var carType *graphql.Object

func GetCarType() *graphql.Object{
	if carType == nil{
		carType = graphql.NewObject(graphql.ObjectConfig{
			Name:"CarType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"brand": &graphql.Field{
					Type:graphql.String,
				},
				"model": &graphql.Field{
					Type:graphql.String,
				},
				"passenger": &graphql.Field{
					Type:graphql.Int,
				},
				"luggage": &graphql.Field{
					Type:graphql.Int,
				},
				"vendors": &graphql.Field{
					Type:graphql.NewList(GetCarVendorType()),
				},
			},
		})
	}

	return carType
}