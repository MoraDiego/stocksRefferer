package main

import (
	"fmt"
	"stocksRefferer/db"
	"stocksRefferer/models"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	fmt.Print(err)
}
func main() {
	items, err := FetchAllItems()
	if err != nil {
		panic(err)
	}
	for i := range items {
		err := models.ProcesarTargets(&items[i])
		if err != nil {
			fmt.Printf("Error al procesar item %d (%s): %v\n", i, items[i].Ticker, err)
		}
	}
	DB := db.Connect()
	DB.AutoMigrate(models.Stock{})
	if err := DB.CreateInBatches(&items, len(items)).Error; err != nil {
		fmt.Print("No insertados")
	}
}
