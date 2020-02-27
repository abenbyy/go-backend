package models

import (
	"fmt"
	"github.com/abenbyy/go-backend/database"
	"time"
)

type Entertainment struct {
	Id			int 		`gorm:primary_key`
	Name 		string
	Type 		string
	IsTrend		bool
	IsBest		bool
	Address 	string
	NeedDate	bool
	Tickets		[]EntertainmentTicket	`gorm:"foreignkey:TicketRefer"`
	Latitude	float64
	Longitude	float64
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time
}

type EntertainmentTicket struct {
	Id		    int 		`gorm:primary_key`
	TicketRefer int
	Name		string
	Price		int
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time
}

func init(){

	db, err := database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()


	db.DropTableIfExists(&EntertainmentTicket{})
	db.DropTableIfExists(&Entertainment{})

	db.AutoMigrate(&Entertainment{})
	db.AutoMigrate(&EntertainmentTicket{})

	db.Model(&EntertainmentTicket{}).AddForeignKey("ticket_refer","entertainments(id)","CASCADE","CASCADE")


	SeedEntertainmentData()
}

func GetEntertainments(enttype string)([]Entertainment, error){
	db, err:= database.Connect()

	if err!=nil{
		panic(err)
	}
	defer db.Close()

	var entertainments []Entertainment

	db.Where("type = ?",enttype).Preload("Tickets").Find(&entertainments)

	return entertainments, nil
}

func GetTrendingEntertainments(enttype string)([]Entertainment, error){
	db, err:= database.Connect()

	if err!=nil{
		panic(err)
	}
	defer db.Close()

	var entertainments []Entertainment

	db.Where("type = ? AND is_trend = ?",enttype, true).Preload("Tickets").Find(&entertainments)

	return entertainments, nil
}

func GetBestEntertainments()([]Entertainment, error){
	db, err := database.Connect()

	if err !=nil{
		panic(err)
	}

	defer db.Close()

	var entertainments []Entertainment
	db.Where("is_best = ?",true).Preload("Tickets").Find(&entertainments)

	return entertainments, nil
}


func CreateEntertainment(ent Entertainment, tickets []EntertainmentTicket){
	db ,err:= database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()

	db.Create(&ent)

	var newent Entertainment
	db.Where("name = ? ",ent.Name).First(&newent)
	fmt.Println(newent)
	CreateTicket(newent.Id, tickets)
}

func CreateTicket(refer int, tickets []EntertainmentTicket){
	db, err:= database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()

	for i := range(tickets){
		tickets[i].TicketRefer = refer
		db.Create(&EntertainmentTicket{
			TicketRefer: tickets[i].TicketRefer,
			Name:        tickets[i].Name,
			Price:       tickets[i].Price,
		})
	}
	//fmt.Print(tickets)



	var tick []EntertainmentTicket
	db.Where("ticket_refer = ?", refer).Find(&tick)
	fmt.Println(tick)
}

func UpdateEntertainment(id int, ent Entertainment)(e error){

	db, err:= database.Connect()
	if err!=nil{
		panic(err)
	}

	defer db.Close()


	fail := db.Model(&Entertainment{}).Where("id = ?",id).Updates(Entertainment{
		Name:      ent.Name,
		Type:      ent.Type,
		Address:   ent.Address,
		Latitude:  ent.Latitude,
		Longitude: ent.Longitude,
	}).Error

	return fail

}

func DeleteEntertainment(id int)(e error){
	db, err := database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()

	fail := db.Where("id = ?",id).Delete(Entertainment{}).Error

	if(fail != nil){
		return e
	}

	return nil

}




