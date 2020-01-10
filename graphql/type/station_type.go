package _type

import "github.com/graphql-go/graphql"

var stationType *graphql.Object

func GetStationType() *graphql.Object{
	if stationType == nil{
		stationType = graphql.NewObject(graphql.ObjectConfig{
			Name:"StationType",
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
			},
		})
	}

	return stationType
}