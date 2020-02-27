package mutations

import (
	res "github.com/abenbyy/go-backend/graphql/mutations/resolver"
	typ "github.com/abenbyy/go-backend/graphql/type"
	"github.com/graphql-go/graphql"
)

func GetRoot() *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name:"RootMutation",
		Fields: graphql.Fields{
			"createuser":{
				Type: typ.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"firstname": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"lastname": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"phone": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.CreateUser,
				Description: "Get User Based on Email",
			},
			"createblog":{
				Type: typ.GetBlogType(),
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.CreateBlog,
				Description: "Create Blog",
			},
			"createentertainment":{
				Type:graphql.String,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"needate": &graphql.ArgumentConfig{
						Type: graphql.Boolean,
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"ticket_name": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.String),
					},
					"ticket_price": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.Int),
					},
				},
				Resolve: res.CreateEntertainment,
				Description: "Create Entertainment",
			},
			"updateentertainment":{
				Type:graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
				},
				Resolve: res.UpdateEntertainment,
				Description: "Update Entertainment",
			},
			"deleteentertainment":{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
				},
				Resolve: res.DeleteEntertainment,
				Description: "Delete Entertainment",
			},
		},
	})
}