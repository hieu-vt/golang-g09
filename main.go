package main

import (
	"g09-to-do-list/middleware"
	gin3 "g09-to-do-list/module/item/transport/gin"
	"g09-to-do-list/module/upload"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	r.Use(middleware.Recover())

	r.Static("/static", "./static")

	v1 := r.Group("/v1")
	{
		v1.PUT("/upload", upload.Upload(db))

		items := v1.Group("/items")
		{
			items.POST("", gin3.CreateNewItem(db))
			items.GET("", gin3.ListItem(db))
			items.GET("/:id", gin3.GetItem(db))
			items.PATCH("/:id", gin3.UpdateItemHandler(db))
			items.DELETE("/:id", gin3.DeleteItem(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(":3000"); err != nil {
		log.Fatalln(err)
	}
}
