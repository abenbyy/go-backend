package resolver

import(
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func CreateChat(p graphql.ResolveParams) (i interface{},e error){
	user1:= p.Args["user1"].(int)
	user2:= p.Args["user2"].(int)


	models.InsertChat(user1, user2)


	return nil,nil

}

func UpdateChat(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)
	content:= p.Args["content"].(string)

	models.UpdateChat(id, content)

	return nil,nil
}