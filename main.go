package main

import "fmt"

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/julienschmidt/httprouter"
//	"html/template"
//	"io"
//	"net/http"
//	"os"
//	"path/filepath"
//	"strconv"
//	"strings"
//	"traces/research"
//)
//
//type DataInput struct {
//	Type string
//}
//
//type DataOutput struct {
//	Result string
//	Data   [][][]string
//}
//
//func Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	home := filepath.Join("html", "research.html")
//
//	tmpl, err := template.ParseFiles(home)
//	if err != nil {
//		http.Error(w, err.Error(), 400)
//		return
//	}
//
//	err = tmpl.ExecuteTemplate(w, "home", nil) //data - передаваемый объект в шаблон
//	if err != nil {
//		http.Error(w, err.Error(), 400)
//		return
//	}
//}
//
//func routes(r *httprouter.Router) {
//	//путь к папке со внешними файлами: html, js, css, изображения и т.д.
//	r.ServeFiles("/css/*filepath", http.Dir("css"))
//	r.ServeFiles("/js/*filepath", http.Dir("js"))
//	r.ServeFiles("/img/*filepath", http.Dir("img"))
//	//что следует выполнять при входящих запросах указанного типа и по указанному адресу
//	r.GET("/", Home)
//
//	r.POST("/get_words", getWords)
//	r.POST("/get_fingers", getFingers)
//	r.GET("/download_file", downloadFile)
//
//	fmt.Println("Сервер запущен. Перейдите по адресу http://localhost:8181/")
//	err := http.ListenAndServe(":8181", r)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func downloadFile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	Filename := "last_level.txt"
//
//	Openfile, err := os.Open(Filename) //Open the file to be downloaded later
//	defer Openfile.Close()             //Close after function return
//
//	if err != nil {
//		http.Error(w, "File not found.", 404) //return 404 if file is not found
//		return
//	}
//
//	tempBuffer := make([]byte, 512) //Create a byte array to read the file later
//	_, err = Openfile.Read(tempBuffer)
//	if err != nil {
//		panic(err)
//	} //Read the file into  byte
//	FileContentType := http.DetectContentType(tempBuffer) //Get file header
//
//	FileStat, _ := Openfile.Stat()                     //Get info from file
//	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string
//
//	//Set the headers
//	w.Header().Set("Content-Type", FileContentType+";"+Filename)
//	w.Header().Set("Content-Length", FileSize)
//
//	Openfile.Seek(0, 0)  //We read 512 bytes from the file already so we reset the offset back to 0
//	io.Copy(w, Openfile) //'Copy' the file to the client
//}
//
///*
//var data TimeDataInput
//  err := json.NewDecoder(request.Body).Decode(&data)
//  if err != nil {
//     fmt.Println(err.Error())
//     return
//  }
//  fmt.Println(data.Name)
//  fmt.Println(data.Time)
//*/
//
//func getWords(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	const firstSeriesLen = 3
//	const maxSeriesLen = 10
//
//	var responseData DataOutput
//	responseData.Data = make([][][]string, maxSeriesLen-firstSeriesLen+1)
//	responseData.Result = "ok"
//
//	file, err := os.Create("last_level.txt")
//	if err != nil {
//		fmt.Println("Unable to create file:", err)
//		os.Exit(1)
//	}
//	defer file.Close()
//
//	var data DataInput
//	err = json.NewDecoder(r.Body).Decode(&data)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	//fmt.Println(data.Type)
//
//	var l research.LevelParam
//	for i := 0; i < maxSeriesLen-firstSeriesLen+1; i++ {
//		l.New(i + firstSeriesLen)
//		var strs [][]string
//		switch data.Type {
//		case "1":
//			strs = l.GetWords()
//		case "2":
//			strs = l.GetLetters()
//		case "3":
//			strs = l.GetNums()
//		}
//		responseData.Data[i] = strs
//
//		for j := 0; j < len(strs); j++ {
//			_, err = file.WriteString(strings.Join(strs[j], " ") + "\n")
//			if err != nil {
//				fmt.Println("Unable to write file:", err)
//				os.Exit(1)
//			}
//		}
//		_, _ = file.WriteString("\n")
//	}
//	_ = json.NewEncoder(w).Encode(responseData)
//	return
//}
//
//func getFingers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	const firstSeriesLen = 3
//	const maxSeriesLen = 10
//
//	var responseData DataOutput
//	responseData.Data = make([][][]string, maxSeriesLen-firstSeriesLen+1)
//	responseData.Result = "ok"
//
//	file, err := os.OpenFile("last_level.txt", os.O_APPEND|os.O_WRONLY, 0600)
//	if err != nil {
//		fmt.Println("Unable to open file:", err)
//		os.Exit(1)
//	}
//	defer file.Close()
//
//	var l research.LevelParam
//	for i := 0; i < maxSeriesLen-firstSeriesLen+1; i++ {
//		l.New(i + firstSeriesLen)
//		strs := l.GetFingers()
//		responseData.Data[i] = strs
//
//		for j := 0; j < len(strs); j++ {
//			_, err = file.WriteString(strings.Join(strs[j], " ") + "\n")
//			if err != nil {
//				fmt.Println("Unable to write file:", err)
//				os.Exit(1)
//			}
//		}
//		_, _ = file.WriteString("\n")
//	}
//	_ = json.NewEncoder(w).Encode(responseData)
//	return
//}
//
////func writeFile(d DataOutput) {
////
////	for i := 0; i < len(d.Data); i++ {
////
////	}
////}
//
//func handleRequest() {
//	//http.HandleFunc("/", homePage)
//	r := httprouter.New()
//	routes(r)
//}
//
//func main() {
//	handleRequest()
//}

