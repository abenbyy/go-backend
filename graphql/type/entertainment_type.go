package _type

import "github.com/graphql-go/graphql"

var entertainmentType *graphql.Object

func GetEntertainmentType() *graphql.Object{
	if entertainmentType == nil{
		entertainmentType = graphql.NewObject(graphql.ObjectConfig{
			Name:"EntertainmentType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"name": &graphql.Field{
					Type:graphql.String,
				},
				"image": &graphql.Field{
					Type:graphql.String,
				},
				"type": &graphql.Field{
					Type:graphql.String,
				},
				"address": &graphql.Field{
					Type:graphql.String,
				},
				"needdate": &graphql.Field{
					Type:graphql.Boolean,
				},
				"tickets": &graphql.Field{
					Type:graphql.NewList(GetEntertainmentTicketType()),
				},
				"latitude": &graphql.Field{
					Type:graphql.Float,
				},
				"longitude": &graphql.Field{
					Type:graphql.Float,
				},
				"description": &graphql.Field{
					Type:graphql.String,
				},
				"terms": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}

	return entertainmentType
}

