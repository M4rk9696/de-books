package sql_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/m4rk9696/de-books/api/sql"
)

var _ = Describe("Sql", func() {

	const writePath

	Describe("UpdateReadable", func() {
		var *sql.DB db
		BeforeEach(func() {
			// Create in memory DB
			db, err := sql.Open("sqlite3", "file:test.db?mode=memory")
			if err != nil {
				log.Fatal("Error opening DB", err)
			}
		})

		AfterEach(func() {
			db.Close()
		})

		It("should be able to serialize db", func() {
			
		})
	})

})
