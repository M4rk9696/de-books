package main

import (
	"database/sql"
	"log"
	"os"
	"os/user"
	"path"

	connector "github.com/m4rk9696/de-books/api"
	s "github.com/m4rk9696/de-books/api/sql"
	apiV1 "github.com/m4rk9696/de-books/api/v1"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	dbPath := path.Join(usr.HomeDir, ".de-books")
	os.MkdirAll(dbPath, os.ModePerm)
	db, err := sql.Open("sqlite3", path.Join(dbPath, "dedb.db"))
	if err != nil {
		log.Fatal("Error opening DB", err)
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
