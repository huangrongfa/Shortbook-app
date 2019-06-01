package main

import (
    "io"
    "net/http"
)

func propxy(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "<h1>欢迎大家进入GO的学习殿堂欢迎大家进入GO的学习殿堂111222</h1>")
}

func main()  {
    http.HandleFunc("/", propxy)
    http.ListenAndServe(":8080", nil)
}
