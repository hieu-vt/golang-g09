package main

import (
	gin2 "g09-to-do-list/module/transport/gin"
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

	log.Println(db)

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", gin2.CreateNewItem(db))
			items.GET("", gin2.ListItem(db))
			items.GET("/:id", gin2.GetItem(db))
			items.PATCH("/:id", gin2.UpdateItemHandler(db))
			items.DELETE("/:id", gin2.DeleteItem(db))
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
