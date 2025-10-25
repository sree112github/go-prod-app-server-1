package config

import (
	"basics/internal/utils"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {

	host := utils.Cfg.DatabaseHostName
	dbPort := utils.Cfg.DatabasePortNumber
	dbUser := utils.Cfg.DatabaseUserName
	dbPassword := utils.Cfg.DatabasePassword
	dbName := utils.Cfg.DatabaseName
	dbSsl := utils.Cfg.DBSSL

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, dbPort, dbUser, dbPassword, dbName, dbSsl)

	var err error

	DB, err = sql.Open("postgres", dns)

	if err != nil {
		log.Fatal("Database Initialization Failed Due to: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database unreachable")
	}

	fmt.Println("✅Database Connection Succeed...")

}
