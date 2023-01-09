package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"traces/controller"
)

func routes(r *httprouter.Router) {
	//путь к папке со внешними файлами: html, js, css, изображения и т.д.
	r.ServeFiles("/templates/*filepath", http.Dir("templates"))
	//что следует выполнять при входящих запросах указанного типа и по указанному адресу
	r.GET("/", controller.Home)
}

func handleRequest() {
	//http.HandleFunc("/", homePage)
	r := httprouter.New()
	routes(r)

	err := http.ListenAndServe(":8181", r)
	if err != nil {
		panic(err)
	}
}

func main() {
	handleRequest()
}
