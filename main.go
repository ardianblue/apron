package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
)

func redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target,
		http.StatusPermanentRedirect)
}

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		userAgent := r.Header.Get("User-Agent")
		fmt.Println( "[+]",currentTime.Format("01-02-2006 15:04:05"), "URL Route:", r.URL, "User-Agent:", userAgent)
		if userAgent == "h4ck3r"{
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("HTTP 403 Forbidden"))
		} else {
			fmt.Fprintf(w, "Hello World " + currentTime.Format("01-02-2006 15:04:05"))
		}

	})

	go http.ListenAndServe(":80", http.HandlerFunc(redirect))
	fmt.Println("Server running on: http://localhost:80")
	panic(http.ListenAndServeTLS(":443", "localhost.crt", "localhost.key", nil))
}
