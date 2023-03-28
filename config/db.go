package config

import "database/sql"

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "postgres"
	Dbname   = "perpus"
)

var Db *sql.DB
