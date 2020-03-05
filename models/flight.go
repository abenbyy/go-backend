package models

import (
	"fmt"
	"github.com/abenbyy/go-backend/database"
	"github.com/abenbyy/go-backend/middleware"
	"time"
)

type Flight struct {
	Id				int			`gorm:primary_key`
	Airline 		Airline		`gorm:"foreignkey:AirlineRefer"`
	AirlineRefer	int
	From			Airport		`gorm:"foreignkey:FromRefer"`
	FromRefer		int
	To				Airport		`gorm:"foreignkey:ToRefer"`
	ToRefer			int
	Departure		time.Time
	Arrival 		time.Time
	Duration		int
	Price			int				`gorm:"type:int"`
	Tax				int				`gorm:"type:int"`
	ServiceCharge	int				`gorm:"type:int"`
	Facilities		[]FlightFacility		`gorm:"foreignkey:FlightRefer"`
	Routes			[]FlightRoute			`gorm:"foreignkey:FlightRouteRefer"`
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		*time.Time		`sql:index`

}

type FlightFacility struct{
	Id			int			`gorm:primary_key`
	FlightRefer uint
	Name		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time		`sql:index`
}

type FlightRoute struct{
	Id			int
	FlightRouteRefer	int
	From		string
	FromCode	string
	To			string
	ToCode		string
	FlightDuration	int
	TransitDuration	int
	AeroplaneType	string
	AeroplaneName	string
	CreatedAt 	time.Time
	UpdatedAt	time.Time
	DeletedAt 	*time.Time
}

var recentflight []Flight

func init(){
	db, err := database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()
	db.DropTableIfExists(&FlightRoute{})
	db.DropTableIfExists(&FlightFacility{})
	db.DropTableIfExists(&Flight{})

	db.DropTableIfExists(&Airline{})
	db.AutoMigrate(&Airline{})

	SeedAirlineData()

	db.AutoMigrate(&Flight{}).AddForeignKey("airline_refer","airlines(Id)","CASCADE","CASCADE").AddForeignKey("from_refer","airports(Id)","CASCADE","CASCADE").AddForeignKey("to_refer","airports(Id)","CASCADE","CASCADE")
	db.AutoMigrate(&FlightFacility{})
	db.AutoMigrate(&FlightRoute{})
	db.Model(&FlightFacility{}).AddForeignKey("flight_refer","flights(id)","CASCADE","CASCADE")
	db.Model(&FlightRoute{}).AddForeignKey("flight_route_refer","flights(id)","CASCADE","CASCADE")

	SeedFlightData()
}
//
//func SliceFlight(value []Flight, idx int)([]Flight){
//	var result []Flight
//	lim := 5
//	i:= lim*idx
//	offset:= i+lim
//	for i < offset{
//		result = append(result,value[i])
//		i++
//	}
//	return result
//}

func GetAllFlights()([]Flight, error){
	db, err:= database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()
	_, err = ValidateKey(middleware.ApiKey)

	if err != nil{
		return nil, err
	}

	var flights []Flight

	db.Find(&flights)
	for i,_ := range flights{

		db.Model(flights[i]).Related(&flights[i].Airline,"airline_refer").Related(&flights[i].From,"from_refer").Related(&flights[i].To,"to_refer")
	}

	return flights, nil
}

func GetFlight(id int)(Flight){
	db, err:= database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()
	var flight Flight
	_, err = ValidateKey(middleware.ApiKey)

	if err != nil{
		return flight
	}

	db.Where("id = ?",id).Preload("Routes").First(&flight)


	db.Model(flight).Related(&flight.Airline,"airline_refer").Related(&flight.From,"from_refer").Related(&flight.To,"to_refer")

	return flight
}

func GetFlights(source string, destination string)([]Flight,error){
	db, err := database.Connect()
	if err != nil{
		panic(err)
	}

	defer db.Close()

	_, err = ValidateKey(middleware.ApiKey)

	if err != nil{
		return nil , err
	}
	var flights []Flight

	db.Where("from_refer IN (?) AND to_refer IN (?)", db.Table("airports").Select("Id").Where("city = ?",source).SubQuery(), db.Table("airports").Select("Id").Where("city = ?",destination).SubQuery()).Preload("Facilities").Preload("Routes").Find(&flights)

	for i,_ := range flights{

		db.Model(flights[i]).Related(&flights[i].Airline,"airline_refer").Related(&flights[i].From,"from_refer").Related(&flights[i].To,"to_refer")
	}


	recentflight = flights

	return flights,nil

}

