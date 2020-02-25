package main

import (
    "encoding/json"
    "fmt"
    "os"
)

func main() {

    argsWithProg := os.Args
    fmt.Println(argsWithProg)

    pluginJson := []byte(os.Args[1])
    fmt.Println(pluginJson)
    var j interface{}
    err := json.Unmarshal(pluginJson, &j)
    fmt.Println(err)
    fmt.Println(j)

}
