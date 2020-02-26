package resolver

import(
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func CreateBlog(p graphql.ResolveParams) (i interface{},e error){
	title:=p.Args["title"].(string)
	content:=p.Args["content"].(string)
	image:=p.Args["image"].(string)

	err := models.CreateBlog(title,content,image)

	if err!=nil{
		panic(err)
	}

	return nil, nil

}