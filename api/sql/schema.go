package sql

const Schema = `
CREATE TABLE IF NOT EXISTS DE_TAGS
  (id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid TEXT,
  tag TEXT);

CREATE TABLE IF NOT EXISTS DE_BOOK_TAGS
  (id INTEGER PRIMARY KEY AUTOINCREMENT,
  book_uuid TEXT,
  tag_uuid TEXT,
  FOREIGN KEY(book_uuid) REFERENCES DE_BOOKS(uuid),
  FOREIGN KEY(tag_uuid) REFERENCES DE_TAGS(uuid));

CREATE TABLE IF NOT EXISTS DE_BOOKS
  (id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid TEXT,
  url TEXT,
  desc TEXT,
  added_at INTEGER);

CREATE VIEW IF NOT EXISTS V_DE_TAG_GROUPS
  AS SELECT b.uuid uuid, group_concat(t.tag) tags
  FROM DE_BOOK_TAGS bt
  JOIN DE_BOOKS b ON b.uuid = bt.book_uuid
  JOIN DE_TAGS t ON bt.tag_uuid = t.uuid group by b.uuid;

CREATE VIEW IF NOT EXISIS V_DE_BOOKS
  AS SELECT b.url, b.desc, t.tags
  FROM DE_BOOKS b
  JOIN V_DE_TAG_GROUPS t ON b.uuid = t.uuid;
`
