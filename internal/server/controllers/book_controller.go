package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/losunioncode/library-managment-system/internal/models"
	"net/http"
)

func HandleRegisterNewBook(c *gin.Context) {
	c.HTML(http.StatusOK, "booklist/booklist-p-register.html", nil)
}

func HandleSearchTitlePage(c *gin.Context) {
	c.HTML(http.StatusOK, "booklist/booklist-p-title.html", nil)
}

func HandleSearchPage(c *gin.Context) {
	c.HTML(http.StatusOK, "booklist/booklist-p.html", nil)
}

func HandleSearchISBNPage(c *gin.Context) {
	c.HTML(http.StatusOK, "booklist/booklist-p-isbn.html", nil)
}

func SearchISBNBook(c *gin.Context) {
	var ISBN string = c.PostForm("ISBN")

	books, err := models.SearchByISBN(ISBN)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(books)
	c.IndentedJSON(http.StatusOK, gin.H{"Books by ISBN": books})
}
func SearchBookTitle(c *gin.Context) {
	var title string = c.PostForm("title")

	books, err := models.SearchBookByTitle(title)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(books)
	c.IndentedJSON(http.StatusOK, gin.H{"Books by title": books})
}

func SearchBookAuthor(c *gin.Context) {
	var author string = c.PostForm("author")

	books, err := models.SearchBookByAuthor(author)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(books)
	c.IndentedJSON(http.StatusOK, gin.H{"Books by author": books})

}
