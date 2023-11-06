package cmd

import (
	"fmt"
	"g09-to-do-list/common"
	"g09-to-do-list/component/tokenprovider/jwt"
	"g09-to-do-list/middleware"
	gin3 "g09-to-do-list/module/item/transport/gin"
	"g09-to-do-list/module/upload"
	userstorage "g09-to-do-list/module/user/storage"
	gin2 "g09-to-do-list/module/user/transport/gin"
	"g09-to-do-list/plugin/sdkgorm"
	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func newService() goservice.Service {
	service := goservice.New(
		goservice.WithName("social-todo-list"),
		goservice.WithVersion("1.0.0"),
		goservice.WithInitRunnable(sdkgorm.NewGormDB("main", common.PluginDBMain)),
	)

	return service
}

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Start social TODO service",
	Run: func(cmd *cobra.Command, args []string) {
		systemSecret := os.Getenv("SECRET")

		service := newService()

		serviceLogger := service.Logger("service")

		if err := service.Init(); err != nil {
			serviceLogger.Fatalln(err)
		}

		service.HTTPServer().AddHandler(func(engine *gin.Engine) {
			engine.Use(middleware.Recover())

			db := service.MustGet(common.PluginDBMain).(*gorm.DB)

			authStore := userstorage.NewSQLStore(db)
			tokenProvider := jwt.NewTokenJWTProvider("jwt", systemSecret)
			middlewareAuth := middleware.RequiredAuth(authStore, tokenProvider)

			v1 := engine.Group("/v1")
			{
				v1.PUT("/upload", upload.Upload(db))

				items := v1.Group("/items")
				{
					items.POST("", middlewareAuth, gin3.CreateNewItem(db))
					items.GET("", gin3.ListItem(db))
					items.GET("/:id", gin3.GetItem(db))
					items.PATCH("/:id", middlewareAuth, gin3.UpdateItemHandler(db))
					items.DELETE("/:id", middlewareAuth, gin3.DeleteItem(db))
				}

				auth := v1.Group("/auth")
				{
					auth.POST("/signup", gin2.CreateUserHandler(db))
					auth.POST("/login", gin2.LoginHandler(db, tokenProvider))
				}

				v1.GET("/profile", middlewareAuth, gin2.Profile())
			}

			engine.GET("/ping", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
				})
			})
		})

		if err := service.Start(); err != nil {
			serviceLogger.Fatalln(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
