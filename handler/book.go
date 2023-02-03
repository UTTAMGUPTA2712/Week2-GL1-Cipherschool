package handler
import (
	"log"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"net/http"
	
	"github.com/uttamgupta2712/Week2-GL1-Cipherschool/database"
	
	"github.com/uttamgupta2712/Week2-GL1-Cipherschool/model"

)

type Handler struct{}
func GetBooks(in *gin.Context) {
	DB *gorm.DB
}

func (h *Handler) GetBooks(in *gin.Context) {
	books, err := database.GetBooks(h.DB) // h.DB is fully configured and can give access of book table
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, books)

}
func (h *Handler) SaveBook(in *gin.Context) {
	book := models.Book{}
	err := in.BindJSOn(&book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}

	err = database.SaveBook(h.DB, &book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, book)
}
