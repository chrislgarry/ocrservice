package main

import (
	"ocrservice/handler"
	"ocrservice/models"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Env struct {
	ocr models.OcrModel
}

func main() {

	//Initialize database
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database " + dsn)
	}
	db.AutoMigrate(&models.OcrAsync{})

	// env := &Env{
	// 	ocr: models.OcrModel{DB: db},
	// }

	// Initialize Echo router
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.POST("/image-sync", handler.SyncOcr)
	e.POST("/image", handler.AsyncOcr)
	e.GET("/image", handler.AsyncOcr)
	e.Logger.Fatal(e.Start(":" + os.Getenv("API_PORT")))
}
