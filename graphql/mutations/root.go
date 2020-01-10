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
		},
	})
}