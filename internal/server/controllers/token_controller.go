package controllers

import (
	"github.com/gin-gonic/gin"
	"github/losunioncode/library-managment-system/internal/models"
	"github/losunioncode/library-managment-system/internal/utils"
	"net/http"
)

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateToken(c *gin.Context) {
	var request TokenRequest
	var user models.User

	var usernameToCheck = c.PostForm("username")
	var passwordToCheck = c.PostForm("password")

	if usernameToCheck == "" || len(usernameToCheck) < 4 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Username missed or didn't fulfill requirements"})
		c.Abort()
		return
	}
	if passwordToCheck == "" || len(passwordToCheck) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Password missed or didn't fulfill requirements"})
		c.Abort()
		return
	}

	request.Username = usernameToCheck
	request.Password = passwordToCheck
	user, err := models.CheckUserExist(request.Username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	credentialsError := user.CheckPassword(request.Password)
	if credentialsError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": credentialsError.Error()})
		c.Abort()
		return
	}

	tokenString, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.SetCookie("Authorization", tokenString, 60*60*24, "/", "", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{"token": tokenString})
}

func HandleHeader(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{"curr_header": "header"})
}