func CreateFlight(f Flight, t []FlightRoute){
	db, err:= database.Connect()

	if err!= nil{
		panic(err)
	}

	defer db.Close()
	_, err = ValidateKey(middleware.ApiKey)

	if err != nil{
		return
	}

	db.Create(&f)

	var flight Flight

	db.Where("airline_refer = ? AND from_refer = ? AND to_refer = ? ",f.AirlineRefer,f.FromRefer,f.ToRefer).First(&flight)

	CreateRoutes(flight.Id, t)

}
func CreateRoutes(id int, t []FlightRoute){
	db, err:= database.Connect()

	if err!= nil{
		panic(err)
	}

	defer db.Close()

	_, err = ValidateKey(middleware.ApiKey)

	if err != nil{
		return
	}
	for i:= range(t){
		t[i].FlightRouteRefer = id
		db.Create(&t[i])
	}

}

func UpdateFlight(id int, f Flight){
	db, err:= database.Connect()

	if err!= nil{
		panic(err)
	}

	defer db.Close()

	_, err = ValidateKey(middleware.ApiKey)

	if err != nil{
		return
	}

	db.Model(&Flight{}).Where("id = ?",id).Updates(Flight{
		AirlineRefer:  f.AirlineRefer,
		FromRefer:     f.FromRefer,
		ToRefer:       f.ToRefer,
		Duration:      f.Duration,
		Price:         f.Price,
	})


}



func DeleteFlight(id int){
	db, err:= database.Connect()

	if err!= nil{
		panic(err)
	}

	defer db.Close()
	_, err = ValidateKey(middleware.ApiKey)

	if err != nil{
		return
	}

	db.Where("id = ?",id).Delete(&Flight{})
}

func FilterFlights(reqairlines interface{},reqfacilities interface{}, reqdepartures interface{},reqarrivals interface{}, reqduration int)([]Flight, error){
	db, err:= database.Connect()
	if err != nil{
		panic(err)
	}

	defer db.Close()
	_, err = ValidateKey(middleware.ApiKey)

	if err != nil{
		return nil , err
	}
	rairlines:= reqairlines.([]interface{})
	rfacilities:= reqfacilities.([]interface{})
	rdepartures:= reqdepartures.([]interface{})
	rarrivals:= reqarrivals.([]interface{})
	fmt.Println("From model:")
	fmt.Println(rdepartures)
	fmt.Println(rarrivals)


	if len(rairlines) == 0 && len(rfacilities)==0 && len(rdepartures)==0 && len(rarrivals)==0 && reqduration == 0{
		return recentflight,nil
	}

	var finalFlights []Flight

	temp := recentflight

	var airlinefiltered []Flight

	if len(rairlines)>0 {
		for i,_:= range temp{
			for j,_:= range rairlines{
				if temp[i].Airline.Name == rairlines[j]{
					airlinefiltered = append(airlinefiltered, temp[i])
				}
			}

		}
	}else{
		airlinefiltered = temp
	}
	fmt.Println("Airline filter pass")
	//fmt.Println(airlinefiltered)
	var facilityfiltered []Flight

	if len(rfacilities)> 0 {
		for i,_ :=range airlinefiltered{
			count:=0
			for j,_:= range airlinefiltered[i].Facilities{

				for k,_ :=range rfacilities{
					if airlinefiltered[i].Facilities[j].Name == rfacilities[k]{
						count++
					}
				}

			}
			if count >= len(rfacilities){
				facilityfiltered = append(facilityfiltered,airlinefiltered[i])
			}
		}
	}else{
		facilityfiltered = airlinefiltered
	}
	fmt.Println("Facility filter pass")
	var departurefiltered []Flight

	if len(rdepartures) > 0{
		for i,_ := range facilityfiltered{
			for j,_:= range rdepartures{
				if facilityfiltered[i].Departure.Hour() >= rdepartures[j].(int) && facilityfiltered[i].Departure.Hour() <= (rdepartures[j].(int)+6) {
					departurefiltered = append(departurefiltered,facilityfiltered[i])
					break
				}
			}
		}
	}else{
		departurefiltered = facilityfiltered
	}
	fmt.Println("Departure filter pass")

	var arrivalfiltered []Flight

	if len(rarrivals) > 0{
		for i,_ := range departurefiltered{
			for j,_:= range rarrivals{
				if departurefiltered[i].Arrival.Hour() >= rarrivals[j].(int) && departurefiltered[i].Arrival.Hour() <= (rarrivals[j].(int)+6) {
					arrivalfiltered = append(arrivalfiltered,departurefiltered[i])
					break
				}
			}
		}
	}else{
		arrivalfiltered = departurefiltered
	}


	var durationfiltered []Flight
	if reqduration > 0{
		for i,_ := range arrivalfiltered{
			if arrivalfiltered[i].Duration <= reqduration*60{
				durationfiltered = append(durationfiltered,arrivalfiltered[i])
			}
		}
	}else{
		durationfiltered = arrivalfiltered
	}


	finalFlights = durationfiltered

	//var subq []Airline
	//fmt.Println("asdasd:")
	//fmt.Println(reqairlines)
	//db.Where("name IN (?)", reqairlines).Find(&subq)
	//fmt.Println("ids:")
	//fmt.Println(subq)
	//var subref [100]interface{}
	//for i,_:=range subq{
	//	subref[i] = subq[i].Id
	//}
	//fmt.Println(subref)
	//db.Model(temp).Where("airline_refer IN (?)",db.Table("airlines").Select("id").Where("name IN (?)",reqairlines).SubQuery())
	////db.Model(temp).Where("airline_refer IN (?)",subref).Find(&flights)
	//for i,_ :=range subref{
	//	fmt.Print("current number: ")
	//	fmt.Println(subref[i])
	//	if subref[i] == nil{
	//		break
	//	}
	//	db.Where("airline_refer LIKE ?",subref[i]).Find(&flights)
	//}
	//fmt.Println("result:")
	//fmt.Println(flights)

	return finalFlights, nil


}

