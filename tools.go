package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func handlerOK(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, 200, r.URL.RequestURI())
	w.WriteHeader(200)
}

func handlerWriteString(w http.ResponseWriter, body string) {
	io.WriteString(w, body)
	if body[len(body)-1] != '\n' {
		w.Write([]byte{'\n'})
	}
}

func handlerOKWithString(w http.ResponseWriter, r *http.Request, body string) {
	handlerOK(w, r)
	handlerWriteString(w, body)
}

func handlerRedirect(w http.ResponseWriter, r *http.Request, status int, location string) {
	log.Println(r.RemoteAddr, status, r.URL.RequestURI(), location)
	w.Header().Set("Location", location)
	w.WriteHeader(status)
}

func handlerNotFound(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, 404, r.URL.RequestURI())
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(404)
	b, err := os.ReadFile("html/404.html")
	if err == nil {
		w.Write(b)
	}
}

func handlerError(w http.ResponseWriter, r *http.Request, status int, err string) {
	log.Println(r.RemoteAddr, status, r.URL.RequestURI(), err)
	w.WriteHeader(status)
	handlerWriteString(w, err)
}
