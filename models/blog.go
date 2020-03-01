package models

import "github.com/abenbyy/go-backend/database"

type Blog struct{
	Id 			int  `gorm:primary_key`
	Title		string
	Category 	string
	Content 	string	`gorm:"type:text;"`
	Image   	string  `gorm:"type:text;"`
	Viewer		int
}

func init(){
	db, err := database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()

	db.DropTableIfExists(&Blog{})
	db.AutoMigrate(&Blog{})

	SeedBlogData()
}

func GetBlog(id int)(Blog, error){

	db, err:= database.Connect()
	if err !=nil{
		panic(err)
	}

	defer db.Close()

	var blog Blog
	db.Where("id = ?",id).First(&blog)

	return blog, nil
}
func GetBlogs()([]Blog, error){
	db, err := database.Connect()

	if err !=nil{
		panic(err)
	}

	defer db.Close()

	var blogs []Blog

	db.Find(&blogs)

	return blogs, err

}

func GetPopularBlogs()([]Blog, error){
	db, err := database.Connect()

	if err != nil{
		panic(err)
	}

	defer db.Close()

	var blogs []Blog

	db.Order("viewer desc").Limit(5).Find(&blogs)

	return blogs, err

}

func CreateBlog(title string, category string, content string, image string)(error){
	db, err:= database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	db.Save(&Blog{
		Title:   title,
		Category: category,
		Content: content,
		Image:   image,
		Viewer:  0,
	})

	return nil
}

func DeleteBlog(id int){
	db, err:= database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	db.Where("id = ?",id).Delete(&Blog{})



}


