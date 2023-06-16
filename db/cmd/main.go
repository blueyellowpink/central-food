package main

import (
	"fmt"

	database "db"
	"db/models"
)

func main() {
	dbConfig := database.GormDatabaseConfig{
		Host:                   "localhost",
		User:                   "postgres",
		Password:               "example",
		Port:                   5432,
		DbName:                 "test",
		SslMode:                false,
		DisableImplicitPrepare: true,
	}
	db := database.NewGormDatabase(dbConfig)
	fmt.Println("Database connected")

	if err := db.Migrate(&models.Person{}); err != nil {
		fmt.Println("Migration Error")
	}

	// one := person.Person{Name: "test", Age: 15}
	// if err := db.Create(&one); err != nil {
	//     fmt.Println("Create Error")
	// }
	// fmt.Println(one)

	// many := []models.Person{
	// 	{Name: "1", Age: 1},
	// 	{Name: "2", Age: 2},
	// }
	// if err := db.CreateMany(many); err != nil {
	// 	fmt.Println("Create many error")
	// }
	// for _, i := range many {
	// 	fmt.Println(i)
	// }
}
