package models

import (
	"database/sql"
	"fmt"
	"github/losunioncode/library-managment-system/internal/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Overdue  int64  `json:"overdue"`
	Type     int64  `json:"type"`
}

func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func CreateNewUser(user User) error {
	_, err := database.DB.Exec(
		"INSERT INTO Userlist (id, name, password, overdue, type) "+
			"VALUES (?, ?, ?, ?, ?)",
		user.ID, user.Username, user.Password, 0, 1)

	if err != nil {
		return fmt.Errorf("Error inserting new userlist: %v", err)
	}

	fmt.Println("User was created successfully")

	return nil
}

func CheckUserExist(username string) (User, error) {
	var user User
	err := database.DB.QueryRow("SELECT * FROM Userlist WHERE Name= ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.Overdue, &user.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}

	return user, nil

}