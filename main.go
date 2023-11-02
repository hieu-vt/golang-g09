package main

import (
	"g09-to-do-list/component/tokenprovider/jwt"
	"g09-to-do-list/middleware"
	gin3 "g09-to-do-list/module/item/transport/gin"
	"g09-to-do-list/module/upload"
	"g09-to-do-list/module/user/storage"
	gin2 "g09-to-do-list/module/user/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DB_CONN")
	secret := os.Getenv("SECRET_KEY")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	tokenProvider := jwt.NewTokenJWTProvider("jwt", secret)

	authStore := storage.NewSQLStore(db)

	r := gin.Default()

	r.Use(middleware.Recover())

	r.Static("/static", "./static")

	v1 := r.Group("/v1")
	{
		v1.PUT("/upload", upload.Upload(db))

		items := v1.Group("/items")
		{
			items.POST("", middleware.RequiredAuth(authStore, tokenProvider), gin3.CreateNewItem(db))
			items.GET("", gin3.ListItem(db))
			items.GET("/:id", gin3.GetItem(db))
			items.PATCH("/:id", middleware.RequiredAuth(authStore, tokenProvider), gin3.UpdateItemHandler(db))
			items.DELETE("/:id", middleware.RequiredAuth(authStore, tokenProvider), gin3.DeleteItem(db))
		}

		auth := v1.Group("/auth")
		{
			auth.POST("/signup", gin2.CreateUserHandler(db))
			auth.POST("/login", gin2.LoginHandler(db, tokenProvider))
		}

		v1.GET("/profile", middleware.RequiredAuth(authStore, tokenProvider), gin2.Profile())
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
