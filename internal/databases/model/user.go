package model

type User struct {
	Id        uint `gorm:"primary_key"`
	FirstName string
	LastName  string
}
