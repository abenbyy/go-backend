package _type

import "github.com/graphql-go/graphql"

var carVendorType *graphql.Object

func GetCarVendorType() *graphql.Object{
	if carVendorType == nil{
		carVendorType = graphql.NewObject(graphql.ObjectConfig{
			Name:"CarVendorType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"name": &graphql.Field{
					Type:graphql.String,
				},
				"city": &graphql.Field{
					Type:graphql.String,
				},
				"price": &graphql.Field{
					Type:graphql.Int,
				},
			},
		})
	}

	return carVendorType
}