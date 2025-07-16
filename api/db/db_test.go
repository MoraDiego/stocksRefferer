package db

import (
	"log"
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

func TestConnect_ReturnsDBInstance(t *testing.T) {

	conn := Connect()

	if conn == nil {
		t.Fatal("Connect() retornó nil, se esperaba una instancia *gorm.DB")
	}

	sqlDB, err := conn.DB()
	if err != nil {
		t.Fatalf("Error al obtener *sql.DB: %v", err)
	}

	// Intentar hacer ping a la base de datos
	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("No se pudo hacer ping a la base de datos: %v", err)
	}
}
