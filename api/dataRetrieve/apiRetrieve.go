package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"stocksRefferer/models"
)

type APIResponse struct {
	Stocks   []models.Stock `json:"items"`
	NextPage string         `json:"next_page"`
}

func FetchAllItems() ([]models.Stock, error) {
	var baseURL = os.Getenv("API_URL")
	var token = os.Getenv("API_TOKEN")

	var allStocks []models.Stock
	page := ""

	client := &http.Client{}

	for {
		// Construir URL con parámetro de página si existe
		fullURL := baseURL
		if page != "" {
			fullURL = fmt.Sprintf("%s?next_page=%s", baseURL, url.QueryEscape(page))
		}
		fmt.Println(fullURL)
		//Preparar request con headers
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Accept", "application/json")

		// Enviar request
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		// Leer respuesta
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		// Parsear JSON
		var apiResp APIResponse
		if err := json.Unmarshal(body, &apiResp); err != nil {
			return nil, err
		}

		// Agregar items a lista acumulada
		allStocks = append(allStocks, apiResp.Stocks...)

		// Verificar si hay más páginas
		if apiResp.NextPage == "" {
			break
		}
		page = apiResp.NextPage
	}
	return allStocks, nil
}
