package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Blog holds the blog fields
type Blog struct {
	ID             int
	Title          string
	Category       string
	BlogCategories []string
	Tag            string
	Tags           []Tag
	Description    string
}

// BlogCategory holds the DB category and name of the category for checking
type BlogCategory struct {
	Name           string
	BlogCategories []Category
}

var blogList []Blog

// showBlogs lists the blogs from the DB
func showBlogs(c *gin.Context) {
	render(c, gin.H{
		"Title": "Blogs",
		"Blogs": listAllBlogs(),
	}, "blogs.html")
}

// showAddBlogpage execute the add blog page (GET)
func showAddBlogPage(c *gin.Context) {
	render(c, gin.H{
		"Title":      "Editor",
		"UserName":   userName,
		"Categories": listAllCatergories(),
		"Tags":       listAllTags(),
	}, "add_blog.html")
}

// AddBlogPage posts the blog details
// Post Method
func AddBlogPage(c *gin.Context) {
	if _, err := insertBlog(c.PostForm("blog_title"), c.PostForm("ckeditor"), c.PostFormArray("category"), c.PostFormArray("tag")); err == nil {
		c.Redirect(301, "/admin/blogs")
	} else {
		c.AbortWithStatus(400)
	}
}

// insertBlog fetches the values from the entered form and stores in a DB
// Returns the Blog detailed values
func insertBlog(title, description string, categories, tags []string) (*Blog, error) {
	_, err := db.Exec("INSERT INTO blogs(title, description, categories, tags, createdAt, updatedAt) VALUES(?,?,?,?,?,?)", title, description, strings.Join(categories, ","), strings.Join(tags, ","), time.Now(), time.Now())
	if err != nil {
		return nil, err
	}
	return &Blog{ID: len(blogList) + 1, Title: title, Description: description, Category: strings.Join(categories, ","), Tag: strings.Join(tags, ",")}, nil
}

// listAllBlogs lists all the blogs from the DB
func listAllBlogs() []Blog {
	var (
		blogLists          = []Blog{}
		blogList           Blog
		id                 int
		title, description string
		categories, tags   string
	)
	rows, err := db.Query("select id, title, description, categories, tags from blogs")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &title, &description, &categories, &tags)
		if err != nil {
			log.Fatal(err)
		}
		blogList = Blog{ID: id, Title: title, Description: description, Category: categories, Tag: tags}
		blogLists = append(blogLists, blogList)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return blogLists
}

// deleteBlog deletes the blog details with the particular selected id
func deleteBlog(c *gin.Context) {
	if blogID, err := strconv.Atoi(c.Param("id")); err == nil {
		_, err := db.Exec("DELETE FROM blogs where id= ?", blogID)
		if err == nil {
			c.Redirect(307, "/admin/blogs")
		}
	} else {
		c.AbortWithStatus(404)
	}
}

// showEditBlogPage shows the edit blog page fetched from the particular ID
// GET method
func showEditBlogPage(c *gin.Context) {
	if blogID, err := strconv.Atoi(c.Query("blogID")); err == nil {
		if blogdet, err := getBlogByID(blogID); err == nil {
			render(c, gin.H{
				"title":      blogdet.Title,
				"Blog":       blogdet,
				"Categories": listAllCatergories(),
				"Tags":       listAllTags(),
			}, "edit_blog.html")
		} else {
			c.AbortWithError(404, err)
		}
	} else {
		c.AbortWithStatus(404)
	}
}

// getBlogByID returns the blog details with respective ID's
func getBlogByID(id int) (*Blog, error) {
	var (
		title, description string
		categories, tags   string
		// category           = BlogCategory{}
	)
	rows, err := db.Query("select id, title, description, categories, tags from blogs where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &title, &description, &categories, &tags)
		if err != nil {
			log.Fatal(err)
		}

		categories := strings.Split(categories, ",")

		return &Blog{ID: id, Title: title, Description: description, BlogCategories: categories}, nil
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return nil, errors.New("Openings not found")
}

// editBlog performs the updation of the edited fields
// POST Method
func editBlog(c *gin.Context) {
	if blogID, err := strconv.Atoi(c.Query("blogID")); err == nil {
		if _, err := updateBlog(blogID, c.PostForm("blog_title"), c.PostForm("ckeditor"), c.PostFormArray("category"), c.PostFormArray("tag")); err == nil {
			c.Redirect(302, "/admin/blogs")
		} else {
			c.AbortWithStatus(400)
		}
	} else {
		c.AbortWithStatus(404)
	}
}

// updateBlog updates the given values to the DB and returns it back
func updateBlog(blogID int, title, description string, categories, tags []string) (*Blog, error) {
	_, err := db.Exec("UPDATE blogs SET title = ?, description = ?, categories = ?, tags = ?, updatedAt = ? WHERE id = ?",
		title, description, strings.Join(categories, ","), strings.Join(tags, ","), time.Now(), blogID)
	if err != nil {
		return nil, err
	}
	return &Blog{ID: blogID, Title: title, Description: description, Category: strings.Join(categories, ","), Tag: strings.Join(tags, ",")}, nil
}
