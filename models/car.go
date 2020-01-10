package models

import "time"

type Car struct {
	Id			int				`gorm:primary_key`
	Brand		string			`gorm:"type:varchar(100);not null"`
	Model		string			`gorm:"type:varchar(100);not null"`
	Price		int				`gorm:int`
	VendorId	int				`gorm:foreignkey`
	Passenger	int 			`gorm:int`
	Luggage		int				`gorm:int`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time		`sql:index`
}