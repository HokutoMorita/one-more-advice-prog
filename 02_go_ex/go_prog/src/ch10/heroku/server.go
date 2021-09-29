package main

import (
    "encoding/json"
    "net/http"
    "path"
    "strconv"
    "fmt"
    "os"
)

type Post struct {
    Id int `json:"id"`
    Content string `json:"content"`
    Author string `json:"author"`
}

func main() {
    server := http.Server{
        Addr: ":" + os.Getenv("PORT"),
    }
    http.HandleFunc("/post/", handleRequest)
    server.ListenAndServe()
}

// main handler function
func handleRequest(w http.ResponseWriter, r *http.Request) {
    var err error
    switch r.Method {
    case "GET":
        err = handleGet(w, r)
    case "POST":
        err = handlePost(w, r)
    case "PUT":
        err = handlePut(w, r)
    case "DELETE":
        err = handleDelete(w, r)
    }
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

// Retrieve a post
// GET /post/1
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
    id, err := strconv.Atoi(path.Base(r.URL.Path))
    if err != nil {
        fmt.Println(err)
        return
    }
    post, err := retrieve(id)
    if err != nil {
        fmt.Println(err)
        return
    }
    output, err := json.MarshalIndent(&post, "", "    ")
    if err != nil {
        fmt.Println(err)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(output)
    return
}

// Create a post
// POST /post/
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
    len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    var post Post
    json.Unmarshal(body, &post)
    err = post.create()
    if err != nil {
        fmt.Println(err)
        return
    }
    w.WriteHeader(200)
    return
}

// Update a post
// PUT /post/1
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
    id, err := strconv.Atoi(path.Base(r.URL.Path))
    if err != nil {
        fmt.Println(err)
        return
    }

    // 既存のデータの取得
    post, err := retrieve(id)
    if err != nil {
        fmt.Println(err)
        return
    }

    // 更新用のデータをリクエスト内から取得
    len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    json.Unmarshal(body, &post)

    // 更新
    err = post.update()
    if err != nil {
        fmt.Println(err)
        return
    }
    w.WriteHeader(200)
    return
}

// Delete a post
// DELETE /post/1
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
    id, err := strconv.Atoi(path.Base(r.URL.Path))
    if err != nil {
        fmt.Println(err)
        return
    }

    // 既存のデータの取得
    post, err := retrieve(id)
    if err != nil {
        fmt.Println(err)
        return
    }

    // 削除
    err = post.delete()
    if err != nil {
        fmt.Println(err)
        return
    }
    w.WriteHeader(200)
    return
}
