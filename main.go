package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	migrate()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome")
	})
	router.Static("/images", "./images")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	router.Run(":" + port)
}

func migrate() (err error){
	db, err := connectDb()
	if err != nil {
		fmt.Println("error on connectDb")
		fmt.Println(err)
		return err
	}

	db.AutoMigrate(&Monster{})
	return nil
}

func connectDb() (db *gorm.DB, err error) {
	// Heroku には DATABASE_URL というのが設定されている
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		// for local env
		// local では sslmode を disable にしないとつながらない
		url = "postgres://postgres:postgres@127.0.0.1:5432/telesan?sslmode=disable"
	}

	return gorm.Open("postgres", url)
}

type Monster struct {
	Name string
	Hp int
	Attack int
	Defence int
}
