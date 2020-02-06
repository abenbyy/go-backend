package models

import (
	"github.com/abenbyy/go-backend/database"
	"time"
)

type Hotel struct{
	Id 				int				`gorm: primary_key`
	Name 			string			`gorm:"type:varchar(100);not null"`
	Address			string			`gorm:"type:varchar(100);not null"`
	Rating			float32
	Star			int
	LocationRate	int
	CleanRate		int
	RoomRate		int
	ServiceRate		int
	Rooms			[]HotelRoom		`gorm:"foreignkey:HotelRoomRefer"`
	Facilities		[]HotelFacility `gorm:"foreignkey:HotelFacilityRefer"`
	City			string
	Province		string
	Latitude		float64
	Longitude		float64
	ZoomLevel		int
	CreatedAt		time.Time
	UpdatedAt		time.Time
	DeletedAt		*time.Time		`sql:index`
}

type HotelFacility struct{
	Id					int
	HotelFacilityRefer	uint
	AC					bool
	Parking				bool
	Receptionist		bool
	Pool				bool
	Lift				bool
	Restaurant			bool
	Wifi				bool
	Spa					bool
	CreatedAt 			time.Time
	UpdatedAt			time.Time
	DeletedAt 			*time.Time
}

type HotelRoom struct {
	Id				int			`gorm: primary_key`
	HotelRoomRefer	uint
	Name			string
	MaxGuest		int
	RoomSize		int
	BedDetail		string
	Breakfast		bool
	Wifi			bool
	Price			int
}

