package entity

import (
	"time"

	"gorm.io/gorm"
)

type EnrollmentSchedule struct {
	gorm.Model
	ScheduleDate    time.Time
	SubjectType     string
	StudyDurationHR float64

	SubjectID *uint
	Subject   Subject `gorm:"foreignKey:SubjectID"`

	TeacherID *uint
	Teacher   Teacher `gorm:"foreignKey:TeacherID"`

	StudentID *uint
}
