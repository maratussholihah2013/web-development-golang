package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", pemandangan)
	http.Handle("/mountains/", http.StripPrefix("/mountains", http.FileServer(http.Dir("./assets"))))
	//OR
	//http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))

	http.ListenAndServe(":8080", nil)
}

func pemandangan(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/mountains/1.jpg">`)
	//OR
	//io.WriteString(w, `<img src="/assets/1.jpg">`)
	io.WriteString(w, `<img src="/mountains/2.jpg">`)
	io.WriteString(w, `<img src="/mountains/3.jpg">`)
}
