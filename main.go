package main

import (
	"sca_api/config"
)

func main() {
	db := config.ConnectDB()
	db.Ping()
}