package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	// db configuration
	Username string
	Password string
	Port     string
	Address  string
	Database string

	// db connection
	DB *gorm.DB
}

type PsqlDb struct {
	*Postgres
}

var (
	PSQL *PsqlDb
)

func InitPostgres() error {
	PSQL = new(PsqlDb)

	PSQL.Postgres = &Postgres{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	// connect to database
	err := PSQL.Postgres.OpenConnection()
	if err != nil {
		return err
	}

	err = PSQL.DB.Raw("SELECT 1").Error
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) OpenConnection() error {
	// init dsn
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.Address, p.Port, p.Username, p.Password, p.Database)

	dbConnection, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	p.DB = dbConnection

	fmt.Println("Successfully connected to database")

	return nil
}
