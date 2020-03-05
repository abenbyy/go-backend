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
			"updateuser":{
				Type: typ.GetUserType(),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"firstname": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"lastname": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"phone": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"city": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"postcode": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"language": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"currency": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.UpdateUser,
				Description: "Update User",
			},
			"createblog":{
				Type: typ.GetBlogType(),
				Args: graphql.FieldConfigArgument{
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"category": &graphql.ArgumentConfig{
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
			"updateblog":{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"title": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"category": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.UpdateBlog,
				Description: "Update Blog",
			},
			"createentertainment":{
				Type:graphql.String,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"startdate": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"terms": &graphql.ArgumentConfig{
						Type: graphql.String,
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
					"image": &graphql.ArgumentConfig{
						Type: graphql.String,
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
					"startdate": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"description": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"terms": &graphql.ArgumentConfig{
						Type: graphql.String,
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
			"deleteblog":{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
				},
				Resolve: res.DeleteBlog,
				Description: "Delete Blog",
			},
			"deletetrip":{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
				},
				Resolve: res.DeleteTrip,
				Description: "Delete Trip",
			},
			"deleteflight":{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
				},
				Resolve: res.DeleteFlight,
				Description: "Delete Flight",
			},
			"createhotel":{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"rating": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"area": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"city": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"province": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"facilities": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.Boolean),
					},
				},
				Resolve: res.CreateHotel,
				Description: "Create Hotel",
			},
			"updatehotel":{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"address": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"rating": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"image": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"area": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"city": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"province": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"latitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"longitude": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"facilities": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.Boolean),
					},
				},
				Resolve: res.UpdateHotel,
				Description: "Update Hotel",
			},
			"deletehotel":{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
				},
				Resolve: res.DeleteHotel,
				Description: "Delete Hotel",
			},
			"createflight":{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"airline_refer": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"from_refer": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"to_refer": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"rfrom": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.String),
					},
					"rfromcode": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.String),
					},
					"rto": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.String),
					},
					"rtocode": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.String),
					},
					"fdurations": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.Int),
					},
					"tdurations": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.Int),
					},
					"types": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.String),
					},
					"names": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.String),
					},

				},
				Resolve: res.InsertFlight,
				Description: "Create Flight",
			},
			"updateflight":{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"airline_refer": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"from_refer": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"to_refer": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
				},
				Resolve: res.UpdateFlight,
				Description: "Update Flight",
			},
			"createtrip":{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"train_refer": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"from_refer": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"to_refer": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
				},
				Resolve: res.CreateTrip,
				Description: "Create Train Trip",
			},
			"updatetrip":{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"train_refer": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"from_refer": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"to_refer": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"price": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
				},
				Resolve: res.UpdateTrip,
				Description: "Create Train Trip",
			},
			"createchat":{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"user1": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"user2": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
				},
				Resolve: res.CreateChat,
				Description: "Create Chat",
			},
			"updatechat":{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,

					},
					"content": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.UpdateChat,
				Description: "Update Chat",
			},
		},
	})
}