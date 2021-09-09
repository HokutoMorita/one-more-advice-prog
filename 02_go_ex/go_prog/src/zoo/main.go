package main

import (
    "fmt"

    "zoo/animals"

)

func main() {
    /* 関数AppNameの呼び出しを追加 */
    fmt.Println(AppName())

    fmt.Println(animals.ElephantFeed())
    fmt.Println(animals.MonkeyFeed())
    fmt.Println(animals.RabbitFeed())
}
