package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *Queries

// ConnectToDatabase function
func ConnectToDatabase() error {
	dbDialect := viper.Get("DB_DIALECT").(string)
	dbHost := viper.Get("DB_HOST").(string)
	dbPort := viper.Get("DB_PORT").(string)
	dbSslMode := viper.Get("DB_SSLMODE").(string)
	dbName := viper.Get("DB_NAME").(string)
	dbUser := viper.Get("DB_USER").(string)
	dbPassword := viper.Get("DB_PASSWORD").(string)

	conn, err := sql.Open(
		dbDialect,
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			dbHost,
			dbPort,
			dbUser,
			dbPassword,
			dbName,
			dbSslMode,
		),
	)
	if err != nil {
		return err
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	DB = New(conn)
	return nil
}
