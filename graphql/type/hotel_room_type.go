package _type

import "github.com/graphql-go/graphql"

var hotelRoomType *graphql.Object

func GetHotelRoomType() *graphql.Object{
	if hotelRoomType == nil{
		hotelRoomType = graphql.NewObject(graphql.ObjectConfig{
			Name:"HotelRoomType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"name": &graphql.Field{
					Type:graphql.String,
				},
				"maxguest": &graphql.Field{
					Type:graphql.Int,
				},
				"roomsize": &graphql.Field{
					Type:graphql.Int,
				},
				"beddetail": &graphql.Field{
					Type:graphql.String,
				},
				"breakfast": &graphql.Field{
					Type:graphql.Boolean,
				},
				"wifi": &graphql.Field{
					Type:graphql.Boolean,
				},
				"price": &graphql.Field{
					Type:graphql.Int,
				},
				"freecancel": &graphql.Field{
					Type: graphql.Boolean,
				},
				"payathotel": &graphql.Field{
					Type: graphql.Boolean,
				},
			},
		})
	}

	return hotelRoomType
}