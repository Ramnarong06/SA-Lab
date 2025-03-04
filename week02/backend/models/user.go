package models


import (

   "time"


   "gorm.io/gorm"

)


type Users struct {

   gorm.Model

   FirstName string    `json:"first_name"`

   LastName  string    `json:"last_name"`

   Email     string    `json:"email"`

   Age       uint8     `json:"age"`

   Address   string    `json:"address"`

   Password  string    `json:"-"`

   BirthDay  time.Time `json:"birthday"`

   GenderID  uint      `json:"gender_id"`

   Gender    *Genders  `gorm:"foreignKey: gender_id" json:"gender"`

}