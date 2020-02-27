package _type

import "github.com/graphql-go/graphql"

var entertainmentTicketType *graphql.Object


func GetEntertainmentTicketType() *graphql.Object{
	if entertainmentTicketType == nil{
		entertainmentTicketType = graphql.NewObject(graphql.ObjectConfig{
			Name:"EntertainmentTicketType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"name": &graphql.Field{
					Type:graphql.String,
				},
				"price": &graphql.Field{
					Type:graphql.Int,
				},
			},
		})
	}

	return entertainmentTicketType
}

