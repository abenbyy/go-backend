package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)


func GetHotel(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)

	hotel:= models.GetHotel(id)

	return hotel, nil
}
func GetAllHotels(p graphql.ResolveParams)(i interface{}, e error){

	hotels:= models.GetAllHotels()


	return hotels, nil
}
func GetHotels(p graphql.ResolveParams) (i interface{},e error){
	city :=p.Args["city"].(string)


	hotels ,err:= models.GetHotels(city)

	return hotels,err
}

func FilterHotels(p graphql.ResolveParams) (i interface{},e error){
	stars := p.Args["stars"]

	hotels, err := models.FilterHotels(stars)

	return hotels, err
}

func GetNearestHotels(p graphql.ResolveParams)(i interface{}, e error){


	lat := p.Args["lat"].(float64)



	lng := p.Args["lng"].(float64)

	amount := p.Args["amount"].(int)


	hotels, err:= models.GetNearestHotels(lat,lng,amount)

	return hotels, err
}