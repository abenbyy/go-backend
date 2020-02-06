package _type

import "github.com/graphql-go/graphql"

var hotelFacilityType *graphql.Object

func GetHotelFacilityType() *graphql.Object{
	if hotelFacilityType == nil{
		hotelFacilityType = graphql.NewObject(graphql.ObjectConfig{
			Name:"HotelFacilityType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"ac":&graphql.Field{
					Type:graphql.Boolean,
				},
				"parking":&graphql.Field{
					Type:graphql.Boolean,
				},
				"receptionist":&graphql.Field{
					Type:graphql.Boolean,
				},
				"pool":&graphql.Field{
					Type:graphql.Boolean,
				},
				"lift":&graphql.Field{
					Type:graphql.Boolean,
				},
				"restaurant":&graphql.Field{
					Type:graphql.Boolean,
				},
				"wifi":&graphql.Field{
					Type:graphql.Boolean,
				},
				"spa":&graphql.Field{
					Type:graphql.Boolean,
				},
			},
		})
	}

	return hotelFacilityType
}