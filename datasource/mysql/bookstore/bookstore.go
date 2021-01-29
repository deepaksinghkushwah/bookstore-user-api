package bookstore

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsername = "mysql_username"
	mysqlPassword = "mysql_password"
	mysqlHost     = "mysql_host"
	mysqlDB       = "mysql_db"
)

var (
	// BookStoreDBLink dbclient
	BookStoreDBLink *sql.DB
)

func init() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalln(".env file not loaded")
	}

	username := os.Getenv(mysqlUsername)
	password := os.Getenv(mysqlPassword)
	host := os.Getenv(mysqlHost)
	db := os.Getenv(mysqlDB)

	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, db)
	BookStoreDBLink, err = sql.Open("mysql", dataSource)

	//defer BookStoreDBLink.Close()

	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Database successfully configured")
}
