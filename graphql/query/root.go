package query

import (
	res "github.com/abenbyy/go-backend/graphql/query/resolver"
	typ "github.com/abenbyy/go-backend/graphql/type"
	"github.com/graphql-go/graphql"
)

func GetRoot() *graphql.Object{
	return graphql.NewObject(graphql.ObjectConfig{
		Name:"RootQuery",
		Fields: graphql.Fields{
			"users":{
				Type:	graphql.NewList(typ.GetUserType()),
				Resolve: res.GetUsers,
				Description: "Get All Users",
			},
			"userbyemail":{
				Type: graphql.NewList(typ.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.GetUserByEmail,
				Description: "Get User Based on Email",
			},
			"userbyphone":{
				Type: graphql.NewList(typ.GetUserType()),
				Args: graphql.FieldConfigArgument{
					"phone": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.GetUserByPhone,
				Description: "Get User Based on Phone",
			},
			"userbyphoneoremail":{
				Type: graphql.NewList(typ.GetUserType()),
				Args:graphql.FieldConfigArgument{
					"arg": &graphql.ArgumentConfig{
						Type: graphql.String,

					},
				},
				Resolve: res.GetUserByPhoneOrEmail,
				Description: "Find User to Validate Login",
			},
			"airports":{
				Type:	graphql.NewList(typ.GetAirportType()),
				Resolve: res.GetAirports,
				Description: "Get All Airports",
			},
			"airlines":{
				Type:	graphql.NewList(typ.GetAirlineType()),
				Resolve: res.GetAirlines,
				Description: "Get All Airlines",
			},
			"searchflight":{
				Type: graphql.NewList(typ.GetFlightType()),
				Args:graphql.FieldConfigArgument{
					"source": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"destination": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetFlights,
				Description: "Search Flights",
			},
			"trains":{
				Type:	graphql.NewList(typ.GetTrainType()),
				Resolve: res.GetTrains,
				Description: "Get All Trains",
			},
			"stations":{
				Type:	graphql.NewList(typ.GetStationType()),
				Resolve: res.GetStations,
				Description: "Get All Stations",
			},
			"searchtrip":{
				Type: graphql.NewList(typ.GetTripType()),
				Args:graphql.FieldConfigArgument{
					"source": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"destination": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetTrips,
				Description: "Search Trips",
			},
		},
	})
}