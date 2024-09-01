package controllers

import (
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github/losunioncode/library-managment-system/internal/models"
	"github/losunioncode/library-managment-system/internal/utils"

	"net/http"
)

func HandlerLoginUserPage(c *gin.Context) {
	c.HTML(http.StatusOK, "userlist/userlist-login.html", nil)

}

func HandlePasswordToChangePage(c *gin.Context) {
	c.HTML(http.StatusOK, "userlist/userlist-password.html", nil)
}

func HandlerCreateUserPage(c *gin.Context) {
	c.HTML(http.StatusOK, "userlist/userlist-create.html", nil)
}

func PasswordToChangeHandler(c *gin.Context) {
	//var user models.User

	tokenId, err := c.Cookie("Authorization")
	passwordToChange := c.PostForm("password")

	if len(passwordToChange) < 4 || passwordToChange == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error has occured while changing password": "Password has not passed the requiremnts."})
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error occurred while trying to get user token Id ": err.Error()})
	}

	err, userId := utils.ValidateToken(tokenId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error occurred while trying to validate token ID from user ": err.Error()})
	}

	user, err := models.GetCurrentUser(userId)
	err = user.HashPassword(passwordToChange)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Could not hash the password, please try again later."})
	}

	err = user.PasswordChangeUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Could not update the password, please try again later."})
	}
	c.JSON(http.StatusOK, gin.H{"result": user})

	//passwordToChange := c.PostForm("password")

}

func HandleLogoutUser(c *gin.Context) {
	c.SetCookie("Authorization", "", -1, "/", "", false, false)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "user was logged out",
	})
}
func RegisterUser(c *gin.Context) {
	var user models.User

	user.ID, _ = gonanoid.Generate("abcde", 11)
	usernameToCreate := c.PostForm("username")
	passwordToCreate := c.PostForm("password")

	if usernameToCreate == "" || len(usernameToCreate) < 4 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Username missed or didn't fulfill requirements"})
		c.Abort()
		return
	}
	if passwordToCreate == "" || len(passwordToCreate) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Password missed or didn't fulfill requirements"})
		c.Abort()
		return
	}
	user.Username = usernameToCreate
	err := user.HashPassword(passwordToCreate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	err = models.CreateNewUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Could not create User ": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"User was created": user.ID})
}
