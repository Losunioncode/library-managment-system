package models

import (
	"database/sql"
	"fmt"
	"github/losunioncode/library-managment-system/internal/database"
)

type BookSample struct {
	Title      string         `json:"title"`
	ISBN       string         `json:"isbn"`
	Author     string         `json:"author"`
	Publisher  string         `json:"publisher"`
	Stock      int64          `json:"stock"`
	Available  int64          `json:"avaible"`
	RemoveInfo sql.NullString `json:"remove_info"`
}

func SearchBookByTitle(title string) ([]BookSample, error) {
	var books []BookSample

	rows, err := database.DB.Query("SELECT * FROM Booklist WHERE Title= ?", title)

	if err != nil {
		return books, fmt.Errorf("handlers-book search by author: %v", err)
	}

	for rows.Next() {
		var book BookSample

		err := rows.Scan(&book.Title, &book.ISBN, &book.Author, &book.Publisher, &book.Stock, &book.Available, &book.RemoveInfo)
		if err != nil {
			return nil, fmt.Errorf("scan error occured : %v", err)
		}
		books = append(books, book)

	}

	err = rows.Err()

	if err != nil {
		return nil, fmt.Errorf("rows error occured : %v", err)
	}

	return books, nil
}

func SearchBookByAuthor(name string) ([]BookSample, error) {
	var books []BookSample

	rows, err := database.DB.Query("SELECT * FROM Booklist WHERE Author= ?", name)

	if err != nil {
		return books, fmt.Errorf("handlers-book search by author: %v", err)
	}

	for rows.Next() {
		var book BookSample

		err := rows.Scan(&book.Title, &book.ISBN, &book.Author, &book.Publisher, &book.Stock, &book.Available, &book.RemoveInfo)
		if err != nil {
			return nil, fmt.Errorf("scan error occured : %v", err)
		}
		books = append(books, book)

	}

	err = rows.Err()

	if err != nil {
		return nil, fmt.Errorf("rows error occured : %v", err)
	}

	return books, nil
}

func SearchByISBN(ISBN string) ([]BookSample, error) {
	var books []BookSample

	rows, err := database.DB.Query("SELECT * FROM Booklist WHERE ISBN= ?", ISBN)

	if err != nil {
		return books, fmt.Errorf("handlers-book search by author: %v", err)
	}

	for rows.Next() {
		var book BookSample

		err := rows.Scan(&book.Title, &book.ISBN, &book.Author, &book.Publisher, &book.Stock, &book.Available, &book.RemoveInfo)
		if err != nil {
			return nil, fmt.Errorf("scan error occured : %v", err)
		}
		books = append(books, book)

	}

	err = rows.Err()

	if err != nil {
		return nil, fmt.Errorf("rows error occured : %v", err)
	}

	return books, nil
}
