package controllers

import (
	"net/http"
	// "html/template"
	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {

	session, _ := store.Get(ctx.Request, "mysession")
	session.Options.MaxAge = -1
	session.Save(ctx.Request, ctx.Writer)
	
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Successfully loged out!"})
}