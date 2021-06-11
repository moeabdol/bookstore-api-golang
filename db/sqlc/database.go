package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/moeabdol/bookstore-api-golang/utils"
)

var DB *Store

// ConnectToDatabase function
func ConnectToDatabase() error {
	conn, err := sql.Open(
		utils.Config.DbDialect,
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			utils.Config.DbHost,
			utils.Config.DbPort,
			utils.Config.DbUser,
			utils.Config.DbPassword,
			utils.Config.DbName,
			utils.Config.DbSslMode,
		),
	)
	if err != nil {
		return err
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	DB = NewStore(conn)
	return nil
}
