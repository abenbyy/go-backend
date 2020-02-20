package _type

import "github.com/graphql-go/graphql"

var hotelType *graphql.Object

func GetHotelType() *graphql.Object{
	if hotelType == nil{
		hotelType = graphql.NewObject(graphql.ObjectConfig{
			Name:"HotelType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"name": &graphql.Field{
					Type:graphql.String,
				},
				"address": &graphql.Field{
					Type:graphql.String,
				},
				"rating": &graphql.Field{
					Type:graphql.Float,
				},
				"star": &graphql.Field{
					Type:graphql.Int,
				},
				"locationrate": &graphql.Field{
					Type:graphql.Int,
				},
				"cleanrate": &graphql.Field{
					Type:graphql.Int,
				},
				"roomrate": &graphql.Field{
					Type:graphql.Int,
				},
				"servicerate": &graphql.Field{
					Type:graphql.Int,
				},
				"latitude": &graphql.Field{
					Type:graphql.Float,
				},
				"longitude": &graphql.Field{
					Type:graphql.Float,
				},
				"zoomlevel": &graphql.Field{
					Type:graphql.Int,
				},
				"area": &graphql.Field{
					Type: graphql.String,
				},
				"city": &graphql.Field{
					Type:graphql.String,
				},
				"province": &graphql.Field{
					Type:graphql.String,
				},
				"facilities": &graphql.Field{
					Type: graphql.NewList(GetHotelFacilityType()),
				},
				"rooms": &graphql.Field{
					Type: graphql.NewList(GetHotelRoomType()),
				},
				"reviews": &graphql.Field{
					Type:  graphql.NewList(GetHotelReviewType()),
				},
			},
		})
	}

	return hotelType
}