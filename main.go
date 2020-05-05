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
	user := os.Getenv("TELESAN_DB_USER")
	if user == "" {
		user = "postgres"
	}

	password := os.Getenv("TELESAN_DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}

	host := os.Getenv("TELESAN_DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	database := os.Getenv("TELESAN_DB_DATABASE")
	if database == "" {
		database = "telesan"
	}
	return gorm.Open("postgres", "postgres://" + user+ ":" + password + "@" + host + ":5432" + "/" + database + "?sslmode=disable")
}

type Monster struct {
	Name string
	Hp int
	Attack int
	Defence int
}
