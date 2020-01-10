package _type

import "github.com/graphql-go/graphql"

var trainType *graphql.Object

func GetTrainType() *graphql.Object{
	if trainType == nil{
		trainType = graphql.NewObject(graphql.ObjectConfig{
			Name:"TrainType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"name": &graphql.Field{
					Type:graphql.String,
				},
				"class": &graphql.Field{
					Type:graphql.String,
				},
				"subclass": &graphql.Field{
					Type:graphql.String,
				},
				"code": &graphql.Field{
					Type:graphql.String,
				},
			},
		})
	}

	return trainType
}