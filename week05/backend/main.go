package main

import (
	"example.com/sa-67-example/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("sa-67.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&entity.Student{}, &entity.Teacher{}, &entity.CourseType{}, &entity.Subject{}, &entity.EnrollmentSchedule{})
}