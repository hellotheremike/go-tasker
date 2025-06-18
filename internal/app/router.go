package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hellotheremike/go-tasker/internal/auth"
	lib "github.com/hellotheremike/go-tasker/internal/db"
	"github.com/hellotheremike/go-tasker/internal/middleware"
	"github.com/hellotheremike/go-tasker/internal/users"
)


func SetupRouter() *gin.Engine {
	router := gin.Default()
	
	db, _ := lib.Connect()

	repo := users.NewRepository(db);
	service := users.NewService(repo);
	handler := users.NewHandler(service);


	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	
	protected := router.Group("/p", middleware.JwtMiddleware()) 
	{
		protected.GET("/users", handler.GetAll)
	}
	router.POST("/jwt", auth.GenerateToken)

	return router;
}