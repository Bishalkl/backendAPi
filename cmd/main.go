package main

import (
	"log"

	"github.com/bishalkl/backendAPi/cmd/api"
	"github.com/bishalkl/backendAPi/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "127.0.1:3306",
		DBName:               "ecom",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
