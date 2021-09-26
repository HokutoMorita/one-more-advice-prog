package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

// connect to the Db
func init() {
    var err error
    Db, err = sql.Open("mysql", "root:one_more_advice_prog@tcp(127.0.0.1:4307)/one_more_advice_prog?collation=utf8mb4_unicode_ci&parseTime=true")
    if err != nil {
        panic(err)
    }
    return
}

// Get a single post
func retrieve(id int) (post Post, err error) {
    post = Post{}
    err = Db.QueryRow("SELECT id, content, author FROM posts WHERE id = ?", id).
        Scan(&post.Id, &post.Content, &post.Author)
    return
}

// Create a new post
func (post *Post) create() (err error) {
    statement := "INSERT INTO posts (content, author) VALUES (?, ?)"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()
    _, err = stmt.Exec(post.Content, post.Author)
    return
}

// Update a post
func (post *Post) update() (err error) {
    _, err = Db.Exec("UPDATE posts set content = ?, author = ? WHERE id = ?", post.Content, post.Author, post.Id)
    return
}

// Delete a post
func (post *Post) delete() (err error) {
    _, err = Db.Exec("DELETE FROM posts WHERE id = ?", post.Id)
    return
}
