package database

const createTable = `
CREATE TABLE IF NOT EXISTS posts
(
	id SERIAL PRIMARY KEY,
	title TEXT,
	content TEXT,
	author TEXT
)
`

var insertPostTable = `
INSERT INTO posts(title, content, author) VALUES($1,$2,$3) RETURNING id
`
