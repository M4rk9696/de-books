package sql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	uuid "github.com/google/uuid"
)

func Migrate(db *sql.DB) error {
	_, err := db.Exec(Schema)
	if err != nil {
		fmt.Println("File reading error", err)
		return err
	}
	return nil
}

type Bookmark struct {
	URL     string   `json:"url"`
	Desc    string   `json:"desc"`
	AddedAt int64    `json:"addedAt"`
	Tags    []string `json:"tags"`
}

func QueryBookmarks(db *sql.DB, tags []string) ([]Bookmark, error) {
	stmt, err := db.Prepare(`select b.uuid, b.url, b.desc, b.added_at, t.tag
		from DE_BOOKS as b
		join DE_BOOK_TAGS as bt on b.uuid = bt.book_uuid
		join DE_TAGS as t on bt.tag_uuid = t.uuid
		where t.tag in (?)`)
	if err != nil {
		fmt.Println("Error creating prepared statement", err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(strings.Join(tags, ","))
	if err != nil {
		log.Println("Error while querying", err)
		return nil, err
	}
	defer rows.Close()

	bookmarks := make(map[string]Bookmark)
	for rows.Next() {
		var id string
		var url string
		var desc string
		var addedAt int64
		var tag string
		err = rows.Scan(&id, &url, &desc, &addedAt, &tag)
		if err != nil {
			fmt.Println("Error while scanning", err)
			return nil, err
		}
		bookmark, ok := bookmarks[id]
		if !ok {
			bookmark = Bookmark{
				url,
				desc,
				addedAt,
				[]string{},
			}
		}
		bookmark.Tags = append(bookmark.Tags, tag)
		bookmarks[id] = bookmark
	}

	bks := []Bookmark{}
	for _, v := range bookmarks {
		bks = append(bks, v)
	}

	return bks, nil
}

var allTags = map[string]string{}

func FetchAllTags(db *sql.DB) {
	rows, err := db.Query("select uuid, tag from DE_TAGS")
	if err != nil {
		log.Println("Error while querying", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		var tag string
		err = rows.Scan(&id, &tag)
		if err != nil {
			fmt.Println("Error while scanning", err)
			return
		}
		allTags[tag] = id
	}
}

func AddBookmark(db *sql.DB, url string, desc string, tags []string) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println("Error creating transaction", err)
		return err
	}

	stmt, err := tx.Prepare("insert into DE_BOOKS (url, uuid, desc, added_at) values (?, ?, ?, ?)")
	if err != nil {
		log.Println("Error creating prepared statement for DE_BOOKS", err)
		return err
	}
	defer stmt.Close()

	bookmarkUUID := uuid.New().String()
	_, err = stmt.Exec(url, bookmarkUUID, desc, time.Now().Unix())
	if err != nil {
		log.Println("Error executing prepared statement for DE_BOOKS", err)
		return err
	}

	for _, tag := range tags {
		tagUUID, ok := allTags[tag]
		// Insert tag if not present
		if !ok {
			tagUUID = uuid.New().String()
			stmt, err = tx.Prepare("insert into DE_TAGS (uuid, tag) values (?, ?)")
			if err != nil {
				log.Println("Error creating prepared statement for DE_TAGS", err)
				return err
			}
			defer stmt.Close()
			_, err = stmt.Exec(tagUUID, tag)
			if err != nil {
				log.Println("Error executing prepared statement for DE_TAGS", err)
				return err
			}
			allTags[tag] = tagUUID
		}

		stmt, err = tx.Prepare("insert into DE_BOOK_TAGS (book_uuid, tag_uuid) values (?, ?)")
		if err != nil {
			log.Println("Error creating prepared statement for DE_BOOK_TAG", err)
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(bookmarkUUID, tagUUID)
		if err != nil {
			log.Println("Error executing prepared statement for DE_BOOK_TAG", err)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Error commiting transaction", err)
		return err
	}
	return nil
}
