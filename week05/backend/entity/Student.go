package entity

import "gorm.io/gorm"
import "time"

type Student struct {
	gorm.Model
	FristName 	string
	LastName 	string
	DateOfBirth time.Time
	Email 		string
	PhoneNumber	string

	//EnrollmentSchedule EnrollmentSchedule
}