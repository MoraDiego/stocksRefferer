package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env") // ajusta el path según dónde esté el archivo
	if err != nil {
		log.Fatalf("Error cargando .env: %v", err)
	}
	os.Exit(m.Run())
}

// TestGetAccionesHandler prueba que GetAccionesHandler responde con 200
func TestGetAccionesHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/acciones", nil)
	rr := httptest.NewRecorder()

	GetAccionesHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("esperado código 200, pero se obtuvo %d", status)
	}
}

// TestGetAccionesByRiskHandler_InvalidRisk prueba un riesgo no válido
func TestGetAccionesByRiskHandler_InvalidRisk(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/accionesPorRiesgo?risk=abc", nil)
	rr := httptest.NewRecorder()

	GetAccionesByRiskHandler(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("esperado código 400, pero se obtuvo %d", status)
	}
}

// TestGetAccionesByRiskHandler_OutOfRange prueba un riesgo fuera del rango
func TestGetAccionesByRiskHandler_OutOfRange(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/accionesPorRiesgo?risk=100", nil)
	rr := httptest.NewRecorder()

	GetAccionesByRiskHandler(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("esperado código 400, pero se obtuvo %d", status)
	}
}

// TestGetAccionesByRiskHandler_Valid prueba una petición válida
func TestGetAccionesByRiskHandler_Valid(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/accionesPorRiesgo?risk=1", nil)
	rr := httptest.NewRecorder()

	GetAccionesByRiskHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("esperado código 200, pero se obtuvo %d", status)
	}
}
