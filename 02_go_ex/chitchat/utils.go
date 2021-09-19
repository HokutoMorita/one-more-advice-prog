package main

import (
    "encoding/json"
    "errors"
    "fmt"
    // "github.com/mushahiroyuki/gowebprog/ch02/chitchat/data"
    "chitchat/data"
    "html/template"
    "log"
    "net/http"
    "os"
    "strings"
)

type Configuration struct {
    Address string
    ReadTimeout int64
    WriteTimeout int64
    Static string
}

var config Configuration
var logger *log.Logger

// 標準出力にログを出力させる関数
func p(a ...interface{}) {
    fmt.Println(a)
}

func init() {
    loadConfig()
    file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("Failed to open log file", err)
    }
    logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
    file, err := os.Open("config.json")
    if err != nil {
        log.Fatalln("Cannot open config file", err)
    }
    decoder := json.NewDecoder(file)
    config = Configuration{}
    err = decoder.Decode(&config)
    if err != nil {
        log.Fatalln("Cannot get configuration from file", err)
    }
}

// エラーメッセージページにリダイレクトさせる関数
func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
    url := []string{"/err?msg=", msg}
    http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// ユーザーがログインしてる かつ セッションがあることを判別する関数
func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {
    cookie, err := request.Cookie("_cookie")
    if err == nil {
        sess = data.Session{Uuid: cookie.Value}
        if ok, _ := sess.Check(); !ok {
            err = errors.New("Invalid session")
        }
    }
    return
}

// HTMLテンプレートをパースする関数
func parseTemplateFiles(filenames ...string) (t *template.Template) {
    var files []string
    t = template.New("layout")
    for _, file := range filenames {
        files = append(files, fmt.Sprintf("templates/%s.html", file))
    }
    t = template.Must(t.ParseFiles(files...))
    return
}

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string)  {
    var files []string
    for _, file := range filenames {
        files = append(files, fmt.Sprintf("templates/%s.html", file))
    }
    templates := template.Must(template.ParseFiles(files...))
    templates.ExecuteTemplate(writer, "layout", data)
}

// ログ情報
func info(args ...interface{}) {
    logger.SetOutput(os.Stdout)
    logger.SetPrefix("INFO ")
    logger.Println(args...)
}

func danger(args ...interface{}) {
    logger.SetOutput(os.Stdout)
    logger.SetPrefix("ERROR ")
    logger.Println(args...)
}

func warning(args ...interface{}) {
    logger.SetOutput(os.Stdout)
    logger.SetPrefix("WARNING ")
    logger.Println(args...)
}

// version
func version() string {
    return "0.1"
}
