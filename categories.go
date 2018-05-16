package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Category holds the category fields
type Category struct {
	ID           int
	CategoryName string
}

var categoryList = []Category{}

// showCategories shows the fetched values of categories in html
func showCategories(c *gin.Context) {
	render(c, gin.H{
		"Title":   "Categories",
		"payload": listAllCatergories(),
	}, "categories.html")
}

// addCategory posts the category details
// Post Method
func addCategory(c *gin.Context) {
	if _, err := insertCategory(c.PostForm("categoryName")); err == nil {
		c.Redirect(301, "/admin/categories")
	} else {
		c.AbortWithStatus(400)
	}
}

// insertCategory fetches the values from the entered form and stores in a DB
// Returns the Category detailed values
func insertCategory(categoryName string) (*Category, error) {
	_, err := db.Exec("INSERT INTO blogCategories(name, createdAt, updatedAt) VALUES(?,?,?)", categoryName, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}
	return &Category{ID: len(categoryList) + 1, CategoryName: categoryName}, nil
}

// listAllCatergories lists all the categories from the DB
func listAllCatergories() []Category {
	var (
		lists = []Category{}
		list  Category
		id    int
		name  string
	)
	rows, err := db.Query("select id, name from blogCategories")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}

		list = Category{ID: id, CategoryName: name}
		lists = append(lists, list)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return lists
}

// editCategory performs the updation of the edited fields
// POST Method
func editCategory(c *gin.Context) {
	if openingsID, err := strconv.Atoi(c.Param("id")); err == nil {
		if _, err := updateCategory(openingsID, c.PostForm("categoryName")); err == nil {
			c.Redirect(301, "/admin/categories")
		} else {
			c.AbortWithStatus(400)
		}
	} else {
		c.AbortWithStatus(404)
	}
}

// updateCategory updates the given values to the DB and returns it back
func updateCategory(openingsID int, name string) (*Category, error) {
	_, err := db.Exec("UPDATE blogCategories SET name = ?, updatedAt = ? WHERE id = ?", name, time.Now(), openingsID)
	if err != nil {
		return nil, err
	}
	return &Category{ID: openingsID, CategoryName: name}, nil
}

// deleteCategory deletes the category details with the particular selected id
func deleteCategory(c *gin.Context) {
	if categoryID, err := strconv.Atoi(c.Param("id")); err == nil {
		_, err := db.Exec("DELETE FROM blogCategories where id= ?", categoryID)
		if err == nil {
			c.Redirect(307, "/admin/categories")
		}
	} else {
		c.AbortWithStatus(404)
	}
}
