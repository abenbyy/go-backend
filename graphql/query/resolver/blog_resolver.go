package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func GetBlogs(p graphql.ResolveParams) (i interface{},e error){
	blogs ,err:= models.GetBlogs()

	return blogs,err
}