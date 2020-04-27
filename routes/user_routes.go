package routes

import (
	"github.com/tsurusekazuki/sampleapp/config"

	"net/http"
	"github.com/gin-gonic/gin"
)

func UserSignUp(ctx *gin.Context) {
	println("post/signup")
	username := ctx.PostForm("username")
	emailAddress := ctx.PostForm("emailAddress")
	password := ctx.PostForm("password")
	passwordConf := ctx.PostForm("passwordConfirmation")

	if password != passwordConf {
		println("Error; password and passwordConf not match")
		ctx.Redirect(http.StatusSeeOther, "//localhost:8000/")
		return
	}

	db := config.DummyDB()
	if err := db.SaveUser(username, emailAddress, password); err != nil {
		println("Error " + err.Error())
	} else {
		println("Signup success")
		println("  username: " + username)
		println("  email: " + emailAddress)
		println("  password: " + password)
	}
	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
}

func UserLogIn(ctx *gin.Context) {
	println("post/login")
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	db := config.DummyDB()
	user, err := db.GetUser(username, password)

	if err != nil {
		println("Error: " + err.Error())
	} else {
		println("Authentication Success!!")
		println("  username: " + user.Username)
		println("  email: " + user.Email)
		println("  password: " + user.Password)
		user.Authenticate()
	}

	ctx.Redirect(http.StatusSeeOther, "//localhost:8080/")
}
