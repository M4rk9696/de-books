package main

import (
	"database/sql"
	"log"
	"os"

	connector "github.com/m4rk9696/de-books/api"
	s "github.com/m4rk9696/de-books/api/sql"
	apiV1 "github.com/m4rk9696/de-books/api/v1"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./dedb.db")
	if err != nil {
		log.Println("Error opening DB", err)
	}
	defer db.Close()

	s.Migrate(db)
	s.FetchAllTags(db)

	// TODO: Remove this with a proper fix
	if os.Getenv("mode") != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(connector.DBConnector(db))

	v1 := router.Group("/v1")
	{
		v1.GET("/bookmarks", apiV1.FetchBookmarks)
		v1.POST("/add", apiV1.AddBookmark)
	}

	router.Run(":8080")
}
