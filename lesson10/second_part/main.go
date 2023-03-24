package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)
type Response []struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
    url := "https://jsonplaceholder.typicode.com/posts"

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

    fmt.Println(response)
}