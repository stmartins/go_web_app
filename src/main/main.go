package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "os"
	"strings"
	"sync"
)

type myHandler struct {
	mu sync.RWMutex
}

func (hd *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	fmt.Println(path)
	data, err := ioutil.ReadFile(string(path))

	hd.mu.Lock()

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
		// fmt.Println(contentType)

		w.Header().Add("Content-Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 Page " + http.StatusText(404)))
	}

	hd.mu.Unlock()
}

func main() {
	// log.Println("ENV: ", os.Environ())
	mx := http.NewServeMux()
	fmt.Println("listening 8080...")

	mx.Handle("/", new(myHandler))
	http.ListenAndServe(":8080", mx)

}
