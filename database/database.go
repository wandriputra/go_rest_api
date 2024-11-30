package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/wandriputra/go_rest_api/config"
	"github.com/wandriputra/go_rest_api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func Connect() {
	// Connect to the database

	p := config.Config("DB_PORT")

	port, err := strconv.ParseInt(p, 10, 32)

	if err != nil {
		panic("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal("Failed to connect to database! \n", err)
		os.Exit(2)
	}

	log.Println("Database connected!")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations...")
	db.AutoMigrate(&model.User{})

	DB = DbInstance{Db: db}
}
