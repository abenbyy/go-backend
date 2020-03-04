package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
)

func CreateHotel(p graphql.ResolveParams)(i interface{}, e error){
	name:= p.Args["name"].(string)
	address:= p.Args["address"].(string)
	rating := p.Args["rating"].(float64)
	image:= p.Args["image"].(string)
	area:= p.Args["area"].(string)
	city:= p.Args["city"].(string)
	province:= p.Args["province"].(string)
	lat:= p.Args["latitude"].(float64)
	long:= p.Args["longitude"].(float64)
	facilites:= p.Args["facilities"].([]interface{})

	if name==""||address==""||rating<0||image==""||area==""||city==""||province==""||len(facilites)==0{
		return nil,nil
	}
	mod:= models.Hotel{
		Name:         name,
		Address:      address,
		Rating:       rating,
		Image:        image,
		Star:         4,
		Area:         area,
		City:         city,
		Province:     province,
		Latitude:     lat,
		Longitude:    long,
		ZoomLevel:    14,
	}

	fac := models.HotelFacility{
		AC:                 facilites[0].(bool),
		Parking:            facilites[1].(bool),
		Receptionist:       facilites[2].(bool),
		Pool:               facilites[3].(bool),
		Lift:               facilites[4].(bool),
		Restaurant:         facilites[5].(bool),
		Wifi:               facilites[6].(bool),
		Spa:                facilites[7].(bool),
	}

	models.CreateHotel(mod,fac)

	return nil,nil
}

func UpdateHotel(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)
	name:= p.Args["name"].(string)
	address:= p.Args["address"].(string)
	rating := p.Args["rating"].(float64)
	image:= p.Args["image"].(string)
	area:= p.Args["area"].(string)
	city:= p.Args["city"].(string)
	province:= p.Args["province"].(string)
	lat:= p.Args["latitude"].(float64)
	long:= p.Args["longitude"].(float64)
	facilites:= p.Args["facilities"].([]interface{})

	if id<=0||name==""||address==""||rating<0||image==""||area==""||city==""||province==""||len(facilites)==0{
		return nil,nil
	}

	mod:= models.Hotel{
		Name:         name,
		Address:      address,
		Rating:       rating,
		Image:        image,
		Area:         area,
		City:         city,
		Province:     province,
		Latitude:     lat,
		Longitude:    long,
	}

	fac := models.HotelFacility{
		AC:                 facilites[0].(bool),
		Parking:            facilites[1].(bool),
		Receptionist:       facilites[2].(bool),
		Pool:               facilites[3].(bool),
		Lift:               facilites[4].(bool),
		Restaurant:         facilites[5].(bool),
		Wifi:               facilites[6].(bool),
		Spa:                facilites[7].(bool),
	}

	models.UpdateHotel(id,mod,fac)

	return nil,nil
}
func DeleteHotel(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)

	models.DeleteHotel(id)

	return nil, nil
}