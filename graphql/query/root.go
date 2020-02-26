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
			"filterflights":{
				Type:graphql.NewList(typ.GetFlightType()),
				Args:graphql.FieldConfigArgument{
					"airlines": &graphql.ArgumentConfig{

						Type: graphql.NewList(graphql.String),
					},
					"facilities": &graphql.ArgumentConfig{
							Type: graphql.NewList(graphql.String),
					},
					"departures":&graphql.ArgumentConfig{
								Type:graphql.NewList(graphql.Int),
					},
					"arrivals": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.Int),
					},
					"duration": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.FilterFlights,
				Description: "Filter Flights",
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
			"searchhotel":{
				Type: graphql.NewList(typ.GetHotelType()),
				Args:graphql.FieldConfigArgument{
					"city": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetHotels,
				Description: "Search Hotels",
			},
			"filterhotel":{
				Type:graphql.NewList(typ.GetHotelType()),
				Args:graphql.FieldConfigArgument{
					"stars": &graphql.ArgumentConfig{
						Type: graphql.NewList(graphql.Int),
					},
				},
				Resolve:res.FilterHotels,
				Description:"Filter Hotel",
			},
			"nearesthotel":{
				Type:graphql.NewList(typ.GetHotelType()),
				Args:graphql.FieldConfigArgument{
					"lat": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"lng": &graphql.ArgumentConfig{
						Type: graphql.Float,
					},
					"amount": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve:res.GetNearestHotels,
				Description:"Get Nearest X Hotel",
			},
			"searchcar":{
				Type: graphql.NewList(typ.GetCarType()),
				Args:graphql.FieldConfigArgument{
					"city": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetCars,
				Description: "Search Cars",
			},

			"searchentertainment":{
				Type: graphql.NewList(typ.GetEntertainmentType()),
				Args:graphql.FieldConfigArgument{
					"type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetEntertainments,
				Description: "Search Entertainments",
			},

			"trendingentertainment":{
				Type: graphql.NewList(typ.GetEntertainmentType()),
				Args:graphql.FieldConfigArgument{
					"type": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: res.GetTrendingEntertainment,
				Description: "Get Trending Entertainments",
			},

			"bestentertainment":{
				Type: graphql.NewList(typ.GetEntertainmentType()),
				Resolve: res.GetBestEntertainment,
				Description: "Get Best Entertainments",
			},

			"allblogs":{
				Type: graphql.NewList(typ.GetBlogType()),
				Resolve: res.GetBlogs,
				Description: "Get All Blogs",
			},
			"popularblogs":{
				Type: graphql.NewList(typ.GetBlogType()),
				Resolve: res.GetPopularBlogs,
				Description: "Get Popular Blogs",
			},
			"blog":{
				Type: typ.GetBlogType(),
				Args:graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: res.GetBlog,
				Description: "Get Blog by Id",
			},
		},
	})
}