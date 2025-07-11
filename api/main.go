package main

import (
	"fmt"
	"net/http"
	"stocksRefferer/handlers"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	fmt.Print(err)
}
func main() {
	http.HandleFunc("/acciones", handlers.GetAccionesHandler)
	http.HandleFunc("/accionesRiesgo", handlers.GetAccionesByRiskHandler)
	fmt.Print(http.ListenAndServe(":8080", nil))
}
