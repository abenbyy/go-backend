package _type

import "github.com/graphql-go/graphql"

var promoType *graphql.Object

func GetPromoType() *graphql.Object{
	if promoType == nil{
		promoType = graphql.NewObject(graphql.ObjectConfig{
			Name:"PromoType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"title": &graphql.Field{
					Type:graphql.String,
				},
				"content": &graphql.Field{
					Type:graphql.String,
				},
				"image": &graphql.Field{
					Type:graphql.String,
				},
				"code": &graphql.Field{
					Type:graphql.String,
				},
				"description": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}

	return promoType
}