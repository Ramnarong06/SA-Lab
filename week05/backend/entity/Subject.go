package entity

import "gorm.io/gorm"

type Subject struct {
	gorm.Model
	SubjectName     string
	SubjectType     string
	TotalStudyHours float64
	Note            string

	CourseTypeID *uint
	CourseType   CourseType `gorm:"foreignKey:CourseTypeID"`

	EnrollmentSchedule []EnrollmentSchedule `gorm:"foreignKey:SubjectID"`
}
