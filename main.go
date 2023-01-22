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
	"strings"
	"traces/levels"
)

type DataInput struct {
	Type string
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
	r.GET("/download_file", downloadFile)

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

/*
var data TimeDataInput
   err := json.NewDecoder(request.Body).Decode(&data)
   if err != nil {
      fmt.Println(err.Error())
      return
   }
   fmt.Println(data.Name)
   fmt.Println(data.Time)
*/

func getWords(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	const firstSeriesLen = 6
	const maxSeriesLen = 10

	var responseData DataOutput
	responseData.Data = make([][][]string, maxSeriesLen-firstSeriesLen+1)
	responseData.Result = "ok"

	file, err := os.Create("last_level.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var data DataInput
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(data.Type)

	var l levels.Level
	for i := 0; i < maxSeriesLen-firstSeriesLen+1; i++ {
		l.New(i + 6)
		var strs [][]string
		switch data.Type {
		case "1":
			strs = l.GetWords()
		case "2":
			strs = l.GetLetters()
		case "3":
			strs = l.GetNums()
		}
		responseData.Data[i] = strs

		for j := 0; j < len(strs); j++ {
			_, err = file.WriteString(strings.Join(strs[j], " ") + "\n")
			if err != nil {
				fmt.Println("Unable to write file:", err)
				os.Exit(1)
			}
		}
		_, _ = file.WriteString("\n")
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
		responseData.Data[i] = l.GetNums()
	}
	_ = json.NewEncoder(w).Encode(responseData)
	return
}

//func writeFile(d DataOutput) {
//
//	for i := 0; i < len(d.Data); i++ {
//
//	}
//}

func handleRequest() {
	//http.HandleFunc("/", homePage)
	r := httprouter.New()
	routes(r)
}

func main() {
	handleRequest()
}
