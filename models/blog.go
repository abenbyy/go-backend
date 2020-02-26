package models

import "github.com/abenbyy/go-backend/database"

type Blog struct{
	Id 		int  `gorm:primary_key`
	Title	string
	Content string	`gorm:"type:text;"`
	Image   string
	Viewer	int
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

func CreateBlog(title string, content string, image string)(error){
	db, err:= database.Connect()

	if err!=nil{
		panic(err)
	}

	defer db.Close()

	db.Save(&Blog{
		Title:   title,
		Content: content,
		Image:   image,
		Viewer:  0,
	})

	return nil
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
		Content: "Hingga kini, Prancis memang masih menjadi destinasi wisata yang didambakan oleh banyak orang. Bukan hanya karena menara Eiffel yang ikonik, namun juga karena negara ini menjadi negara yang romantis dan tersimpan banyak hal-hal unik di dalamnya.<br><br> Jika kamu ingin menjadikan Prancis sebagai destinasi liburan kamu bersama dengan orang tersayang. Tempat wisata di Prancis berikut ini bisa menjadi referensi tempat yang bisa kamu kunjungi selama di sana.",
		Viewer:  341,
	})

	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/tempat-wisata-prancis.jpg",
		Title:   "Bonjour! Ini Dia 5 Tempat Wisata Fantastis di Prancis, Kerennya Nggak Abis-Abis!",
		Content: "Hingga kini, Prancis memang masih menjadi destinasi wisata yang didambakan oleh banyak orang. Bukan hanya karena menara Eiffel yang ikonik, namun juga karena negara ini menjadi negara yang romantis dan tersimpan banyak hal-hal unik di dalamnya.<br><br> Jika kamu ingin menjadikan Prancis sebagai destinasi liburan kamu bersama dengan orang tersayang. Tempat wisata di Prancis berikut ini bisa menjadi referensi tempat yang bisa kamu kunjungi selama di sana.",
		Viewer:  341,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/tempat-wisata-prancis.jpg",
		Title:   "Bonjour! Ini Dia 5 Tempat Wisata Fantastis di Prancis, Kerennya Nggak Abis-Abis!",
		Content: "Hingga kini, Prancis memang masih menjadi destinasi wisata yang didambakan oleh banyak orang. Bukan hanya karena menara Eiffel yang ikonik, namun juga karena negara ini menjadi negara yang romantis dan tersimpan banyak hal-hal unik di dalamnya.<br><br> Jika kamu ingin menjadikan Prancis sebagai destinasi liburan kamu bersama dengan orang tersayang. Tempat wisata di Prancis berikut ini bisa menjadi referensi tempat yang bisa kamu kunjungi selama di sana.",
		Viewer:  341,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/tempat-wisata-prancis.jpg",
		Title:   "Bonjour! Ini Dia 5 Tempat Wisata Fantastis di Prancis, Kerennya Nggak Abis-Abis!",
		Content: "Hingga kini, Prancis memang masih menjadi destinasi wisata yang didambakan oleh banyak orang. Bukan hanya karena menara Eiffel yang ikonik, namun juga karena negara ini menjadi negara yang romantis dan tersimpan banyak hal-hal unik di dalamnya.<br><br> Jika kamu ingin menjadikan Prancis sebagai destinasi liburan kamu bersama dengan orang tersayang. Tempat wisata di Prancis berikut ini bisa menjadi referensi tempat yang bisa kamu kunjungi selama di sana.",
		Viewer:  341,
	})
	db.Create(&Blog{
		Image:   "../../../assets/blogthumbnails/tempat-wisata-prancis.jpg",
		Title:   "Bonjour! Ini Dia 5 Tempat Wisata Fantastis di Prancis, Kerennya Nggak Abis-Abis!",
		Content: "Hingga kini, Prancis memang masih menjadi destinasi wisata yang didambakan oleh banyak orang. Bukan hanya karena menara Eiffel yang ikonik, namun juga karena negara ini menjadi negara yang romantis dan tersimpan banyak hal-hal unik di dalamnya.<br><br> Jika kamu ingin menjadikan Prancis sebagai destinasi liburan kamu bersama dengan orang tersayang. Tempat wisata di Prancis berikut ini bisa menjadi referensi tempat yang bisa kamu kunjungi selama di sana.",
		Viewer:  341,
	})


}


