package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/raylicola/NFlaquiz/database"
	"github.com/raylicola/NFlaquiz/models"
)

// ユーザー情報の取得
func Auth(c *gin.Context){
	cookie, err := c.Cookie("jwt")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err_msg": err.Error()})
		return
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET_KEY"), nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err_msg": err.Error()})
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	id := claims["sub"]

	var user models.User
	database.DB.Where("id = ?", id).First(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}