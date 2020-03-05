package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func GetAllChat(p graphql.ResolveParams) (i interface{},e error){
	id:= p.Args["userid"].(int)


	chats, err := models.GetAllChats(id)

	return chats,err
}

func GetChat(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)

	chats, err:= models.GetChatById(id)

	return chats,err
}