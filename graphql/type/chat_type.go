package _type

import "github.com/graphql-go/graphql"

var chatType *graphql.Object

func GetChatType() *graphql.Object{
	if chatType == nil{
		chatType = graphql.NewObject(graphql.ObjectConfig{
			Name:"ChatType",
			Fields:graphql.Fields{
				"id": &graphql.Field{
					Type:graphql.Int,
				},
				"user1": &graphql.Field{
					Type:graphql.Int,
				},
				"user2": &graphql.Field{
					Type:graphql.Int,
				},
				"content": &graphql.Field{
					Type:graphql.String,
				},

			},
		})
	}

	return chatType
}