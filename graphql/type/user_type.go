package _type

import "github.com/graphql-go/graphql"

var userType *graphql.Object

func GetUserType() *graphql.Object{
	if userType == nil{
		userType = graphql.NewObject(graphql.ObjectConfig{
			Name:"UserType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"firstname": &graphql.Field{
					Type:graphql.String,
				},
				"lastname": &graphql.Field{
					Type:graphql.String,
				},
				"password": &graphql.Field{
					Type:graphql.String,
				},
				"email": &graphql.Field{
					Type:graphql.String,
				},
				"phone": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}

	return userType
}