func init(){
	db, err:= database.Connect()

	if err != nil{
		panic(err)
	}
	defer db.Close()

	db.DropTableIfExists(&HotelRoom{})
	db.DropTableIfExists(&HotelFacility{})
	db.DropTableIfExists(&Hotel{})

	db.AutoMigrate(&Hotel{})
	db.AutoMigrate(&HotelFacility{})
	db.AutoMigrate(&HotelRoom{})
	db.Model(&HotelFacility{}).AddForeignKey("hotel_facility_refer","hotels(id)","CASCADE","CASCADE")
	db.Model(&HotelRoom{}).AddForeignKey("hotel_room_refer","hotels(id)","CASCADE","CASCADE")

	db.Create(&Hotel{
		Id:			1,
		Name:       "RedDoorz GBK Senayan",
		Address:    "Jl. Palmerah Barat No.110, Palmerah, Kec. Palmerah, Jakarta Barat, Jakarta 11480",
		Rating:     9.3,
		Star:		3,
		LocationRate: 89,
		CleanRate:	  90,
		RoomRate:	  86,
		ServiceRate:  88,
		City:       "Jakarta",
		Province:   "DKI Jakarta",
		Latitude:   -6.2069054,
		Longitude:  106.7209194,
		ZoomLevel:	15,
	})

	db.Create(&HotelFacility{
		HotelFacilityRefer: 1,
		AC:                 true,
		Parking:            true,
		Receptionist:       true,
		Pool:               false,
		Lift:               true,
		Restaurant:         true,
		Wifi:               true,
		Spa:                false,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 1,
		Name:           "Standard Room",
		MaxGuest:       1,
		RoomSize:       15,
		BedDetail:      "1 Queen Bed",
		Breakfast:      false,
		Wifi:           true,
		Price:          275000,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 1,
		Name:           "Double Room",
		MaxGuest:       2,
		RoomSize:       25,
		BedDetail:      "2 Single Bed",
		Breakfast:      false,
		Wifi:           true,
		Price:          354000,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 1,
		Name:           "VIP Room",
		MaxGuest:       3,
		RoomSize:       25,
		BedDetail:      "1 Kings Bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          730000,
	})


	db.Create(&Hotel{
		Id:			2,
		Name:       "Kempinski Hotel Jakarta",
		Address:    "Jalan M.H. Thamrin No.1, RT.1/RW.5, Menteng, Menteng, Jakarta Pusat, 10310",
		Rating:     9.5,
		Star:		5,
		LocationRate: 99,
		CleanRate:	  97,
		RoomRate:	  99,
		ServiceRate:  93,
		City:       "Jakarta",
		Province:   "DKI Jakarta",
		Latitude:   -6.2009372,
		Longitude:  106.7853241,
		ZoomLevel:	14,
	})

	db.Create(&HotelFacility{
		HotelFacilityRefer: 2,
		AC:                 true,
		Parking:            true,
		Receptionist:       true,
		Pool:               true,
		Lift:               true,
		Restaurant:         true,
		Wifi:               true,
		Spa:                false,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 2,
		Name:           "Deluxe Room",
		MaxGuest:       2,
		RoomSize:       44,
		BedDetail:      "Large Bed(King size)",
		Breakfast:      false,
		Wifi:           true,
		Price:          3618000,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 2,
		Name:           "Grand Deluxe Room",
		MaxGuest:       2,
		RoomSize:       64,
		BedDetail:      "Twin Bed, large bed (King size)",
		Breakfast:      false,
		Wifi:           true,
		Price:          4101000,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 2,
		Name:           "Executive Grand Deluxe Room",
		MaxGuest:       2,
		RoomSize:       64,
		BedDetail:      "Twin bed, Large bed (King size)",
		Breakfast:      true,
		Wifi:           true,
		Price:          4827000,
	})


	db.Create(&Hotel{
		Id:			3,
		Name:       "The Media Hotel Towers",
		Address:    "Jalan Gunung Sahari Raya No. 3",
		Rating:     7.8,
		Star:		5,
		LocationRate: 86,
		CleanRate:	  70,
		RoomRate:	  70,
		ServiceRate:  71,
		City:       "Jakarta",
		Province:   "DKI Jakarta",
		Latitude:   -6.1702381,
		Longitude:  106.7752316,
		ZoomLevel:	13,
	})

	db.Create(&HotelFacility{
		HotelFacilityRefer: 3,
		AC:                 true,
		Parking:            true,
		Receptionist:       true,
		Pool:               true,
		Lift:               true,
		Restaurant:         true,
		Wifi:               true,
		Spa:                true,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 3,
		Name:           "Premium Room",
		MaxGuest:       2,
		RoomSize:       32,
		BedDetail:      "Double bed, Twin bed",
		Breakfast:      false,
		Wifi:           true,
		Price:           838530,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 3,
		Name:           "Premium Queen Room",
		MaxGuest:       2,
		RoomSize:       32,
		BedDetail:      "Double bed, Twin bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          1573000,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 3,
		Name:           "Tower Room",
		MaxGuest:       2,
		RoomSize:       64,
		BedDetail:      "Double bed, Twin bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          1936000,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 3,
		Name:           "Club Suite Room",
		MaxGuest:       2,
		RoomSize:       64,
		BedDetail:      "Double bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          2202000,
	})

	db.Create(&Hotel{
		Id:			4,
		Name:       "Aryaduta Suite Semanggi",
		Address:    "Garnisun Dalam No.8 Street Karet Semanggi",
		Rating:     8.7,
		Star:		4,
		LocationRate: 95,
		CleanRate:	  78,
		RoomRate:	  83,
		ServiceRate:  85,
		City:       "Jakarta",
		Province:   "DKI Jakarta",
		Latitude:   -6.2117281,
		Longitude:  106.7837998,
		ZoomLevel:	14,
	})

	db.Create(&HotelFacility{
		HotelFacilityRefer: 4,
		AC:                 true,
		Parking:            true,
		Receptionist:       true,
		Pool:               true,
		Lift:               true,
		Restaurant:         true,
		Wifi:               true,
		Spa:                true,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 4,
		Name:           "One Bedroom Suite",
		MaxGuest:       2,
		RoomSize:       115,
		BedDetail:      "Double bed",
		Breakfast:      false,
		Wifi:           true,
		Price:          1026000,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 4,
		Name:           "One Bedroom Suite",
		MaxGuest:       2,
		RoomSize:       115,
		BedDetail:      "Double bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          1121000,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 4,
		Name:           "Two Bedroom Suite",
		MaxGuest:       3,
		RoomSize:       125,
		BedDetail:      "Double bed, Single bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          1225500,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 4,
		Name:           "Three Bedroom Suite",
		MaxGuest:       4,
		RoomSize:       135,
		BedDetail:      "Double bed, Single bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          1320500,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 4,
		Name:           "Penthouse",
		MaxGuest:       4,
		RoomSize:       435,
		BedDetail:      "Double bed, Twin bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          4750000,
	})

	db.Create(&Hotel{
		Id:			5,
		Name:       "Artotel Thamrin Jakarta",
		Address:    "Jl. Sunda No. 3 - Thamrin Jakarta pusat",
		Rating:     8.7,
		Star:		3,
		LocationRate: 95,
		CleanRate:	  89,
		RoomRate:	  83,
		ServiceRate:  88,
		City:       "Jakarta",
		Province:   "DKI Jakarta",
		Latitude:   -6.1962685,
		Longitude:  106.7861597,
		ZoomLevel:	14,
	})

	db.Create(&HotelFacility{
		HotelFacilityRefer: 5,
		AC:                 true,
		Parking:            true,
		Receptionist:       true,
		Pool:               false,
		Lift:               true,
		Restaurant:         true,
		Wifi:               true,
		Spa:                true,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 5,
		Name:           "Studio 20 Room",
		MaxGuest:       2,
		RoomSize:       20,
		BedDetail:      "Double bed, Twin bed",
		Breakfast:      false,
		Wifi:           true,
		Price:          880000,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 5,
		Name:           "Studio 20 Room",
		MaxGuest:       2,
		RoomSize:       20,
		BedDetail:      "Double bed, Twin bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          930000,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 5,
		Name:           "Studio 40 Room",
		MaxGuest:       2,
		RoomSize:       40,
		BedDetail:      "Double bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          1550000,
	})


	db.Create(&Hotel{
		Id:			6,
		Name:       "Oasis Amir Hotel",
		Address:    "Jalan Senen Raya Kav. 135 - 137",
		Rating:    	7.5,
		Star:		4,
		LocationRate: 83,
		CleanRate:	  59,
		RoomRate:	  59,
		ServiceRate:  62,
		City:       "Jakarta",
		Province:   "DKI Jakarta",
		Latitude:   -6.1912715,
		Longitude:  106.7809346,
		ZoomLevel:	13,
	})

	db.Create(&HotelFacility{
		HotelFacilityRefer: 6,
		AC:                 true,
		Parking:            true,
		Receptionist:       false,
		Pool:               false,
		Lift:               true,
		Restaurant:         true,
		Wifi:               true,
		Spa:                false,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 6,
		Name:           "Superior Double Room",
		MaxGuest:       2,
		RoomSize:       27,
		BedDetail:      "Double bed",
		Breakfast:      false,
		Wifi:           true,
		Price:          617500,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 6,
		Name:           "Superior Twin Room",
		MaxGuest:       2,
		RoomSize:       27,
		BedDetail:      "Double bed",
		Breakfast:      false,
		Wifi:           true,
		Price:          617500,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 6,
		Name:           "Superior Double Room",
		MaxGuest:       2,
		RoomSize:       27,
		BedDetail:      "Double bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          712500,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 6,
		Name:           "Superior Twin Room",
		MaxGuest:       2,
		RoomSize:       27,
		BedDetail:      "Double bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          712500,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 6,
		Name:           "Deluxe Room",
		MaxGuest:       2,
		RoomSize:       29,
		BedDetail:      "Double bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          760000,
	})

	db.Create(&Hotel{
		Id:			7,
		Name:       "Ashley Hotel Sabang",
		Address:    "Jl. H. Agus Salim No.22ab, RT.2/RW.1, Kb. Sirih, Menteng, Jakarta 10340, Indonesia",
		Rating:    	8.6,
		Star:		4,
		LocationRate: 90,
		CleanRate:	  82,
		RoomRate:	  76,
		ServiceRate:  78,
		City:       "Jakarta",
		Province:   "DKI Jakarta",
		Latitude:   -6.1946674,
		Longitude:  106.7861138,
		ZoomLevel:	14,
	})

	db.Create(&HotelFacility{
		HotelFacilityRefer: 7,
		AC:                 true,
		Parking:            true,
		Receptionist:       true,
		Pool:               false,
		Lift:               true,
		Restaurant:         true,
		Wifi:               true,
		Spa:                false,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 7,
		Name:           "Executive Double Room",
		MaxGuest:       2,
		RoomSize:       25,
		BedDetail:      "Double bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          780000,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 7,
		Name:           "Deluxe Double Room",
		MaxGuest:       2,
		RoomSize:       20,
		BedDetail:      "Double bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          910000,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 7,
		Name:           "Deluxe Twin Room",
		MaxGuest:       2,
		RoomSize:       25,
		BedDetail:      "Twin bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          975000,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 7,
		Name:           "Ashley Room",
		MaxGuest:       2,
		RoomSize:       31,
		BedDetail:      "Double bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          1105000,
	})

	db.Create(&Hotel{
		Id:			8,
		Name:       "Century Park Hotel",
		Address:    "Jalan Pintu Satu Senayan, Daerah Khusus Ibukota Jakarta 10270, Indonesia",
		Rating:    	8.8,
		Star:		4,
		LocationRate: 96,
		CleanRate:	  91,
		RoomRate:	  85,
		ServiceRate:  88,
		City:       "Jakarta",
		Province:   "DKI Jakarta",
		Latitude:   -6.2123667,
		Longitude:  106.780622,
		ZoomLevel:	14,
	})

	db.Create(&HotelFacility{
		HotelFacilityRefer: 8,
		AC:                 true,
		Parking:            true,
		Receptionist:       false,
		Pool:               true,
		Lift:               true,
		Restaurant:         true,
		Wifi:               true,
		Spa:                false,
	})

	db.Create(&HotelRoom{

		HotelRoomRefer: 8,
		Name:           "Deluxe Twin",
		MaxGuest:       2,
		RoomSize:       40,
		BedDetail:      "Twin bed",
		Breakfast:      false,
		Wifi:           true,
		Price:          1030000,
	})
	db.Create(&HotelRoom{

		HotelRoomRefer: 8,
		Name:           "Deluxe King",
		MaxGuest:       2,
		RoomSize:       40,
		BedDetail:      "Double bed",
		Breakfast:      true,
		Wifi:           true,
		Price:          1180000,
	})
	db.Create(&HotelRoom{

		HotelRoomRefer: 8,
		Name:           "Executive Twin",
		MaxGuest:       2,
		RoomSize:       40,
		BedDetail:      "Twin bed",
		Breakfast:      false,
		Wifi:           true,
		Price:          1330000,
	})


}

var recenthotels []Hotel

func GetHotels(city string)([]Hotel,error){

	db,err:=database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()

	var hotels []Hotel

	db.Where("city IN (?)",city).Preload("Rooms").Preload("Facilities").Find(&hotels)

	recenthotels = hotels


	return hotels,nil
}


