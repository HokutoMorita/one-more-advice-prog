package data

import (
    "time"
)

type User struct {
    Id int
    Uuid string
    Name string
    Email string
    Password string
    CreatedAt time.Time
}

type Session struct {
    Id int
    Uuid string
    Email string
    UserId int
    CreatedAt time.Time
}

// ユーザーセッションを作成する
func (user *User) CreateSession() (session Session, err error) {
    statement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, user_id, created_at"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()

    // use QueryRow to return a row and scan the returned id into the Session struct
    err = stmt.QueryRow(createdUUID(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
    return
}

// ユーザーセッションを取得する
func (user *User) Session() (session Session, err error) {
    session = Session{}
    err = Db.QueryRow(
        "SELECT
            id,
            uuid,
            email,
            user_id,
            created_at
        FROM
            sessions
        WHERE
            user_id = $1"
        , user.Id
    ).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
    return
}

// ユーザーセッションのバリデーションをする
func (session *Session) Check() (valid bool, err error) {
    err = Db.QueryRow(
        "SELECT
            id,
            uuid,
            email,
            user_id,
            created_at
        FROM
            sessions
        WHERE
            uuid = $1"
        , session.Uuid
    ).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
    if err != nil {
        valid = false
        return
    }
    if session.Id != 0 {
        valid = true
    }
    return
}

// uuidからユーザーセッションを削除する
func (session *Session) DeleteByUUID() (err error) {
    statement := "delete from sessions where uuid = $1"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()

    _, err = stmt.Exec(session.Uuid)
    return
}

// ユーザーセッションからユーザー情報を取得する
func (session *Session) User() (user User, err error) {
    user = User{}
    err = Db.QueryRow(
        "SELECT
            id,
            uuid,
            name,
            email,
            created_at
        FROM
            users
        WHERE
            id = $1"
        , session.UserId
    ).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
    return
}

// すべてのユーザーセッションを削除する
func SessionDeleteAll() (err error) {
    statement := "delete from sessions"
    _, err = Db.Exec(statement)
    return
}

// ユーザー情報を作成する
func (user *User) Create() (err error) {
    // Postgres does not automatically return the last insert id, because it would be wrong to assume
    // you're always using a sequence.You need to use the RETURNING keyword in your insert to get this
    // information from postgres.
    statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()
    // use QueryRow to return a row and scan the returned id into the User struct
    err = stmt.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
    return
}

// ユーザー情報を削除する
func (user *User) Delete() (err error) {
    statement := "delete from users where id = $1"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()
    _, err = stmt.Exec(user.Id)
    return
}

// ユーザー情報を更新する
func (user *User) Update() (err error) {
    statement := "update users set name = $2, email = $3 where id = $1"
    stmt, err := Db.Prepare(statement)
    if err != nil {
        return
    }
    defer stmt.Close()
    _, err = stmt.Exec(user.Id, user.Name, user.Email)
    return
}

// 全てのユーザー情報を削除する
func UserDeleteAll() (err error) {
    statement := "delete from users"
    _, err = Db.Exec(statement)
    return
}

// 全てのユーザー情報を取得する
func Users() (users []User, err error) {
    rows, err := Db.Query(
        "SELECT
            id,
            uuid,
            name,
            email,
            password,
            created_at
        FROM
            users"
    )
    if err != nil {
        return
    }
    for rows.Next() {
        user := User{}
        if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
            return
        }
        users = append(users, user)
    }
    rows.Close()
    return
}

// Eメールアドレスからユーザー情報を取得する
func UserByEmail(email string) (user User, err error) {
    user = User{}
    err = Db.QueryRow(
        "SELECT
            id,
            uuid,
            name,
            email,
            password,
            created_at
        FROM
            users
        WHERE
            email = $1"
        , email
    ).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
    return
}

// uuidからユーザー情報を取得する
func UserByUUID(uuid string) (user User, err error) {
    user = User{}
    err = Db.QueryRow(
        "SELECT
            id,
            uuid,
            name,
            email,
            password,
            created_at
        FROM
            users
        WHERE
            uuid = $1"
        , uuid
    ).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
    return
}
