package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

// connect to the Db
func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/gwp")
	if err != nil {
		panic(err)
	}
}

// Get a single post
func retrieve(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// Create a new post
func (post *Post) create() (err error) {
	statement := "insert into posts (content, author) values (?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

// Update a post
func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set content = ?, author = ? where id = ?", post.Id, post.Content, post.Author)
	return
}

// Delete a post
func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = ?", post.Id)
	return
}
