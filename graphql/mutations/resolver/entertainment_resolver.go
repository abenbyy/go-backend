package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"

)

func CreateEntertainment(p graphql.ResolveParams) (i interface{},e error){

	name:= p.Args["name"].(string)
	etype:= p.Args["type"].(string)
	address:= p.Args["address"].(string)
	needate:= p.Args["needate"].(bool)
	latitude:= p.Args["latitude"].(float64)
	longitude:= p.Args["longitude"].(float64)

	ticketname := p.Args["ticket_name"].([]interface{})
	ticketprice := p.Args["ticket_price"].([]interface{})




	var tickets []models.EntertainmentTicket
	tickets = nil
	for i := range(ticketname){
		ticket:= models.EntertainmentTicket{
			TicketRefer: 0,
			Name:        ticketname[i].(string),
			Price:       ticketprice[i].(int),
		}
		tickets = append(tickets, ticket)
	}
	ent := models.Entertainment{
		Name:      name,
		Type:      etype,
		IsTrend:   false,
		IsBest:    false,
		Address:   address,
		NeedDate:  needate,
		Latitude:  latitude,
		Longitude: longitude,
	}


	models.CreateEntertainment(ent, tickets)


	return nil, nil

}

func UpdateEntertainment(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)
	name:= p.Args["name"].(string)
	etype:= p.Args["type"].(string)
	address:= p.Args["address"].(string)
	latitude:= p.Args["latitude"].(float64)
	longitude:= p.Args["longitude"].(float64)

	ent := models.Entertainment{
		Name:      name,
		Type:      etype,
		Address:   address,
		Latitude:  latitude,
		Longitude: longitude,

	}

	err := models.UpdateEntertainment(id, ent)

	return err, nil
}

func DeleteEntertainment(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)

	err := models.DeleteEntertainment(id)



	return err,nil

}

