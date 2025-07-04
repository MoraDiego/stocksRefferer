package main

import (
	"fmt"
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
}
