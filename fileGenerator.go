package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nguyenthenguyen/docx"
	"io"
	"net/http"
	"os"
	"strconv"
	"traces/research"
)

func getLetters(arr []int) (str string) {
	m := map[int]string{
		1:  "д",
		2:  "к",
		3:  "л",
		4:  "с",
		5:  "з",
		6:  "б",
		7:  "м",
		8:  "ч",
		9:  "в",
		10: "р",
	}
	for i := 0; i < len(arr); i++ {
		str += m[arr[i]]
	}
	return str
}

func getNums(arr []int) (str string) {
	for _, v := range arr {
		str += strconv.Itoa(v)
	}
	return str
}

func createResearchFile() {
	r, err := docx.ReadDocxFile("research.docx")

	if err != nil {
		fmt.Println("Unable to create file:", err)
	}
	defer r.Close()

	docx1 := r.Editable()
	levels := research.Get()

	writer := func(arr [][][]int, f func([]int) string) {
		for i := 0; i < len(arr); i++ {
			var text string
			for j := 0; j < len(arr[i]); j++ {
				text = f(arr[i][j])
				err = docx1.Replace("REPL", text, 1)
				if err != nil {
					fmt.Println("Unable to replace in file:", err)
				}
			}
		}
	}
	writer(levels.Words, getLetters)
	writer(levels.Fingers, getNums)
	err = docx1.WriteToFile("tmp.docx")
	if err != nil {
		fmt.Println("Unable to replace in file:", err)
	}
}

func downloadFile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	createResearchFile()
	Filename := "tmp.docx"

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
