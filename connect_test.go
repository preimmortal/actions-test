package main

import (
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestDBConnection(t *testing.T) {
	dsn := "host=localhost user=pre password=pre dbname=test port=5432 sslmode=disable TimeZone=America/Los_Angeles"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal("Could not connect to Database")
	}
	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "D42", Price: 100})
	t.Log("Finished Database Connection Test")
}
