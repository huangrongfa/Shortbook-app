package main

import (
	"io" 
	"net/http" 
	"os/exec" 
	"log"
)

func reLaunch () {
	cmd := exec.Command("sh", "./depley.sh")
	err := cmd.Start()
	if err!= nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func propxy(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>欢迎学习go语言,并实现自动化部署</h1>")
	reLaunch()
}

func main()  {
	http.HandleFunc("/", propxy)
	http.ListenAndServe(":5000", nil)
}