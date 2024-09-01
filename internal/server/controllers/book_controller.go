package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/losunioncode/library-managment-system/internal/models"
	"github/losunioncode/library-managment-system/internal/utils"
	"net/http"
	"time"
)

func HandleRegisterNewBook(c *gin.Context) {
	c.HTML(http.StatusOK, "booklist-recordlist/booklist-p-borrow.html", nil)
}

func HandleSearchTitlePage(c *gin.Context) {
	c.HTML(http.StatusOK, "booklist/booklist-p-title.html", nil)
}

func HandleExtendDeadlinePage(c *gin.Context) {
	c.HTML(http.StatusOK, "booklist-recordlist/booklist-p-extend.html", nil)

}

func HandleReturnBookPage(c *gin.Context) {
	c.HTML(http.StatusOK, "booklist-recordlist/booklist-p-return.html", nil)
}

func HandleSearchPage(c *gin.Context) {
	c.HTML(http.StatusOK, "booklist/booklist-p.html", nil)
}

func HandleCheckDeadlinePage(c *gin.Context) {
	c.HTML(http.StatusOK, "booklist-recordlist/booklist-p-deadline.html", nil)
}

func HandleBorrowBookPage(c *gin.Context) {
	c.HTML(http.StatusOK, "booklist-recordlist/booklist-p-borrow.html", nil)
}
func HandleSearchISBNPage(c *gin.Context) {
	c.HTML(http.StatusOK, "booklist/booklist-p-isbn.html", nil)
}

func CheckBookDeadline(c *gin.Context) {
	var ISBN string = c.PostForm("ISBN")
	tokenUserAuth, _ := c.Cookie("Authorization")
	err, userId := utils.ValidateToken(tokenUserAuth)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err, deadline := models.QueryDeadlineBook(ISBN, userId)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"isbn": ISBN, "deadline": deadline})

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

func ExtendBookDeadline(c *gin.Context) {
	var ISBN string = c.PostForm("ISBN")
	tokenUserAuth, _ := c.Cookie("Authorization")
	var current_Time = time.Now()
	err, userId := utils.ValidateToken(tokenUserAuth)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.BookRecordExtend(ISBN, userId, current_Time)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "Book's record has been extended"})

}

func ReturnBookLibrary(c *gin.Context) {

	var ISBN string = c.PostForm("ISBN")
	tokenUserAuth, _ := c.Cookie("Authorization")
	var currentDate = time.Now()
	err, userId := utils.ValidateToken(tokenUserAuth)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.ReturnBookToLibrary(ISBN, userId, currentDate)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "Book was returned!"})
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

func BorrowBookFromBooklist(c *gin.Context) {
	tokenUser, err := c.Cookie("Authorization")
	booksISBN := c.PostForm("ISBN")
	borrowNewDate := time.Now()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error has occured": err.Error()})

	}
	err, userId := utils.ValidateToken(tokenUser)

	err = models.BorrowBookFromLibrary(booksISBN, userId, borrowNewDate)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Error has occurred while borrowing new book": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"userId": userId, "borrowNewDate": borrowNewDate})

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
