// internal/database/database.go
package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

// Config armazena as configurações de conexão com o banco de dados
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewConnection cria uma nova conexão com o banco de dados
func NewConnection() (*sql.DB, error) {
	config := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com o banco de dados: %v", err)
	}

	// Configurar o pool de conexões
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Testar a conexão
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao pingar o banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso")

	return db, nil
}
