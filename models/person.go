package models

type Person struct {
	ModelBase
	FirstName string `json:"first_name" gorm:"type:string;size:100;not null" validate:"required,min=0,max=100"`
	LastName  string `json:"last_name" gorm:"type:string;size:100;not null" validate:"required,min=0,max=100"`
	Address   string `json:"address" gorm:"type:string;size:255;not null" validate:"required,min=0,max=255"`
}
