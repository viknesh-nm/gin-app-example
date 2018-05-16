package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	router *gin.Engine
	err    error
	db     *sql.DB
)

func main() {
	// Please fill up the details
	var (
		driverName = ""
		userName   = ""
		password   = ""
		sqlPort    = ""
		dbName     = ""
	)

	db, err = sql.Open(driverName, fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s", userName, password, sqlPort, dbName))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	router.StaticFS("/css", gin.Dir("css", true))
	router.StaticFS("/img", gin.Dir("img", true))
	router.StaticFS("/js", gin.Dir("js", true))
	router.StaticFS("/fonts", gin.Dir("fonts", true))

	router.StaticFS("/admin/css", gin.Dir("css", true))
	router.StaticFS("/admin/js", gin.Dir("js", true))
	router.StaticFS("/admin/img", gin.Dir("img", true))
	router.StaticFS("/admin/fonts", gin.Dir("fonts", true))

	router.StaticFS("/admin/tags/edit/css", gin.Dir("css", true))
	router.StaticFS("/admin/tags/edit/js", gin.Dir("js", true))
	router.StaticFS("/admin/tags/edit/img", gin.Dir("img", true))
	router.StaticFS("/admin/tags/edit/fonts", gin.Dir("fonts", true))

	router.StaticFS("/admin/blogs/edit/css", gin.Dir("css", true))
	router.StaticFS("/admin/blogs/edit/js", gin.Dir("js", true))
	router.StaticFS("/admin/blogs/edit/img", gin.Dir("img", true))
	router.StaticFS("/admin/blogs/edit/fonts", gin.Dir("fonts", true))

	router.LoadHTMLGlob("templates/*/*")

	// Initialize the routes
	initRoutes()

	// Start serving the application
	router.Run(":9000")
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
