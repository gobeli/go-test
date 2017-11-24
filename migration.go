package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type ToDo struct {
	Finished bool
	Text string
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&ToDo{})
	fmt.Println("Migration created")
	if err := db.First(&ToDo{}).Error; err != nil {
		todo := ToDo{Finished:false, Text: "My very first Todo"}
		db.Create(&todo)
		fmt.Println("Example created")
	}
}