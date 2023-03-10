package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"traces/research"
)

type DataOutput struct {
	Result string
	Data   [][][]int
}

type DataResearch struct {
	Result string
	Levels research.Levels
}

func Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	home := filepath.Join("html", "home.html")

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

func Research(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res := filepath.Join("html", "research.html")

	tmpl, err := template.ParseFiles(res)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = tmpl.ExecuteTemplate(w, "res", nil) //data - передаваемый объект в шаблон
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func ResearchResults(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	results := filepath.Join("html", "results.html")

	tmpl, err := template.ParseFiles(results)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = tmpl.ExecuteTemplate(w, "results", nil) //data - передаваемый объект в шаблон
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
	r.GET("/research", Research)
	r.GET("/research/results", ResearchResults)

	r.POST("/get_words", getWords)
	//r.POST("/get_fingers", getFingers)
	r.POST("/get_research", getResearch)
	r.GET("/download_file", downloadFile)

	fmt.Println("Сервер запущен. Перейдите по адресу http://localhost:8181/")
	err := http.ListenAndServe(":8181", r)
	if err != nil {
		panic(err)
	}
}

func downloadFile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Filename := "last_level.txt"

	Openfile, err := os.Open(Filename) //Open the file to be downloaded later
	defer Openfile.Close()             //Close after function return

	if err != nil {
		http.Error(w, "File not found.", 404) //return 404 if file is not found
		return
	}

	tempBuffer := make([]byte, 512) //Create a byte array to read the file later
	_, err = Openfile.Read(tempBuffer)
	if err != nil {
		panic(err)
	} //Read the file into  byte
	FileContentType := http.DetectContentType(tempBuffer) //Get file header

	FileStat, _ := Openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	//Set the headers
	w.Header().Set("Content-Type", FileContentType+";"+Filename)
	w.Header().Set("Content-Length", FileSize)

	Openfile.Seek(0, 0)  //We read 512 bytes from the file already so we reset the offset back to 0
	io.Copy(w, Openfile) //'Copy' the file to the client
}

func getResearch(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var responseData DataResearch
	responseData.Levels = research.Get()
	fmt.Println(responseData.Levels)
	responseData.Result = "ok"

	_ = json.NewEncoder(w).Encode(responseData)
	return
}

func getWords(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var responseData DataOutput
	responseData.Data = research.GetWords()
	fmt.Println(responseData.Data)
	responseData.Result = "ok"

	_ = json.NewEncoder(w).Encode(responseData)
	return
}

func handleRequest() {
	//http.HandleFunc("/", homePage)
	r := httprouter.New()
	routes(r)
}

func main() {
	handleRequest()
}
