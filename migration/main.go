package main

import (
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading env file")
	}

	// DANGER : only use for migrate database !!
	// err = database.Connect().AutoMigrate(
	// 	&entities.User{},
	// 	&entities.UserAddress{},
	//
	// 	&entities.Class{},
	// 	&entities.Schedule{},
	// 	&entities.Booking{},
	// 	&entities.Trainer{},
	//
	// 	&entities.CreditType{},
	// 	&entities.Credit{},
	// 	&entities.CreditUsed{},
	// 	&entities.UserCredit{},
	//
	// 	&entities.Product{},
	//
	// 	&entities.Order{},
	// 	&entities.OrderItem{},
	// )
	// if err != nil {
	// 	fmt.Printf("Error Migration: %v\n", err)
	// }
}
