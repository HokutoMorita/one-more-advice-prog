package data

import (
    "time"
)

type Thread struct {
    Id int
    Uuid string
    Topic string
    UserId int
    CreatedAt time.Time
}

type Post struct {
    Id int
    Uuid string
    Body string
    UserId int
    ThreadId int
    CreatedAt time.Time
}

// CreatedAtをページ上で表示できるように整形するメソッド
func (thread *Thread) CreatedAtDate() string {
    return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (post *Post) CreatedAtDate() string {
    return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// スレッド内の投稿番号を取得するメソッド
func (thread *Thread) NumReplies() (count int) {
    rows, err := Db.Query("SELECT count(*) FROM posts WHERE thread_id = $1", thread.Id)
    if err != nil {
        // countは0がリターンされる？？
        return
    }
    for rows.Next() {
        if err = rows.Scan(&count); err != nil {
            // countはインクリメントされた値がリターンされる？？
            return
        }
    }
    rows.Close()
    // countはインクリメントされた値がリターンされる？？
    return
}

// スレッドに投稿した投稿データを取得するメソッド
func (thread *Thread) Posts() (posts []Post, err error) {
    queryStr := `
        SELECT 
            id,
            uuid,
            body,
            user_id,
            thread_id,
            created_at
        FROM 
            posts
        WHERE 
            thread_id = $1`
    rows, err := Db.Query(queryStr, thread.Id)
    if err != nil {
        // 空のリストが返される？？
        return
    }
    for rows.Next() {
        post := Post{}
        if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
            // この時点で追加されてるpostsを返す
            return
        }
        posts = append(posts, post)
    }
    rows.Close()
    return
}

// 新しいスレッドを作成するメソッド
func (user *User) CreateThread(topic string) (conv Thread, err error) {
    // statement := "insert into threads (uuid, topic, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, topic, user_id, created_at"
    statement := "insert into threads (uuid, topic, user_id, created_at) values (?, ?, ?, ?)"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()
    // use QueryRow to return a row and scan the returned id into the Session struct
    // err = stmt.QueryRow(createUUID(), topic, user.Id, time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
    _, err = stmt.Exec(createUUID(), topic, user.Id, time.Now())
    return
}

// スレッドに新しい投稿を作成するメソッド
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
    statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, thread_id, created_at"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()
    // use QueryRow to return a row and scan the returned id into the Session struct
    err = stmt.QueryRow(createUUID(), body, user.Id, conv.Id, time.Now()).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
    return
}

// すべてのスレッドを取得する
func Threads() (threads []Thread, err error) {
    queryStr := `
        SELECT
                id,
                uuid,
                topic,
                user_id,
                created_at
            FROM
                threads
            ORDER BY
                created_at DESC`
    rows, err := Db.Query(queryStr)
    if err != nil {
        return
    }
    for rows.Next() {
        conv := Thread{}
        if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
            return
        }
        threads = append(threads, conv)
    }
    rows.Close()
    return
}

// UUIDに対応するスレッドを取得する
func ThreadByUUID(uuid string) (conv Thread, err error) {
    conv = Thread{}
    queryStr := `
        SELECT
                id,
                uuid,
                topic,
                user_id,
                created_at
            FROM
                threads
            WHERE
                uuid = ?`
    err = Db.QueryRow(queryStr, uuid).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
    return
}

// このスレッドを始めたユーザーを取得する
func (thread *Thread) User() (user User) {
    user = User{}
    queryStr := `
    SELECT
            id,
            uuid,
            name,
            email,
            created_at
        FROM
            users
        WHERE
            id = $1`
    Db.QueryRow(queryStr, thread.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
    return
}

// この投稿を書いたユーザーを取得する
func (post *Post) User() (user User) {
    user = User{}
    queryStr := `
        SELECT
                id,
                uuid,
                name,
                email,
                created_at
            FROM
                users
            WHERE
                id = $1`
    Db.QueryRow(queryStr, post.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
    return
}
