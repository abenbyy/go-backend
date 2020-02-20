package _type

import "github.com/graphql-go/graphql"

var hotelReviewType *graphql.Object

func GetHotelReviewType() *graphql.Object{
	if hotelReviewType == nil{
		hotelReviewType = graphql.NewObject(graphql.ObjectConfig{
			Name:"HotelReviewType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"name": &graphql.Field{
					Type:graphql.String,
				},
				"category": &graphql.Field{
					Type:graphql.String,
				},
				"title": &graphql.Field{
					Type:graphql.String,
				},
				"content": &graphql.Field{
					Type:graphql.String,
				},
				"date": &graphql.Field{
					Type:graphql.String,
				},
				"overall": &graphql.Field{
					Type:graphql.Float,
				},
			},
		})
	}

	return hotelReviewType
}