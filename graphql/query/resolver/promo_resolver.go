package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func GetPromos(p graphql.ResolveParams) (i interface{},e error){
	promos ,err:= models.GetPromos()

	return promos,err
}

func GetPromo(p graphql.ResolveParams) (i interface{}, e error){
	id := p.Args["id"].(int)

	promo, err:= models.GetPromo(id)

	return promo, err
}

func GetOtherPromos(p graphql.ResolveParams)(i interface{}, e error){
	id := p.Args["id"].(int)

	promos, err:= models.GetOtherPromo(id)

	return promos, err

}
