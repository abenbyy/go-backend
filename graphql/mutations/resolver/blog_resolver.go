package resolver

import(
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func CreateBlog(p graphql.ResolveParams) (i interface{},e error){
	title:=p.Args["title"].(string)
	content:=p.Args["content"].(string)
	category:=p.Args["category"].(string)
	image:=p.Args["image"].(string)


	if title == "" || content == "" || category == ""{
		return nil, nil
	}

	err := models.CreateBlog(title, category,content,image)

	if err!=nil{
		panic(err)
	}

	return nil, nil

}

func UpdateBlog(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)
	title:=p.Args["title"].(string)
	content:=p.Args["content"].(string)
	category:=p.Args["category"].(string)
	image:=p.Args["image"].(string)


	if title == "" || content == "" || category == ""{
		return nil, nil
	}

	bl := models.Blog{

		Title:    title,
		Category: category,
		Content:  content,
		Image:    image,

	}
	models.UpdateBlog(id, bl)

	return nil, nil
}

func DeleteBlog(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)

	models.DeleteBlog(id)

	return nil, nil
}