func SeedFlightData(){

	db, err:= database.Connect()
	if err!=nil{
		panic(err)
	}

	defer db.Close()

	db.Create(&Flight{

		AirlineRefer:  2,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,5,30,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,8,0,0,0,time.Local),
		Duration:      150,
		Price:         790000,
		Tax:           10000,
		ServiceCharge: 0,
		//Facilities:		[]Facility{
		//	{
		//
		//		Name:        "Baggage",
		//	},
		//	{
		//
		//		Name:        "Meal",
		//	},
		//},
	})
	db.Create(&Flight{

		AirlineRefer:  2,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,6,30,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,9,30,0,0,time.Local),
		Duration:      180,
		Price:         1200000,
		Tax:           0,
		ServiceCharge: 0,
		//Facilities:		[]Facility{
		//	{
		//		FlightRefer: 2,
		//		Name:        "Baggage",
		//	},
		//},
	})

	db.Create(&Flight{

		AirlineRefer:  2,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,9,30,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,12,0,0,0,time.Local),
		Duration:      150,
		Price:         990000,
		Tax:           0,
		ServiceCharge: 10000,
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,10,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,12,45,0,0,time.Local),
		Duration:      165,
		Price:         1500000,
		Tax:           0,
		ServiceCharge: 0,
		//Facilities:  	[]Facility{
		//	{
		//		FlightRefer: 4,
		//		Name:        "Baggage",
		//	},
		//	{
		//		FlightRefer: 4,
		//		Name:        "Entertainment",
		//	},
		//	{
		//		FlightRefer: 4,
		//		Name:        "Meal",
		//	},
		//},
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,21,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,23,0,0,0,time.Local),
		Duration:      120,
		Price:         2000000,
		Tax:           0,
		ServiceCharge: 0,
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,21,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,23,0,0,0,time.Local),
		Duration:      120,
		Price:         2000000,
		Tax:           0,
		ServiceCharge: 0,
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,21,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,23,0,0,0,time.Local),
		Duration:      120,
		Price:         2000000,
		Tax:           0,
		ServiceCharge: 0,
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,21,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,23,0,0,0,time.Local),
		Duration:      120,
		Price:         2000000,
		Tax:           0,
		ServiceCharge: 0,
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,21,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,23,0,0,0,time.Local),
		Duration:      120,
		Price:         2000000,
		Tax:           0,
		ServiceCharge: 0,
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,21,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,23,0,0,0,time.Local),
		Duration:      120,
		Price:         2000000,
		Tax:           0,
		ServiceCharge: 0,
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,21,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,23,0,0,0,time.Local),
		Duration:      120,
		Price:         2000000,
		Tax:           0,
		ServiceCharge: 0,
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,21,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,23,0,0,0,time.Local),
		Duration:      120,
		Price:         2000000,
		Tax:           0,
		ServiceCharge: 0,
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,21,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,23,0,0,0,time.Local),
		Duration:      120,
		Price:         2000000,
		Tax:           0,
		ServiceCharge: 0,
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,21,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,23,0,0,0,time.Local),
		Duration:      120,
		Price:         2000000,
		Tax:           0,
		ServiceCharge: 0,
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,21,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,23,0,0,0,time.Local),
		Duration:      120,
		Price:         2000000,
		Tax:           0,
		ServiceCharge: 0,
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,21,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,23,0,0,0,time.Local),
		Duration:      120,
		Price:         2000000,
		Tax:           0,
		ServiceCharge: 0,
	})
	db.Create(&Flight{

		AirlineRefer:  1,
		FromRefer:     1,
		ToRefer:       2,
		Departure:     time.Date(2020,1,21,21,0,0,0,time.Local),
		Arrival:       time.Date(2020,1,21,23,1,0,0,time.Local),
		Duration:      121,
		Price:         2000000,
		Tax:           0,
		ServiceCharge: 0,
	})


	db.Create(&FlightFacility{

		FlightRefer: 1,
		Name:        "Baggage",
	})

	db.Create(&FlightFacility{

		FlightRefer: 1,
		Name:        "Meal",
	})

	db.Create(&FlightFacility{

		FlightRefer: 2,
		Name:        "Baggage",
	})
	db.Create(&FlightFacility{

		FlightRefer: 4,
		Name:        "Baggage",
	})
	db.Create(&FlightFacility{

		FlightRefer: 4,
		Name:        "Entertainment",
	})
	db.Create(&FlightFacility{

		FlightRefer: 4,
		Name:        "Meal",
	})

	db.Create(&FlightRoute{
		FlightRouteRefer: 1,
		From:             "Soekarno-Hatta",
		FromCode:         "CGK",
		To:               "Kuala Namu",
		ToCode:           "KNO",
		FlightDuration:   90,
		TransitDuration:  0,
		AeroplaneType:    "Boeing 737-800",
		AeroplaneName:    "JT-396",
	})

	db.Create(&FlightRoute{
		FlightRouteRefer: 1,
		From:             "Kuala Namu",
		FromCode:         "KNO",
		To:               "Changi",
		ToCode:           "CHG",
		FlightDuration:   30,
		TransitDuration:  30,
		AeroplaneType:    "Boeing 737-800",
		AeroplaneName:    "JT-277",
	})

	db.Create(&FlightRoute{
		FlightRouteRefer: 4,
		From:             "Soekarno-Hatta",
		FromCode:         "CGK",
		To:               "Sultan Mahmud Badaruddin II",
		ToCode:           "PLM",
		FlightDuration:   0,
		TransitDuration:  90,
		AeroplaneType:    "Boeing 737-800",
		AeroplaneName:    "GA-123",
	})
	db.Create(&FlightRoute{
		FlightRouteRefer: 4,
		From:             "Sultan Mahmud Badaruddin II",
		FromCode:         "PLM",
		To:               "Kuala Namu",
		ToCode:           "KNO",
		FlightDuration:   30,
		TransitDuration:  15,
		AeroplaneType:    "Boeing 737-800",
		AeroplaneName:    "GA-271",
	})
	db.Create(&FlightRoute{
		FlightRouteRefer: 4,
		From:             "Kuala Namu",
		FromCode:         "KNO",
		To:               "Changi",
		ToCode:           "CHG",
		FlightDuration:   15,
		TransitDuration:  15,
		AeroplaneType:    "Canadair Regional Jet 1000",
		AeroplaneName:    "GA-180",
	})
}