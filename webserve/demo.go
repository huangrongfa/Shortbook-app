package main

import (
    "io"
    "net/http"
)

func propxy(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "<h1>欢迎学习go语言哈哈</h1>")
}

func main()  {
    http.HandleFunc("/", propxy)
    http.ListenAndServe(":8080", nil)
}