package controllers

import (
	"github.com/Pcorner0/corpmin-api/models"
	"github.com/Pcorner0/corpmin-api/database"
	"github.com/Pcorner0/corpmin-api/utils/errors"
	"github.com/Pcorner0/corpmin-api/utils/token"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RepoUsers struct {
	DB *gorm.DB
}

func NewRepoUsers() *RepoUsers {
	db := database.InitDB()
	db.AutoMigrate(&models.Users{})
	return &RepoUsers{
		DB: db,
	}
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegistrarUser
func (u *RepoUsers) RegistrarUser(c *gin.Context) {
	var input models.Users

	if err := c.ShouldBindJSON(&input); err != nil {
		err := errors.NewBadRequestError("Invalid json body")
		c.JSON(err.Status, err)
		return
	}

	user := models.Users{}

	user.Email = input.Email
	user.Password = input.Password

	_, err := input.SaveUser(u.DB)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}

func (u *RepoUsers) LogIn(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.Users{}

	user.Email = input.Email
	user.Password = input.Password

	token, err, Usuario := models.LoginCheck(u.DB, user.Email, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"name":     Usuario.Nombre,
		"apellido": Usuario.Apellidop,
		"rol":      Usuario.Rol,
		"office":   Usuario.Office,
	})
}

func (u *RepoUsers) CurrenteUser(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserByID(u.DB, userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

func (u *RepoUsers) LogOut(c *gin.Context) {
	c.GetHeader("jwt")
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
