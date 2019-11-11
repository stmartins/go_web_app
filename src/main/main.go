package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	// "os"
	"strings"
	"sync"
)

type myHandler struct {
	mu sync.RWMutex
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func (hd *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// fmt.Println(r)

	path := r.URL.Path[1:]
	fmt.Println(path)
	data, err := ioutil.ReadFile(string(path))

	hd.mu.Lock()
	defer hd.mu.Unlock()

	if err == nil {
		var contentType string

		if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else {
			contentType = "text/plain"
		}

		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 Page " + http.StatusText(404)))
	}

}

// func idx(w http.ResponseWriter, r *http.Request) {
// 	err := tpl.ExecuteTemplate(w, "carroussel.html", nil)
// 	path := r.URL.Path[1:]
// 	fmt.Println(path)
//
// 	if err != nil {
// 		fmt.Println("error la")
// 		http.Error(w, "internal server error ", http.StatusInternalServerError)
// 	}
// }

func test(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "test.html", nil)

	if err != nil {
		fmt.Println("error")
		http.Error(w, "internal server error ", http.StatusInternalServerError)
	}
}

func main() {

	mx := http.NewServeMux()

	fmt.Println("listening 8080...")

	mx.Handle("/", new(myHandler))

	// mx.HandleFunc("/", idx)

	// rh := http.RedirectHandler("http://localhost:8080/templates/test.html", 307)
	mx.HandleFunc("/test", test)

	http.ListenAndServe(":8080", mx)

}
