package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	migrate()
	router := gin.Default()
	router.GET("/monsters", getMonsters)
	router.POST("/monster", postMonster)
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

type getMonstersQuery struct {
	Names string `form:"names" binding:"required"`
}

func getMonsters(c *gin.Context) {
	var query getMonstersQuery

	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{})
		return
	}

	fmt.Println(query.Names)
	names := strings.Split(query.Names, ",")
	fmt.Println(names)

	db, _ := connectDb()
	defer db.Close()

	var monsters []Monster
	db.Where("name IN (?)", names).Find(&monsters)
	c.JSON(200, monsters)
}

func postMonster(c *gin.Context) {
	var query Monster

	if err := c.BindJSON(&query); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{})
		return
	}

	db, _ := connectDb()
	defer db.Close()
	db.Where(Monster{Name: query.Name}).Assign(Monster{
		Health:  query.Health,
		Attack:  query.Attack,
		Defence: query.Defence,
	}).FirstOrCreate(&Monster{})
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
	Name    string `gorm:"primary_key" json:"name" binding:"required"`
	Health  int    `json:"health" binding:"required"`
	Attack  int    `json:"attack" binding:"required"`
	Defence int    `json:"defence" binding:"required"`
}
