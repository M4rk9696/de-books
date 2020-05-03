package v1

import (
	"database/sql"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	s "github.com/m4rk9696/de-books/api/sql"
)

func FetchBookmarks(c *gin.Context) {
	tags := strings.Split(c.DefaultQuery("tags", "default"), ",")
	db, ok := c.Get("connection")
	// Create a method to create new connection
	if !ok {
		c.JSON(402, gin.H{
			"error": "Internal Server Error",
		})
		log.Fatal("Connection to db not present")
	}

	res, err := s.QueryBookmarks(db.(*sql.DB), tags)
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Error fetching bookmarks",
		})
	}

	c.JSON(200, res)
}

func AddBookmark(c *gin.Context) {
	url := c.PostForm("url")
	if url == "" {
		c.JSON(401, gin.H{
			"error": "URL not present",
		})
		return
	}
	desc := c.PostForm("desc")
	tags := strings.Split(c.DefaultPostForm("tags", "default"), ",")
	db, ok := c.Get("connection")
	// Create a method to create new connection
	if !ok {
		c.JSON(402, gin.H{
			"error": "Internal Server Error",
		})
		log.Fatal("Connection to db not present")
	}

	s.AddBookmark(db.(*sql.DB), url, desc, tags)
	c.JSON(200, gin.H{
		"result": "success",
	})
}
