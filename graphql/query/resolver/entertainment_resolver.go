package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)


func GetEntertainment(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)

	ent:=models.GetEntertainment(id)

	return ent, nil
}
func GetAllEntertainments(p graphql.ResolveParams)(i interface{}, e error){
	ents:= models.GetAllEntertainments()

	return ents, nil
}
func GetEntertainments(p graphql.ResolveParams) (i interface{},e error){
	typ :=p.Args["type"].(string)


	entertainments ,err:= models.GetEntertainments(typ)

	return entertainments,err
}

func GetTrendingEntertainment(p graphql.ResolveParams) (i interface{},e error){
	typ :=p.Args["type"].(string)


	entertainments ,err:= models.GetTrendingEntertainments(typ)

	return entertainments,err
}

func GetBestEntertainment(p graphql.ResolveParams) (i interface{},e error){
	entertainments ,err:= models.GetBestEntertainments()

	return entertainments,err
}




