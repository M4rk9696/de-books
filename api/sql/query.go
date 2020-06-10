package sql

const ToReadable = `select printf("%s	%s	%s", b.desc, b.url, group_concat(t.tag))
from DE_BOOKS AS b
join DE_BOOK_TAGS as bt on b.uuid = bt.book_uuid
join DE_TAGS as t on bt.tag_uuid = t.uuid
group by b.uuid`
