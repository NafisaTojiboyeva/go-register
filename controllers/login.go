package controllers

import (
	"log"
	"net/http"

	"github.com/NafisaTojiboyeva/go-register/config"
	"github.com/NafisaTojiboyeva/go-register/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("mysession"))

func Login(ctx *gin.Context) {

	// ctx.Request.ParseForm()

	// phone := ctx.Request.Form.Get("phone")
	// password := ctx.Request.Form.Get("password")

	var loginData models.LoginInfo

	if err := ctx.BindJSON(&loginData); err != nil {
		log.Print(err)
		return
	}

	db, err := config.DB()

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Can not connected to database"})
	}

	var user models.User

	result := db.Where("phone = ? and password = ?", loginData.Phone, loginData.Password).Find(&user)

	if result.RowsAffected == 1 {

		session, _ := store.Get(ctx.Request, "mysession")

		session.Values["user_id"] = user.Id

		session.Values["fullname"] = user.Fullname

		session.Save(ctx.Request, ctx.Writer)

		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully loged in! Welcome " + user.Fullname})
	} else {

		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invailed confirmation code!"})
	}
}
