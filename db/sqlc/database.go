package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/moeabdol/bookstore-api-golang/utils"
)

var DB *Store

// ConnectToDatabase function
func ConnectToDatabase() {
	conn, err := sql.Open(
		utils.Config.DBDialect,
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			utils.Config.DBHost,
			utils.Config.DBPort,
			utils.Config.DBUser,
			utils.Config.DBPassword,
			utils.Config.DBName,
			utils.Config.DBSslmode,
		),
	)
	if err != nil {
		utils.Log.Fatalf("Unable to connect to database %s", err)
	}

	err = conn.Ping()
	if err != nil {
		utils.Log.Fatalf("Unable to connect to database %s", err)
	}

	DB = NewStore(conn)
}
