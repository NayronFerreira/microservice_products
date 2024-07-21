package database

import (
	"database/sql"
	"fmt"

	"github.com/NayronFerreira/microservice_products/configs"
	_ "github.com/go-sql-driver/mysql"
)

func SetupDB(config *configs.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := sql.Open(config.DBDriver, dsn)
	if err != nil {
		return nil, err
	}

	if err := setupTables(db, config); err != nil {
		return nil, err
	}

	return db, nil
}

func setupTables(db *sql.DB, config *configs.Config) error {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + config.DBName)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS %s (
        id VARCHAR(255) PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        model VARCHAR(255) NOT NULL,
        code VARCHAR(50) NOT NULL,
        price DECIMAL(10,2) NOT NULL,
        color VARCHAR(50) NOT NULL
    );
    `, config.DBTable)

	_, err = db.Exec(query)
	return err
}
