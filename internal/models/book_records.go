package models

import (
	"database/sql"
	"errors"
	"fmt"
	"github/losunioncode/library-managment-system/internal/database"
	"log"
	"time"
)

type BookRecords struct {
	recordId   string       `json:"record_Id"`
	bookId     string       `json:"book_id"`
	userId     string       `json:"user_id"`
	isReturned bool         `json:"is_returned"`
	borrowDate time.Time    `json:"borrow_date"`
	returnDate sql.NullTime `json:"return_date"`
	Deadline   time.Time    `json:"deadline"`

	Extendtimes int64 `json:"extendtimes"`
}

func CheckDeadlineBook(deadline, returnedDate time.Time, userId string) error {
	if deadline.Before(returnedDate) {
		err := ChangeOverdueUser(userId)
		if err != nil {
			return err
		}
	}

	return nil

}

func BorrowBookFromLibrary(bookISBN, userId string, borrowDate time.Time) error {
	var recordId int

	err := database.DB.QueryRow("SELECT record_id FROM RecordList WHERE book_id = ? AND user_id= ? AND IsReturned = 0 ", bookISBN, userId).Scan(&recordId)

	if err != nil && err != sql.ErrNoRows {
		log.Fatal("Check whether borrowed", err)
		return err
	}

	if err == nil {
		log.Fatal("Book record already exists")
		return errors.New("Book record already exists")
	}

	var availableAmount, stock int

	row := database.DB.QueryRow("SELECT Available, Stock FROM Booklist WHERE ISBN = ? AND Stock > 0", bookISBN)

	err = row.Scan(&availableAmount, &stock)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return errors.New("Could not find a book")
		}
	}

	if availableAmount <= 0 {
		return errors.New("Book is not available currently")
	}

	if stock <= 0 {
		return errors.New("Book is not in stock")
	}

	_, err = database.DB.Exec("UPDATE Booklist SET Available = Available - 1 WHERE ISBN = ?", bookISBN)

	if err != nil {
		return errors.New("Could not update the book record")
	}

	deadline := borrowDate.AddDate(0, 1, 0)
	_, err = database.DB.Exec("INSERT INTO Recordlist (book_id, user_id, IsReturned, borrow_date, deadline, extendtimes) VALUES (?, ?, 0, ?, ?, 0)",
		bookISBN, userId, borrowDate, deadline)

	if err != nil {
		log.Fatal(err)
		return errors.New("Could not create new Record List!")
	}
	log.Println("New Record List was created !")

	return nil
}

func BookRecordExtend(ISBN, userId string, currTime time.Time) error {
	var deadlineRecord time.Time
	var IsReturned bool

	err := database.DB.QueryRow("SELECT deadline, IsReturned FROM RecordList WHERE book_id = ? AND user_id = ?", ISBN, userId).Scan(&deadlineRecord, &IsReturned)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Errorf("Book record does not exist")
			return err
		}
		fmt.Errorf("Error occured while querying for book's deadline value!")
		return err
	}

	if IsReturned {
		return errors.New("Book was already returned !")
	}

	newDeadlineRecord := currTime.AddDate(0, 1, 0)

	_, err = database.DB.Exec("UPDATE RecordList SET deadline= ?, extendtimes=extendtimes+1 WHERE book_id=?", newDeadlineRecord, ISBN)

	if err != nil {
		fmt.Errorf("Error occured while updating book's deadline value!")
		return err
	}

	return nil

}

func ReturnBookToLibrary(ISBN, userId string, returnDate time.Time) error {
	var recordReturn bool
	var deadline time.Time
	err := database.DB.QueryRow("SELECT IsReturned, deadline FROM RecordList WHERE book_id = ? AND user_id= ?", ISBN, userId).Scan(&recordReturn, &deadline)

	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Errorf("Could not find the record's book!")
			return err
		}
		fmt.Errorf("Error has occurred during quering record list")
		return err
	}

	if recordReturn != false {
		fmt.Errorf("The Book has already been received")
		return errors.New("The Book has already been received")

	}

	err = CheckDeadlineBook(deadline, returnDate, userId)

	if err != nil {
		return errors.New("Could not update user overdue status")
	}

	_, err = database.DB.Exec("UPDATE RecordList SET IsReturned = 1, return_date = ? WHERE book_id = ?", returnDate, ISBN)

	if err != nil {
		fmt.Errorf("The error while updating the book record! %v", err)
		return err
	}

	return nil
}

func QueryDeadlineBook(ISBN, userId string) (error, time.Time) {
	var deadline time.Time

	err := database.DB.QueryRow("SELECT deadline FROM RecordBook WHERE book_id= ? AND user_id= ?", ISBN, userId).Scan(&deadline)

	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Errorf("Error occured while querying for book's deadline!")
			return err, deadline
		}
		fmt.Errorf("Could not receive deadline for your book!")
		return err, deadline
	}

	return nil, deadline

}
