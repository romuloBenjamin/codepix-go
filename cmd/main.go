package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/romuloBenjamin/codepix-go/infraestructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
}
