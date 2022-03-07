package database

import (
	"fmt"
	"go-fiber/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	p := config.Config("DB_PORT")
	port, _ := strconv.Atoi(p)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"),
		config.Config("DB_PASS"),
		config.Config("DB_HOST"),
		port,
		config.Config("DB_NAME"))
	//dsn := "root:MXXyrg56614@tcp(node9848-env-6465275.user.cloudjkt02.com:3306)/fiber_app?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("failed connect to database")
	}
	fmt.Println("connection opened to database")
}
