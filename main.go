package main

import (
	"github.com/ivandersr/go-todo/src/config"

	"github.com/ivandersr/go-todo/src/routes"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	routes.Routes()
}
