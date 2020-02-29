package models

import (
	"github.com/abenbyy/go-backend/database"
	"time"
)

type Promo struct{
	Id		int  `gorm:primary_key`
	Title	string
	Content	string
	Image	string
	Code	string
	Description string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

}

func init(){
	db, err := database.Connect()
	if err!=nil{
		panic(err)

	}

	defer db.Close()

	db.DropTableIfExists(&Promo{})
	db.AutoMigrate(&Promo{})

	SeedPromoData()

}

func GetPromos()([]Promo, error){
	db, err:= database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	var promos []Promo

	db.Find(&promos)

	return promos, nil
}
func GetPromo(id int)(Promo,error){
	db, err:= database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()

	var promo Promo

	db.Where("id = ?",id).First(&promo)

	return promo, nil

}

func GetOtherPromo(id int)([]Promo, error){

	db, err :=database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	var promos []Promo

	db.Not("id = ?",id).Find(&promos)

	return promos,nil

}

func SeedPromoData(){

	db, err:= database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()

	db.Create(&Promo{
		Title:     "Diskon hingga Rp200.000 untuk Tiket Pesawat Domestik",
		Content:   "Saatnya jelajahi tempat-tempat eksotis di Indonesia! Mau terbang ke semua rute domestik dengan maskapai Citilink, Garuda Indonesia, Transnusa, Sriwijaya, atau NAM Air dapat diskon hingga Rp200.000 kalau menggunakan Debit Online Jenius atau Kartu Kredit BNI di setiap transaksi kamu. Dapatkan promo ini hanya di aplikasi tiket.com Baca Syarat dan Ketentuan di bawah, ya!",
		Image:     "../../../assets/promoimages/diskon-tiket-pesawat-domestik-200.jpg",
		Description: "Diskon Rp200.000",
		Code:	   "PAYDAYBNI1",
	})

	db.Create(&Promo{
		Title:     "Diskon Hotel di Jepang Rp150.000",
		Content:   "Saatnya bikin mimpi liburan ke Jepang jadi kenyataan~ tiket.com punya caranya biar liburan kamu makin hemat, lho. Nikmati diskon Rp150.000 untuk hotel pilihan di Jepang! Pilih hotel favoritmu dan pesan pake Kartu Kredit Citibank untuk dapetin diskonnya. Baca Syarat dan Ketentuan di bawah ini, ya!",
		Image:     "../../../assets/promoimages/diskon-hotel-jepang-150.jpg",
		Description: "Hotel di Jepang Diskon Rp150.000",
		Code:	   "CITIVACAYJAPAN",
	})



	db.Create(&Promo{
		Title:     "Diskon hingga Rp500.000 untuk Hotel Domestik",
		Content:   "Bayangin deh, Indonesia yang cantik ini memanggil nama kamu untuk menjelajahi kekayaannya. Menikmati pemandangan gunungnya, melihat senja di pantai, atau hanya sekadar menyaksikan bintang-bintang. Nggak perlu ragu lagi untuk liburan karena ada diskon hingga Rp500.000 untuk nginep di hotel domestik bareng keluarga atau sahabat kamu. Nikmati diskonnya hanya di aplikasi tiket.com Baca Syarat dan Ketentuan di bawah, ya!",
		Image:     "../../../assets/promoimages/diskon-hotel-domestik-500.jpg",
		Description: "Diskon 7% hingga Rp500.000",
		Code:	   "GAJIANHEMAT",
	})

	db.Create(&Promo{
		Title:     "Diskon hingga Rp350.000 untuk Terbang ke Rute Internasional",
		Content:   "Coba sebutin negara mana aja yang pengin kamu kunjungi selama ini? Jepang? Korea? Paris? Ke mana aja bisa~ Makin hemat juga karena ada diskon hingga Rp350.000 untuk penerbangan ke rute internasional yang kamu mau. Jangan nunggu lama lagi buat dapetin promo ini hanya di tiket.com Baca Syarat dan Ketentuan di bawah, ya!",
		Image:     "../../../assets/promoimages/diskon-tiket-pesawat-internasional-350.jpg",
		Description: "Diskon Rp350.000",
		Code:	   "PAYDAYSTANCHART",
	})


	db.Create(&Promo{
		Title:     "Diskon hingga Rp400.000 untuk Terbang ke Jepang",
		Content:   "Saatnya bikin liburan impian jadi kenyataan! Mau eksplor Tokyo, Sapporo, Kyoto, atau kota lainnya di Jepang? Nikmati diskon hingga Rp400.000 untuk tiket penerbangan ke Jepang bersama Japan Airlines, Nippon Airways, Garuda Indonesia, Cathay Pacific, Singapore Airlines, Philippine Airlines, Malaysia Airlines dan Thai Airways dengan menggunakan kartu kredit Citibank kamu. Untuk liburan seru di Jepang, #adatiketnya Baca Syarat dan Ketentuan di bawah ini, ya!",
		Image:     "../../../assets/promoimages/diskon-tiket-pesawat-jepang-400.jpg",
		Description: "Tiket Pesawat Internasional Diskon Rp400.000",
		Code:	   "JAPANCITI",
	})


	//JUST COPIES FOR DUMMY
	db.Create(&Promo{
		Title:     "Diskon hingga Rp400.000 untuk Terbang ke Jepang",
		Content:   "Saatnya bikin liburan impian jadi kenyataan! Mau eksplor Tokyo, Sapporo, Kyoto, atau kota lainnya di Jepang? Nikmati diskon hingga Rp400.000 untuk tiket penerbangan ke Jepang bersama Japan Airlines, Nippon Airways, Garuda Indonesia, Cathay Pacific, Singapore Airlines, Philippine Airlines, Malaysia Airlines dan Thai Airways dengan menggunakan kartu kredit Citibank kamu. Untuk liburan seru di Jepang, #adatiketnya Baca Syarat dan Ketentuan di bawah ini, ya!",
		Image:     "../../../assets/promoimages/diskon-tiket-pesawat-jepang-400.jpg",
		Description: "Tiket Pesawat Internasional Diskon Rp400.000",
		Code:	   "JAPANCITI",
	})
	db.Create(&Promo{
		Title:     "Diskon hingga Rp400.000 untuk Terbang ke Jepang",
		Content:   "Saatnya bikin liburan impian jadi kenyataan! Mau eksplor Tokyo, Sapporo, Kyoto, atau kota lainnya di Jepang? Nikmati diskon hingga Rp400.000 untuk tiket penerbangan ke Jepang bersama Japan Airlines, Nippon Airways, Garuda Indonesia, Cathay Pacific, Singapore Airlines, Philippine Airlines, Malaysia Airlines dan Thai Airways dengan menggunakan kartu kredit Citibank kamu. Untuk liburan seru di Jepang, #adatiketnya Baca Syarat dan Ketentuan di bawah ini, ya!",
		Image:     "../../../assets/promoimages/diskon-tiket-pesawat-jepang-400.jpg",
		Description: "Tiket Pesawat Internasional Diskon Rp400.000",
		Code:	   "JAPANCITI",
	})
	db.Create(&Promo{
		Title:     "Diskon hingga Rp400.000 untuk Terbang ke Jepang",
		Content:   "Saatnya bikin liburan impian jadi kenyataan! Mau eksplor Tokyo, Sapporo, Kyoto, atau kota lainnya di Jepang? Nikmati diskon hingga Rp400.000 untuk tiket penerbangan ke Jepang bersama Japan Airlines, Nippon Airways, Garuda Indonesia, Cathay Pacific, Singapore Airlines, Philippine Airlines, Malaysia Airlines dan Thai Airways dengan menggunakan kartu kredit Citibank kamu. Untuk liburan seru di Jepang, #adatiketnya Baca Syarat dan Ketentuan di bawah ini, ya!",
		Image:     "../../../assets/promoimages/diskon-tiket-pesawat-jepang-400.jpg",
		Description: "Tiket Pesawat Internasional Diskon Rp400.000",
		Code:	   "JAPANCITI",
	})


}



