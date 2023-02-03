package routers

import (
	"github.com/gin-gonic/gin"
	
	"github.com/uttamgupta2712/Week2-GL1-Cipherschool/database"
	
	"github.com/uttamgupta2712/Week2-GL1-Cipherschool/handler"
)

func Router() *gin.Engine {
	router := gin.Default() // get the dfault engine for further customization
	api := handler.Handler{
		DB: database.GetDB(), //set the handler DB
	}
	router.GET("/books", api.GetBooks) //set thefunction for thiss  https://localhost:8080/books : Get Request
	// while calling the handler function , gin will pass *gin.Context in the handler function
	router.POST("/books", api.SaveBook)

	return router

}
