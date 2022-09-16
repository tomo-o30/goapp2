package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(10))
	t.Execute(w, rand.Intn(10) > 5)
}
