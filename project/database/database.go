package database

import (
	"fmt"

	"github.com/Manuel11713/go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:Spartan11713@/go_auth?parseTime=true"), &gorm.Config{})
	DB = db
	if err != nil {
		panic("could not connect to the db")
	}
	DB.AutoMigrate(&models.User{}, &models.PasswordReset{})
	fmt.Println("DB is working", DB)
}
