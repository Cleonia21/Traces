package controller

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"path/filepath"
	"traces/levels"
)

func Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var l levels.Level
	l.New(10)
	data := l.GetWords()
	home := filepath.Join("templates", "html", "homePage.html")

	tmpl, err := template.ParseFiles(home)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = tmpl.ExecuteTemplate(w, "home", data) //data - передаваемый объект в шаблон
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}
