package main

import (
    "fmt"
    "github.com/mushahiroyuki/gowebprog/ch02/chitchat/data"
    "net/http"
)

// GET /threads/new
// 新しいスレッドページを表示する
func newThread(writer http.ResponseWriter, request *http.Request) {
    _, err := session(writer, request)
    if err != nil {
        http.Redirect(writer, request, "/login", 302)
    } else {
        generateHTML(writer, nil, "layout", "private.navbar", "new.thread")
    }
}

// POST /signup
// ユーザーアカウントを作成する
func createThread(writer http.ResponseWriter, request *http.Request) {
    sess, err := session(writer, request)
    if err != nil {
        http.Redirect(writer, request, "/login", 302)
    } else {
        err = request.ParseForm()
        if err != nil {
            danger(err, "Cannot parse form")
        }
        user, err := sess.User()
        if err != nil {
            danger(err, "Cannnot parse form")
        }
        user, err := sess.User()
        if err != nil {
            danger(err, "Cannot get user from session")
        }
        topic := request.PostFormValue("topic")
        if _, err := user.CreateThread(topic); err != nil {
            danger(err, "Cannot create thread")
        }
        http.Redirect(writer, request, "/", 302)
    }
}

// GET /thread/read
// スレッドの詳細を表示する
func readThread(writer http.ResponseWriter, request *http.Request) {
    vals := request.URL.Query()
    uuid := vals.Get("id")
    thread, err := data.ThreadByUUID(uuid)
    if err != nil {
        error_message(writer, request, "Cannot read thread")
    } else {
        _, err := session(writer, request)
        if err != nil {
            generateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
        } else {
            generateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
        }
    }
}

// POST /thread/post
// 投稿を作成する
func postThread(writer http.ResponseWriter, request *http.Request) {
    sess, err := session(writer, request)
    if err != nil {
        http.Redirect(writer, request, "/login", 302)
    } else {
        err = request.ParseForm()
        if err != nil {
            danger(err, "Cannot parse form")
        }
        user, err := sess.User()
        if err != nil {
            danger(err, "Cannot get user from session")
        }
        body := request.PostFormValue("body")
        uuid := request.PostFormValue("uuid")
        thread, err := data.ThreadByUUID(uuid)
        if err != nil {
            error_message(writer, request, "Cannot read thread")
        }
        if _, err := user.CreatePost(thread, body); err != nil {
            danger(err, "Cannot create post")
        }
        url := fmt.Sprint("/thread/read?id=", uuid)
        http.Redirect(writer, request, url, 302)
    }
}
