package main

import (
    "encoding/json"
    "fmt"
    "os"
)

func main() {

    argsWithProg := os.Args
    fmt.Println(argsWithProg)

    json := os.Args[1]
    fmt.Println(json)
    var j interface{}
    err := json.Unmarshal(json, &j)
    fmt.Println(j)

}
