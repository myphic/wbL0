package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type TemplateData struct {
	Output string
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	ts, err := template.ParseFiles("./ui/html/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	uid := r.FormValue("search")

	log.Printf("Order from cache: %s", uid)
	find, err := FindOrderInCache(uid)
	if err != nil {
		err = ts.Execute(w, "Order not found")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	} else {
		ord, err := json.Marshal(find)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		Template := TemplateData{string(ord)}
		err = ts.Execute(w, Template)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}

}

func CreateServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
