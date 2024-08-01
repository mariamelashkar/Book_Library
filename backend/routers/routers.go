package routers

import (
	"bookstore/handlers"
	//"bookstore/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine){
	router.POST("/register", handlers.RegisterUser)
	router.POST("/login", handlers.LoginUser)

	authorized := router.Group("/")
	//authorized.Use(middlewares.authMiddleware())
	{
		authorized.POST("/uploadbook",handlers.UploadBook)
		authorized.PATCH("/borrowbook/:Name", handlers.BorrowBook)
	}

	router.GET("/getbook", handlers.Getbook)
	router.GET("/getbook/:Name", handlers.Getbookbyname)
	router.POST("/createbook", handlers.CreateBook)

}