func SeedEntertainmentData(){
	db, err := database.Connect()

	if err !=nil{
		panic(err)
	}
	defer db.Close()


	//1
	db.Create(&Entertainment{
		Name:      "Spa Bali Oddysey",
		Type:      "Activity",
		Address:   "Jl. Ranggagading No.3, Tamansari",
		NeedDate:  true,
		IsTrend:   true,
		IsBest:    false,
		Latitude:  -6.902478599999999,
		Longitude: 107.60972979999997,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 1,
		Name:        "Bali Life Pack",
		Price:       179000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 1,
		Name:        "Signature Deluxe",
		Price:       199000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 1,
		Name:        "Amazing Four Hands",
		Price:       239000,
	})


	//2
	db.Create(&Entertainment{
		Name:      "Osaka Amazing Pass Tickets",
		Type:      "Activity",
		Address:   "Osaka, Japan",
		NeedDate:  true,
		IsTrend:   true,
		IsBest:     true,
		Latitude:  34.7923196,
		Longitude: 135.4390672,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 2,
		Name:        "Osaka Amazing Pass Children",
		Price:       300000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 2,
		Name:        "Osaka Amazing Pass Children",
		Price:       350000,
	})


	//3
	db.Create(&Entertainment{
		Name:      "Kelor Island Tour Packages",
		Type:      "Activity",
		Address:   "Kepulauan Seribu, Indonesia",
		NeedDate:  true,
		IsTrend:   true,
		IsBest:    true,
		Latitude:  -5.6122404,
		Longitude: 106.61699640,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 3,
		Name:        "One Day Trip",
		Price:       82500,
	})


	//4
	db.Create(&Entertainment{
		Name:      "Spa Bali Oddysey",
		Type:      "Activity",
		Address:   "Jl. Ranggagading No.3, Tamansari",
		NeedDate:  true,
		IsTrend:   false,
		IsBest:    false,
		Latitude:  -6.902478599999999,
		Longitude: 107.60972979999997,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 4,
		Name:        "Bali Life Pack",
		Price:       179000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 4,
		Name:        "Signature Deluxe",
		Price:       199000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 4,
		Name:        "Amazing Four Hands",
		Price:       239000,
	})


	//5
	db.Create(&Entertainment{
		Name:      "Osaka Amazing Pass Tickets",
		Type:      "Activity",
		Address:   "Osaka, Japan",
		NeedDate:  true,
		IsTrend:   false,
		IsBest:     false,
		Latitude:  34.7923196,
		Longitude: 135.4390672,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 5,
		Name:        "Osaka Amazing Pass Children",
		Price:       300000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 5,
		Name:        "Osaka Amazing Pass Children",
		Price:       350000,
	})


	//6
	db.Create(&Entertainment{
		Name:      "Kelor Island Tour Packages",
		Type:      "Activity",
		Address:   "Kepulauan Seribu, Indonesia",
		NeedDate:  true,
		IsTrend:   false,
		IsBest:    false,
		Latitude:  -5.6122404,
		Longitude: 106.61699640,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 6,
		Name:        "One Day Trip",
		Price:       82500,
	})

	//7
	db.Create(&Entertainment{
		Name:      "Timezone Lippo Mall Puri",
		Type:      "Attraction",
		Address:   "Lippo Mall Puri, Jalan Puri Indah Raya, RT.3/RW.2",
		NeedDate:  true,
		IsTrend:   true,
		IsBest:    false,
		Latitude:  -6.1878581,
		Longitude: 106.73871800,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 7,
		Name:        "Voucher IDR 200.000",
		Price:       120000,
	})

	//8
	db.Create(&Entertainment{
		Name:      "KIDZANIA JAKARTA",
		Type:      "Attraction",
		Address:   "Pacific Place Mall, Sudirman Central Business District",
		NeedDate:  true,
		IsTrend:   true,
		IsBest:    false,
		Latitude:  -6.2240509,
		Longitude: 106.8098201000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 8,
		Name:        "Half Day Pass",
		Price:       185000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 8,
		Name:        "One Day Pass",
		Price:       225000,
	})


	//9
	db.Create(&Entertainment{
		Name:      "Trans Studio Bandung",
		Type:      "Attraction",
		Address:   "Trans Studio Bandung, Jalan Gatot Subroto",
		NeedDate:  true,
		IsTrend:   true,
		IsBest:    true,
		Latitude:  -6.92506319999,
		Longitude: 107.6365103000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 9,
		Name:        "Weekday Ticket",
		Price:       185000,
	})

	//10
	db.Create(&Entertainment{
		Name:      "Timezone Lippo Mall Puri",
		Type:      "Attraction",
		Address:   "Lippo Mall Puri, Jalan Puri Indah Raya, RT.3/RW.2",
		NeedDate:  true,
		IsTrend:   false,
		IsBest:    false,
		Latitude:  -6.1878581,
		Longitude: 106.73871800,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 10,
		Name:        "Voucher IDR 200.000",
		Price:       120000,
	})

	//11
	db.Create(&Entertainment{
		Name:      "KIDZANIA JAKARTA",
		Type:      "Attraction",
		Address:   "Pacific Place Mall, Sudirman Central Business District",
		NeedDate:  true,
		IsTrend:   false,
		IsBest:    false,
		Latitude:  -6.2240509,
		Longitude: 106.8098201000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 11,
		Name:        "Half Day Pass",
		Price:       185000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 11,
		Name:        "One Day Pass",
		Price:       225000,
	})


	//12
	db.Create(&Entertainment{
		Name:      "Trans Studio Bandung",
		Type:      "Attraction",
		Address:   "Trans Studio Bandung, Jalan Gatot Subroto",
		NeedDate:  true,
		IsTrend:   false,
		IsBest:    false,
		Latitude:  -6.92506319999,
		Longitude: 107.6365103000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 12,
		Name:        "Weekday Ticket",
		Price:       185000,
	})

	//13
	db.Create(&Entertainment{
		Name:      "Head in the Clouds Jakarta",
		Type:      "Event",
		Address:   "JIExpo Kemayoran, Jalan Rajawali Selatan Raya, RT.2/RW.6",
		NeedDate:  false,
		IsTrend:   true,
		IsBest:    true,
		Latitude:  -6.1463445,
		Longitude: 106.84581630000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 13,
		Name:        "General Admission - REGULAR",
		Price:       1488000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 13,
		Name:        "Exclusive GA Bundle",
		Price:       2088000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 13,
		Name:        "Exclusive VIP Bundle",
		Price:       3488000,
	})

	//14
	db.Create(&Entertainment{
		Name:      "Asian Sound Syndicate",
		Type:      "Event",
		Address:   "Helipad Parking Ground, Jl. Pintu Satu Senayan",
		NeedDate:  false,
		IsTrend:   true,
		IsBest:    true,
		Latitude:  -6.21493429999,
		Longitude: 106.80010319999,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 14,
		Name:        "GA - Early Bird",
		Price:       575000,
	})

	db.Create(&EntertainmentTicket{
		TicketRefer: 14,
		Name:        "The Syndicate",
		Price:       1265000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 14,
		Name:        "Ultimates",
		Price:       1840000,
	})

	//15
	db.Create(&Entertainment{
		Name:      "Music Blast Live 2020",
		Type:      "Event",
		Address:   "Revo Town Bekasi",
		NeedDate:  false,
		IsTrend:   true,
		IsBest:    false,
		Latitude:  37.7749295,
		Longitude: -122.4194155000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 15,
		Name:        "General Admission",
		Price:       120000,
	})


	//16
	db.Create(&Entertainment{
		Name:      "Head in the Clouds Jakarta",
		Type:      "Event",
		Address:   "JIExpo Kemayoran, Jalan Rajawali Selatan Raya, RT.2/RW.6",
		NeedDate:  false,
		IsTrend:   false,
		IsBest:    false,
		Latitude:  -6.1463445,
		Longitude: 106.84581630000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 16,
		Name:        "General Admission - REGULAR",
		Price:       1488000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 16,
		Name:        "Exclusive GA Bundle",
		Price:       2088000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 16,
		Name:        "Exclusive VIP Bundle",
		Price:       3488000,
	})

	//17
	db.Create(&Entertainment{
		Name:      "Asian Sound Syndicate",
		Type:      "Event",
		Address:   "Helipad Parking Ground, Jl. Pintu Satu Senayan",
		NeedDate:  false,
		IsTrend:   false,
		IsBest:    false,
		Latitude:  -6.21493429999,
		Longitude: 106.80010319999,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 17,
		Name:        "GA - Early Bird",
		Price:       575000,
	})

	db.Create(&EntertainmentTicket{
		TicketRefer: 17,
		Name:        "The Syndicate",
		Price:       1265000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 17,
		Name:        "Ultimates",
		Price:       1840000,
	})

	//18
	db.Create(&Entertainment{
		Name:      "Music Blast Live 2020",
		Type:      "Event",
		Address:   "Revo Town Bekasi",
		NeedDate:  false,
		IsTrend:   false,
		IsBest:    false,
		Latitude:  37.7749295,
		Longitude: -122.4194155000,
	})
	db.Create(&EntertainmentTicket{
		TicketRefer: 18,
		Name:        "General Admission",
		Price:       120000,
	})
}





