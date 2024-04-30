package repositories

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email                string `validate:"omitempty,optional_between=6 75,email" gorm:"type:varchar(75);not null;unique" json:"email"`
	Password             string `gorm:"not null" json:"-"`
	Name                 string `validate:"optional_between=1 100" gorm:"type:varchar(100);not null" json:"name"`
	Role                 int    `validate:"optional_between=1 6" gorm:"not null" json:"role"`
	Status               int    `validate:"optional_between=1 2" json:"status"`
	PlainPassword        string `validate:"optional_between=8 50" gorm:"-" json:"plainPassword"`
	PlainPasswordConfirm string `validate:"optional_between=8 50" gorm:"-" json:"plainPasswordConfirm"`
}
