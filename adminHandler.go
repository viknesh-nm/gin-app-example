package main

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var userName string

// showLoginPage shows the login page
func showLoginPage(c *gin.Context) {
	render(c, gin.H{
		"Title": "Login",
	}, "login.html")
}

// performLogin performs the login modules for the user
func performLogin(c *gin.Context) {
	var databasePassword string

	err := db.QueryRow("SELECT username, password FROM tableUsers WHERE username=?", c.PostForm("username")).Scan(&userName, &databasePassword)

	if err != nil {
		c.Redirect(301, "/admin/login")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(c.PostForm("password")))
	if err != nil {
		c.Redirect(301, "/admin/login")
		return
	}
	uID, _ := uuid.NewV4()
	c.SetCookie("session", uID.String(), 3600, "", "", false, true)
	c.Set("is_logged_in", true)

	render(c, gin.H{
		"Title": "DashBoard", "UserName": userName, "payload": listAllJobs()}, "index.html",
	)
}

// showIndexPage shows the Dashboard field after the login performance
func showIndexPage(c *gin.Context) {
	render(c, gin.H{
		"Title": "Dashboard", "UserName": userName,
		"payload": listAllJobs(),
	}, "index.html")
}

// logout clears the session values and redirects to login
func logout(c *gin.Context) {
	c.SetCookie("session", "", -1, "", "", false, true)
	c.Redirect(307, "/admin/login")
}
