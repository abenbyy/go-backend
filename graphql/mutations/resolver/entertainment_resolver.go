package resolver

import (
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"

)

func CreateEntertainment(p graphql.ResolveParams) (i interface{},e error){

	name:= p.Args["name"].(string)
	etype:= p.Args["type"].(string)
	address:= p.Args["address"].(string)
	image:= p.Args["image"].(string)
	startdate:= p.Args["startdate"].(string)
	latitude:= p.Args["latitude"].(float64)
	longitude:= p.Args["longitude"].(float64)

	description:= p.Args["description"].(string)
	terms:= p.Args["terms"].(string)

	ticketname := p.Args["ticket_name"].([]interface{})
	ticketprice := p.Args["ticket_price"].([]interface{})

	if name ==""||etype == ""||address==""||image==""||startdate==""||description==""||terms==""|| len(ticketprice)==0||len(ticketname) ==0{
		return nil, nil
	}

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
		Image: image,
		IsTrend:   false,
		IsBest:    false,
		Address:   address,
		NeedDate:  false,
		StartDate: startdate,
		Latitude:  latitude,
		Longitude: longitude,
		Description: description,
		Terms: terms,
	}


	models.CreateEntertainment(ent, tickets)


	return nil, nil

}

func UpdateEntertainment(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)
	name:= p.Args["name"].(string)
	etype:= p.Args["type"].(string)
	address:= p.Args["address"].(string)
	startdate:= p.Args["startdate"].(string)
	latitude:= p.Args["latitude"].(float64)
	longitude:= p.Args["longitude"].(float64)
	description:= p.Args["description"].(string)
	terms:= p.Args["terms"].(string)
	image:= p.Args["image"].(string)


	if id<=0||name ==""||etype == ""||address==""||image==""||startdate==""||description==""||terms==""{
		return nil, nil
	}

	ent := models.Entertainment{
		Name:      name,
		Type:      etype,
		Image:     image,
		Address:   address,
		StartDate: startdate,
		Latitude:  latitude,
		Longitude: longitude,
		Description: description,
		Terms: terms,

	}

	err := models.UpdateEntertainment(id, ent)

	return err, nil
}

func DeleteEntertainment(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)

	err := models.DeleteEntertainment(id)



	return err,nil

}

