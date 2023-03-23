package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Response struct {
    Data []string `json:"data"`
}

func main() {
    url := "https://example.com/data.json"

    resp, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    var response Response
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        panic(err)
    }

    fmt.Println(response.Data)
}