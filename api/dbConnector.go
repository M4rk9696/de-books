package api

import (
	"github.com/gin-gonic/gin"
	"database/sql"
)

func DBConnector(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("connection", db)
		c.Next()
	}
}
