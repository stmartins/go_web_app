package main

import (
	"log"
	"net/http"
	"time"
)

func timeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("the time is " + tm))
	}
	return http.HandlerFunc(fn)
}

func main() {
	mux := http.NewServeMux()

	rh := http.RedirectHandler("http://intra.42.fr", 302)
	mux.Handle("/42", rh)

	th := timeHandler(time.RFC822)
	mux.Handle("/time", th)

	th = timeHandler(time.RFC1123)
	mux.Handle("/time/rfc1123", th)

	th = timeHandler(time.RFC3339)
	mux.Handle("/time/rfc3339", th)

	log.Println("listening on port 8080...")
	http.ListenAndServe(":8080", mux)
}
