package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
	"traces/levels"
)

type DataInput struct {
	Text string
}

type DataOutput struct {
	Result string
	Data   [][][]string
}

func Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	home := filepath.Join("html", "levels.html")

	tmpl, err := template.ParseFiles(home)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = tmpl.ExecuteTemplate(w, "home", nil) //data - передаваемый объект в шаблон
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func routes(r *httprouter.Router) {
	//путь к папке со внешними файлами: html, js, css, изображения и т.д.
	r.ServeFiles("/css/*filepath", http.Dir("css"))
	r.ServeFiles("/js/*filepath", http.Dir("js"))
	r.ServeFiles("/img/*filepath", http.Dir("img"))
	//что следует выполнять при входящих запросах указанного типа и по указанному адресу
	r.GET("/", Home)

	r.POST("/get_words", getWords)
	r.POST("/get_fingers", getFingers)

	err := http.ListenAndServe(":8181", r)
	if err != nil {
		panic(err)
	}
}

func getWords(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	const firstSeriesLen = 6
	const maxSeriesLen = 10

	var responseData DataOutput
	responseData.Data = make([][][]string, maxSeriesLen-firstSeriesLen+1)
	responseData.Result = "ok"

	var l levels.Level
	for i := 0; i < maxSeriesLen-firstSeriesLen+1; i++ {
		l.New(i + 6)
		responseData.Data[i] = l.GetWords()
	}
	_ = json.NewEncoder(w).Encode(responseData)
	return
}

func getFingers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	const firstSeriesLen = 6
	const maxSeriesLen = 10

	var responseData DataOutput
	responseData.Data = make([][][]string, maxSeriesLen-firstSeriesLen+1)
	responseData.Result = "ok"

	var l levels.Level
	for i := 0; i < maxSeriesLen-firstSeriesLen+1; i++ {
		l.New(i + 6)
		responseData.Data[i] = l.GetWords()
	}
	_ = json.NewEncoder(w).Encode(responseData)
	return
}

//func writeFile() {
//	file, err := os.Create("hello.txt")
//
//}

func handleRequest() {
	//http.HandleFunc("/", homePage)
	r := httprouter.New()
	routes(r)
}

func main() {
	handleRequest()
}
