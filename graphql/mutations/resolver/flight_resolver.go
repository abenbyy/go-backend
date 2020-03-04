package resolver

import(
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
	"time"
)


func InsertFlight(p graphql.ResolveParams)(i interface{}, e error){

	airline:=p.Args["airline_refer"].(int)
	from:=p.Args["from_refer"].(int)
	to:=p.Args["to_refer"].(int)
	duration:=p.Args["duration"].(int)
	price:=p.Args["price"].(int)
	rfrom:=p.Args["rfrom"].([]interface{})
	rfromcode:=p.Args["rfromcode"].([]interface{})
	rto:=p.Args["rto"].([]interface{})
	rtocode:=p.Args["rtocode"].([]interface{})
	fdurations:= p.Args["fdurations"].([]interface{})
	tdurations:= p.Args["tdurations"].([]interface{})
	aeroplanetypes:= p.Args["types"].([]interface{})
	aeroplanenames:= p.Args["names"].([]interface{})

	if airline <= 0||from <=0||to <=0||duration <=0||price<=0||len(rfrom)<=0||len(rfromcode)<=0||len(rto)<=0||len(rtocode)<=0||len(fdurations)<=0||len(tdurations)<=0||len(aeroplanenames)<=0||len(aeroplanetypes)<=0{
		return nil,nil
	}
	f:= models.Flight{

		AirlineRefer:  airline,
		FromRefer:     from,
		ToRefer:       to,
		Departure:     time.Date(2020,3,20,0,0,0,0,time.Local),
		Arrival:       time.Date(2020,3,20,0,0,0,0,time.Local),
		Duration:      duration,
		Price:         price,
		Tax:           0,
		ServiceCharge: 0,

	}

	var routes []models.FlightRoute
	routes = nil

	for i:=range(rfrom){
		route:= models.FlightRoute{
			From:             rfrom[i].(string),
			FromCode:         rfromcode[i].(string),
			To:               rto[i].(string),
			ToCode:           rtocode[i].(string),
			FlightDuration:   fdurations[i].(int),
			TransitDuration:  tdurations[i].(int),
			AeroplaneType:    aeroplanetypes[i].(string),
			AeroplaneName:    aeroplanenames[i].(string),
		}

		routes = append(routes, route)
	}


	models.CreateFlight(f, routes)


	return nil,nil
}

func UpdateFlight(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)
	airline:=p.Args["airline_refer"].(int)
	from:=p.Args["from_refer"].(int)
	to:=p.Args["to_refer"].(int)
	duration:=p.Args["duration"].(int)
	price:=p.Args["price"].(int)

	f:= models.Flight{
		AirlineRefer:  airline,
		FromRefer:     from,
		ToRefer:       to,
		Duration:      duration,
		Price:         price,
	}

	models.UpdateFlight(id,f)


	return nil,nil
}
func DeleteFlight(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)

	models.DeleteFlight(id)

	return nil, nil
}