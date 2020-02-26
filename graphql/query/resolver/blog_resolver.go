package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)


func GetBlog(p graphql.ResolveParams) (i interface{},e error){
	id:= p.Args["id"].(int)


	blog , err:= models.GetBlog(id)

	return blog, err
}

func GetBlogs(p graphql.ResolveParams) (i interface{},e error){
	blogs ,err:= models.GetBlogs()

	return blogs,err
}

func GetPopularBlogs(p graphql.ResolveParams)(i interface{}, e error){

	blogs, err := models.GetPopularBlogs()


	return blogs, err
}