func SeedBlogData(){

	db, err := database.Connect()

	if err !=nil{
		panic(err)
	}

	defer db.Close()

	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/tempat-wisata-prancis.jpg",
		Title:   "Bonjour! Ini Dia 5 Tempat Wisata Fantastis di Prancis, Kerennya Nggak Abis-Abis!",
		Category: "Travel",
		Content: "Hingga kini, Prancis memang masih menjadi destinasi wisata yang didambakan oleh banyak orang. Bukan hanya karena menara Eiffel yang ikonik, namun juga karena negara ini menjadi negara yang romantis dan tersimpan banyak hal-hal unik di dalamnya. Jika kamu ingin menjadikan Prancis sebagai destinasi liburan kamu bersama dengan orang tersayang. Tempat wisata di Prancis berikut ini bisa menjadi referensi tempat yang bisa kamu kunjungi selama di sana.",
		Viewer:  341,
	})

	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/wisata-pemandangan-terindah.jpg",
		Title:   "15 Pemandangan Alam Terindah di Indonesia, Indahnya Bikin Mata Terpana!",
		Category: "Travel",
		Content: "Kamu suka jalan-jalan melihat keindahan pemandangan alam di Indonesia? Di internet, banyak banget gambar pemandangan alam di Indonesia yang pastinya bikin kamu kebayang-bayang buat menjejakkan kaki secara langsung di sana. Ngomong-ngomong soal pemandangan alam Indonesia, kalau kamu sendiri lebih suka pemandangan alam berupa obyek wisata apa, sobat tiket? Apakah kamu lebih suka danau, laut, pegunungan, atau situs bersejarah bernilai seni tinggi? Nah, kalau kamu mau cuci mata sekalian mencari referensi tempat wisata di Indonesia yang oke punya, simak deh 15 gambar pemandangan alam di Indonesia yang ngehits banget ini. Kira-kira tempat mana yang yang bakal kamu kunjungi pas liburan nanti?",
		Viewer:  120,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/kuliner-labuan-bajo.jpg",
		Title:   "Ayo Cicipi Kuliner di Labuan Bajo yang Menggoda, Dijamin Ngiler!",
		Category: "Culinary",
		Content: "Wisata Kuliner di Labuan Bajo – Tentunya, nggak cukup kalau eksplor satu kota baru tanpa mencoba kulinernya. Di Labuan Bajo, sobat tiket bisa menikmati makanan khas Flores ataupun makanan autentik negara lain, sambil memanjakan mata dengan panorama alam yang indah.",
		Viewer:  221,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/tempat-wisata-prancis.jpg",
		Title:   "Bonjour! Ini Dia 5 Tempat Wisata Fantastis di Prancis, Kerennya Nggak Abis-Abis!",
		Category: "Travel",
		Content: "Hingga kini, Prancis memang masih menjadi destinasi wisata yang didambakan oleh banyak orang. Bukan hanya karena menara Eiffel yang ikonik, namun juga karena negara ini menjadi negara yang romantis dan tersimpan banyak hal-hal unik di dalamnya. Jika kamu ingin menjadikan Prancis sebagai destinasi liburan kamu bersama dengan orang tersayang. Tempat wisata di Prancis berikut ini bisa menjadi referensi tempat yang bisa kamu kunjungi selama di sana.",
		Viewer:  110,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/wisata-seoul.jpg",
		Title:   "10 Tempat Wisata di Seoul yang Pas untuk Liburan Bareng Keluarga",
		Category: "Travel",
		Content: "Saat ini, liburan bareng keluarga ke luar negeri semakin populer. Sobat tiket pasti juga setuju karena kamu bisa dapetin tiket pesawat dengan harga murah di tiket.com. Nah, salah satu destinasi yang cukup populer untuk itu adalah Seoul, ibukota dari Korea Selatan.",
		Viewer:  17600,
	})

	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/wisata-labuan-bajo.jpg",
		Title:   "10 Obyek Wisata Keluarga Labuan Bajo yang Wajib Kamu Kunjungi",
		Category: "Travel",
		Content: "Wisata Keluarga di Labuan Bajo – Labuan Bajo menjadi salah satu destinasi wisata yang paling diminati oleh para pelancong karena keindahan alam dan lautnya. Obyek wisata Labuan Bajo sangat beragam, nggak Cuma Pulau Komodo aja ya, sobat tiket. Labuan Bajo sendiri terletak di Nusa Tenggara Barat, dan merupakan ibu kota dari Manggarai Barat.",
		Viewer:  16122,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/wisata-pemandangan-terindah.jpg",
		Title:   "15 Pemandangan Alam Terindah di Indonesia, Indahnya Bikin Mata Terpana!",
		Category: "Travel",
		Content: "Kamu suka jalan-jalan melihat keindahan pemandangan alam di Indonesia? Di internet, banyak banget gambar pemandangan alam di Indonesia yang pastinya bikin kamu kebayang-bayang buat menjejakkan kaki secara langsung di sana. Ngomong-ngomong soal pemandangan alam Indonesia, kalau kamu sendiri lebih suka pemandangan alam berupa obyek wisata apa, sobat tiket? Apakah kamu lebih suka danau, laut, pegunungan, atau situs bersejarah bernilai seni tinggi? Nah, kalau kamu mau cuci mata sekalian mencari referensi tempat wisata di Indonesia yang oke punya, simak deh 15 gambar pemandangan alam di Indonesia yang ngehits banget ini. Kira-kira tempat mana yang yang bakal kamu kunjungi pas liburan nanti?",
		Viewer:  120,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/wisata-pemandangan-terindah.jpg",
		Title:   "15 Pemandangan Alam Terindah di Indonesia, Indahnya Bikin Mata Terpana!",
		Category: "Travel",
		Content: "Kamu suka jalan-jalan melihat keindahan pemandangan alam di Indonesia? Di internet, banyak banget gambar pemandangan alam di Indonesia yang pastinya bikin kamu kebayang-bayang buat menjejakkan kaki secara langsung di sana. Ngomong-ngomong soal pemandangan alam Indonesia, kalau kamu sendiri lebih suka pemandangan alam berupa obyek wisata apa, sobat tiket? Apakah kamu lebih suka danau, laut, pegunungan, atau situs bersejarah bernilai seni tinggi? Nah, kalau kamu mau cuci mata sekalian mencari referensi tempat wisata di Indonesia yang oke punya, simak deh 15 gambar pemandangan alam di Indonesia yang ngehits banget ini. Kira-kira tempat mana yang yang bakal kamu kunjungi pas liburan nanti?",
		Viewer:  120,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/wisata-pemandangan-terindah.jpg",
		Title:   "15 Pemandangan Alam Terindah di Indonesia, Indahnya Bikin Mata Terpana!",
		Category: "Travel",
		Content: "Kamu suka jalan-jalan melihat keindahan pemandangan alam di Indonesia? Di internet, banyak banget gambar pemandangan alam di Indonesia yang pastinya bikin kamu kebayang-bayang buat menjejakkan kaki secara langsung di sana. Ngomong-ngomong soal pemandangan alam Indonesia, kalau kamu sendiri lebih suka pemandangan alam berupa obyek wisata apa, sobat tiket? Apakah kamu lebih suka danau, laut, pegunungan, atau situs bersejarah bernilai seni tinggi? Nah, kalau kamu mau cuci mata sekalian mencari referensi tempat wisata di Indonesia yang oke punya, simak deh 15 gambar pemandangan alam di Indonesia yang ngehits banget ini. Kira-kira tempat mana yang yang bakal kamu kunjungi pas liburan nanti?",
		Viewer:  120,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/wisata-pemandangan-terindah.jpg",
		Title:   "15 Pemandangan Alam Terindah di Indonesia, Indahnya Bikin Mata Terpana!",
		Category: "Travel",
		Content: "Kamu suka jalan-jalan melihat keindahan pemandangan alam di Indonesia? Di internet, banyak banget gambar pemandangan alam di Indonesia yang pastinya bikin kamu kebayang-bayang buat menjejakkan kaki secara langsung di sana. Ngomong-ngomong soal pemandangan alam Indonesia, kalau kamu sendiri lebih suka pemandangan alam berupa obyek wisata apa, sobat tiket? Apakah kamu lebih suka danau, laut, pegunungan, atau situs bersejarah bernilai seni tinggi? Nah, kalau kamu mau cuci mata sekalian mencari referensi tempat wisata di Indonesia yang oke punya, simak deh 15 gambar pemandangan alam di Indonesia yang ngehits banget ini. Kira-kira tempat mana yang yang bakal kamu kunjungi pas liburan nanti?",
		Viewer:  120,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/wisata-pemandangan-terindah.jpg",
		Title:   "15 Pemandangan Alam Terindah di Indonesia, Indahnya Bikin Mata Terpana!",
		Category: "Travel",
		Content: "Kamu suka jalan-jalan melihat keindahan pemandangan alam di Indonesia? Di internet, banyak banget gambar pemandangan alam di Indonesia yang pastinya bikin kamu kebayang-bayang buat menjejakkan kaki secara langsung di sana. Ngomong-ngomong soal pemandangan alam Indonesia, kalau kamu sendiri lebih suka pemandangan alam berupa obyek wisata apa, sobat tiket? Apakah kamu lebih suka danau, laut, pegunungan, atau situs bersejarah bernilai seni tinggi? Nah, kalau kamu mau cuci mata sekalian mencari referensi tempat wisata di Indonesia yang oke punya, simak deh 15 gambar pemandangan alam di Indonesia yang ngehits banget ini. Kira-kira tempat mana yang yang bakal kamu kunjungi pas liburan nanti?",
		Viewer:  120,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/wisata-pemandangan-terindah.jpg",
		Title:   "15 Pemandangan Alam Terindah di Indonesia, Indahnya Bikin Mata Terpana!",
		Category: "Travel",
		Content: "Kamu suka jalan-jalan melihat keindahan pemandangan alam di Indonesia? Di internet, banyak banget gambar pemandangan alam di Indonesia yang pastinya bikin kamu kebayang-bayang buat menjejakkan kaki secara langsung di sana. Ngomong-ngomong soal pemandangan alam Indonesia, kalau kamu sendiri lebih suka pemandangan alam berupa obyek wisata apa, sobat tiket? Apakah kamu lebih suka danau, laut, pegunungan, atau situs bersejarah bernilai seni tinggi? Nah, kalau kamu mau cuci mata sekalian mencari referensi tempat wisata di Indonesia yang oke punya, simak deh 15 gambar pemandangan alam di Indonesia yang ngehits banget ini. Kira-kira tempat mana yang yang bakal kamu kunjungi pas liburan nanti?",
		Viewer:  120,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/wisata-pemandangan-terindah.jpg",
		Title:   "15 Pemandangan Alam Terindah di Indonesia, Indahnya Bikin Mata Terpana!",
		Category: "Travel",
		Content: "Kamu suka jalan-jalan melihat keindahan pemandangan alam di Indonesia? Di internet, banyak banget gambar pemandangan alam di Indonesia yang pastinya bikin kamu kebayang-bayang buat menjejakkan kaki secara langsung di sana. Ngomong-ngomong soal pemandangan alam Indonesia, kalau kamu sendiri lebih suka pemandangan alam berupa obyek wisata apa, sobat tiket? Apakah kamu lebih suka danau, laut, pegunungan, atau situs bersejarah bernilai seni tinggi? Nah, kalau kamu mau cuci mata sekalian mencari referensi tempat wisata di Indonesia yang oke punya, simak deh 15 gambar pemandangan alam di Indonesia yang ngehits banget ini. Kira-kira tempat mana yang yang bakal kamu kunjungi pas liburan nanti?",
		Viewer:  120,
	})

	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/tempat-wisata-prancis.jpg",
		Title:   "Bonjour! Ini Dia 5 Tempat Wisata Fantastis di Prancis, Kerennya Nggak Abis-Abis!",
		Category: "Travel",
		Content: "Hingga kini, Prancis memang masih menjadi destinasi wisata yang didambakan oleh banyak orang. Bukan hanya karena menara Eiffel yang ikonik, namun juga karena negara ini menjadi negara yang romantis dan tersimpan banyak hal-hal unik di dalamnya. Jika kamu ingin menjadikan Prancis sebagai destinasi liburan kamu bersama dengan orang tersayang. Tempat wisata di Prancis berikut ini bisa menjadi referensi tempat yang bisa kamu kunjungi selama di sana.",
		Viewer:  110,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/tempat-wisata-prancis.jpg",
		Title:   "Bonjour! Ini Dia 5 Tempat Wisata Fantastis di Prancis, Kerennya Nggak Abis-Abis!",
		Category: "Travel",
		Content: "Hingga kini, Prancis memang masih menjadi destinasi wisata yang didambakan oleh banyak orang. Bukan hanya karena menara Eiffel yang ikonik, namun juga karena negara ini menjadi negara yang romantis dan tersimpan banyak hal-hal unik di dalamnya. Jika kamu ingin menjadikan Prancis sebagai destinasi liburan kamu bersama dengan orang tersayang. Tempat wisata di Prancis berikut ini bisa menjadi referensi tempat yang bisa kamu kunjungi selama di sana.",
		Viewer:  110,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/tempat-wisata-prancis.jpg",
		Title:   "Bonjour! Ini Dia 5 Tempat Wisata Fantastis di Prancis, Kerennya Nggak Abis-Abis!",
		Category: "Travel",
		Content: "Hingga kini, Prancis memang masih menjadi destinasi wisata yang didambakan oleh banyak orang. Bukan hanya karena menara Eiffel yang ikonik, namun juga karena negara ini menjadi negara yang romantis dan tersimpan banyak hal-hal unik di dalamnya. Jika kamu ingin menjadikan Prancis sebagai destinasi liburan kamu bersama dengan orang tersayang. Tempat wisata di Prancis berikut ini bisa menjadi referensi tempat yang bisa kamu kunjungi selama di sana.",
		Viewer:  110,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/tempat-wisata-prancis.jpg",
		Title:   "Bonjour! Ini Dia 5 Tempat Wisata Fantastis di Prancis, Kerennya Nggak Abis-Abis!",
		Category: "Travel",
		Content: "Hingga kini, Prancis memang masih menjadi destinasi wisata yang didambakan oleh banyak orang. Bukan hanya karena menara Eiffel yang ikonik, namun juga karena negara ini menjadi negara yang romantis dan tersimpan banyak hal-hal unik di dalamnya. Jika kamu ingin menjadikan Prancis sebagai destinasi liburan kamu bersama dengan orang tersayang. Tempat wisata di Prancis berikut ini bisa menjadi referensi tempat yang bisa kamu kunjungi selama di sana.",
		Viewer:  110,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/tempat-wisata-prancis.jpg",
		Title:   "Bonjour! Ini Dia 5 Tempat Wisata Fantastis di Prancis, Kerennya Nggak Abis-Abis!",
		Category: "Travel",
		Content: "Hingga kini, Prancis memang masih menjadi destinasi wisata yang didambakan oleh banyak orang. Bukan hanya karena menara Eiffel yang ikonik, namun juga karena negara ini menjadi negara yang romantis dan tersimpan banyak hal-hal unik di dalamnya. Jika kamu ingin menjadikan Prancis sebagai destinasi liburan kamu bersama dengan orang tersayang. Tempat wisata di Prancis berikut ini bisa menjadi referensi tempat yang bisa kamu kunjungi selama di sana.",
		Viewer:  110,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/tempat-wisata-prancis.jpg",
		Title:   "Bonjour! Ini Dia 5 Tempat Wisata Fantastis di Prancis, Kerennya Nggak Abis-Abis!",
		Category: "Travel",
		Content: "Hingga kini, Prancis memang masih menjadi destinasi wisata yang didambakan oleh banyak orang. Bukan hanya karena menara Eiffel yang ikonik, namun juga karena negara ini menjadi negara yang romantis dan tersimpan banyak hal-hal unik di dalamnya. Jika kamu ingin menjadikan Prancis sebagai destinasi liburan kamu bersama dengan orang tersayang. Tempat wisata di Prancis berikut ini bisa menjadi referensi tempat yang bisa kamu kunjungi selama di sana.",
		Viewer:  110,
	})



}


