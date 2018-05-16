package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Tag holds the tag fields
type Tag struct {
	ID      int
	TagName string
}

var tagList = []Tag{}

// showTags shows the fetched values of tags in html
func showTags(c *gin.Context) {
	render(c, gin.H{
		"Title":   "Tags",
		"payload": listAllTags(),
	}, "tags.html")
}

// addTag posts the tag details
// Post Method
func addTag(c *gin.Context) {
	if _, err := insertTag(c.PostForm("tagName")); err == nil {
		c.Redirect(301, "/admin/tags")
	} else {
		c.AbortWithStatus(400)
	}
}

// insertTag fetches the values from the entered form and stores in a DB
// Returns the Tag detailed values
func insertTag(tagName string) (*Tag, error) {
	_, err := db.Exec("INSERT INTO blogTags(name, createdAt, updatedAt) VALUES(?,?,?)", tagName, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}
	return &Tag{ID: len(tagList) + 1, TagName: tagName}, nil
}

// listAllTags lists all the tags from the DB
func listAllTags() []Tag {
	var (
		lists = []Tag{}
		list  Tag
		id    int
		name  string
	)
	rows, err := db.Query("select id, name from blogTags")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}

		list = Tag{ID: id, TagName: name}
		lists = append(lists, list)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return lists
}

// editTag performs the updation of the edited fields
// POST Method
func editTag(c *gin.Context) {
	if openingsID, err := strconv.Atoi(c.Param("id")); err == nil {
		if _, err := updateTag(openingsID, c.PostForm("tagName")); err == nil {
			c.Redirect(301, "/admin/tags")
		} else {
			c.AbortWithStatus(400)
		}
	} else {
		c.AbortWithStatus(404)
	}
}

// updateTag updates the given values to the DB and returns it back
func updateTag(openingsID int, name string) (*Tag, error) {
	_, err := db.Exec("UPDATE blogTags SET name = ?, updatedAt = ? WHERE id = ?", name, time.Now(), openingsID)
	if err != nil {
		return nil, err
	}
	return &Tag{ID: openingsID, TagName: name}, nil
}

// deleteTag deletes the tag details with the particular selected id
func deleteTag(c *gin.Context) {
	if tagID, err := strconv.Atoi(c.Param("id")); err == nil {
		_, err := db.Exec("DELETE FROM blogTags where id= ?", tagID)
		if err == nil {
			c.Redirect(307, "/admin/tags")
		}
	} else {
		c.AbortWithStatus(404)
	}
}
