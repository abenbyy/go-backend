package resolver

import(
	"github.com/abenbyy/go-backend/models"
	"github.com/graphql-go/graphql"
	"time"
)

func CreateTrip(p graphql.ResolveParams)(i interface{}, e error){
	train:= p.Args["train_refer"].(int)
	from:= p.Args["from_refer"].(int)
	to:= p.Args["to_refer"].(int)
	duration:= p.Args["duration"].(int)
	price:= p.Args["price"].(int)

	if train<=0|| from<=0||to<=0||duration<=0||price<=0{
		return nil, nil
	}

	t:= models.Trip{
		TrainRefer: train,

		FromRefer:  from,

		ToRefer:    to,
		Departure:  time.Date(2020,3,20,1,0,0,0,time.Local),
		Arrival:    time.Date(2020,3,20,3,0,0,0,time.Local),
		Duration:   duration,
		Price:      price,
		Tax:        0,
		Service:    0,
	}

	models.CreateTrip(t)
	return nil, nil
}

func UpdateTrip(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)
	train:= p.Args["train_refer"].(int)
	from:= p.Args["from_refer"].(int)
	to:= p.Args["to_refer"].(int)
	duration:= p.Args["duration"].(int)
	price:= p.Args["price"].(int)


	t:= models.Trip{
		TrainRefer: train,

		FromRefer:  from,

		ToRefer:    to,
		Departure:  time.Date(2020,3,20,1,0,0,0,time.Local),
		Arrival:    time.Date(2020,3,20,3,0,0,0,time.Local),
		Duration:   duration,
		Price:      price,
		Tax:        0,
		Service:    0,
	}

	models.UpdateTrip(id,t)
	return nil, nil
}



func DeleteTrip(p graphql.ResolveParams)(i interface{}, e error){
	id:= p.Args["id"].(int)

	models.DeleteTrip(id)

	return nil, nil
}