package main

import "github.com/gin-gonic/gin"

// initRoutes inits all the router to handle
func initRoutes() {
	router.Use(setUserStatus())

	router.GET("/contact", showContactForm)
	router.POST("/contact", contactPost)
	router.GET("/admin", ensureLoggedIn(), func(c *gin.Context) {
		c.Redirect(307, "/admin/job_openings")
	})
	router.GET("/test", func(c *gin.Context) {
		c.HTML(200, "test.html", nil)
	})

	// Admin Handler
	adminRoutes := router.Group("/admin")
	{
		// Login-Logut
		adminRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)
		adminRoutes.GET("/logout", ensureLoggedIn(), logout)

		// JOB-Details
		adminRoutes.POST("/job_openings", ensureNotLoggedIn(), performLogin)
		adminRoutes.GET("/job_openings", ensureLoggedIn(), showIndexPage)

		adminRoutes.GET("/add_new_job", ensureLoggedIn(), showNewJobPage)
		adminRoutes.POST("/add_new_job", ensureLoggedIn(), addNewJob)
		adminRoutes.GET("/edit", ensureLoggedIn(), showEditPage)
		adminRoutes.POST("/edit", ensureLoggedIn(), editPage)
		adminRoutes.GET("/delete/:id", ensureLoggedIn(), deleteJobList)

		// Blog-Details
		adminRoutes.GET("/blogs", ensureLoggedIn(), showBlogs)
		adminRoutes.GET("/add_blog", ensureLoggedIn(), showAddBlogPage)
		adminRoutes.POST("/add_blog", ensureLoggedIn(), AddBlogPage)
		adminRoutes.GET("/editBlog", ensureLoggedIn(), showEditBlogPage)
		adminRoutes.POST("/editBlog", ensureLoggedIn(), editBlog)
		adminRoutes.GET("/blogs/delete/:id", ensureLoggedIn(), deleteBlog)

		// Category
		adminRoutes.GET("/categories", ensureLoggedIn(), showCategories)
		adminRoutes.POST("/categories", ensureLoggedIn(), addCategory)
		adminRoutes.POST("/categorieEdit/:id", ensureLoggedIn(), editCategory)
		adminRoutes.GET("/categories/delete/:id", ensureLoggedIn(), deleteCategory)

		// Tag
		adminRoutes.GET("/tags", ensureLoggedIn(), showTags)
		adminRoutes.POST("/tags", ensureLoggedIn(), addTag)
		adminRoutes.POST("/tags/edit/:id", ensureLoggedIn(), editTag)
		adminRoutes.GET("/tags/delete/:id", ensureLoggedIn(), deleteTag)
	}
}
