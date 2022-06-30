package app

import (
	"database/sql"
)

var (
	Db  *sql.DB
	Err error
)

type Config struct {
	Db_Host        	string `env:"DB_ADDRESS"`
	Db_Port        	int    `env:"DB_PORT"`
	Db_Username     string `env:"DB_USERNAME"`
	Db_Password    	string `env:"DB_PASSWORD"`
	Db_Dbname      	string `env:"DB_NAME"`
	PORT           	int    `env:"_PORT"`
}