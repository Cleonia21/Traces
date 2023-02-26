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

type DataInput struct {
	Type string
}

type DataOutput struct {
	Result string
	Data   [][][]int
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
	//r.POST("/get_fingers", getFingers)
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

func getWords(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var responseData DataOutput
	responseData.Data = research.GetWords()
	fmt.Println(responseData.Data[2])
	responseData.Result = "ok"

	//var data DataInput
	//err := json.NewDecoder(r.Body).Decode(&data)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(data.Type)

	_ = json.NewEncoder(w).Encode(responseData)
	return
}

//
//func writeFile(d DataOutput) {
//
//	file, err := os.Create("last_level.txt")
//	if err != nil {
//		fmt.Println("Unable to create file:", err)
//		os.Exit(1)
//	}
//	defer file.Close()
//
//	for j := 0; j < len(strs); j++ {
//		_, err = file.WriteString(strings.Join(strs[j], " ") + "\n")
//		if err != nil {
//			fmt.Println("Unable to write file:", err)
//			os.Exit(1)
//		}
//	}
//	_, _ = file.WriteString("\n")
//}

func handleRequest() {
	//http.HandleFunc("/", homePage)
	r := httprouter.New()
	routes(r)
}

func main() {
	handleRequest()
}