var l int

func mapCheck3(m map[int][]int) bool {
	if m[0][0] == m[3][0] && m[0][1] == m[3][1] && m[0][2] == m[3][2] {
		return false
	}
	if m[0][0] == m[4][0] && m[0][1] == m[4][1] && m[0][2] == m[4][2] {
		return false
	}
	if m[1][0] == m[4][0] && m[1][1] == m[4][1] && m[1][2] == m[4][2] {
		return false
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if m[i][j] == m[i+2][j] {
				return false
			}
		}
		for j := 0; j < 2; j++ {
			if m[i][j] == m[i+2][0] && m[i][j+1] == m[i+2][1] ||
				m[i][j] == m[i+2][1] && m[i][j+1] == m[i+2][2] {
				return false
			}
		}
	}

	return true
}

type patterns struct {
}

func my(m map[int][]int, depth int) {
	if depth == 4 {
		if !mapCheck3(m) {
			return
		}
		fmt.Println("{")
		fmt.Printf("{%d, %d, %d, %d},\n", m[0][0], m[0][1], m[0][2], m[0][3])
		fmt.Printf("{%d, %d, %d, %d},\n", m[1][0], m[1][1], m[1][2], m[1][3])
		fmt.Printf("{%d, %d, %d, %d},\n", m[2][0], m[2][1], m[2][2], m[2][3])
		fmt.Printf("{%d, %d, %d, %d},\n", m[3][0], m[3][1], m[3][2], m[3][3])
		fmt.Printf("{%d, %d, %d, %d},\n", m[4][0], m[4][1], m[4][2], m[4][3])
		fmt.Println("},")
		return
	}
	lev := m[depth]
	depth++
	m[depth] = []int{lev[1], lev[0], lev[3], lev[2]}
	my(m, depth)
	m[depth] = []int{lev[1], lev[3], lev[0], lev[2]}
	my(m, depth)
	m[depth] = []int{lev[3], lev[2], lev[1], lev[0]}
	my(m, depth)
	m[depth] = []int{lev[3], lev[2], lev[0], lev[1]}
	my(m, depth)

	//fmt.Println(i, ")")
	//i++

}

func main() {
	m := make(map[int][]int, 5)
	m[0] = []int{1, 2, 3, 4}
	my(m, 0)

}
