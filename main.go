package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	go func() {
		err := updateLinks()
		if err != nil {
			log.Println(err)
		}
	}()

	srv := http.Server{
		Addr:    "127.0.0.1:2410",
		Handler: http.HandlerFunc(handler),
	}
	fmt.Print("http://", srv.Addr, "/\n")

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "" || r.URL.Path[0] != '/' {
		w.WriteHeader(400)
		return
	}

	key := r.URL.Path[1:]

	switch key {
	case "", "index.html":
		http.ServeFile(w, r, "html/index.html")
		return
	case "ok", "ok.txt":
		handlerOKWithString(w, r, "ok")
		return
	case "favicon.ico":
		w.WriteHeader(404)
		return
	case "ua", "useragent":
		handlerOKWithString(w, r, r.UserAgent())
		return
	// case "time":
	// 	handlerOK(w, r)
	// 	fmt.Fprint(w, "UnixMicro: ", time.Now().UnixMicro(), "\n")
	// 	return
	case "head", "header", "headers":
		handlerOK(w, r)
		r.Header.Write(w)
		return
	}

	// key 不能包含部分字符
	for _, v := range key {
		switch v {
		case '/', '.', '\\':
			handlerNotFound(w, r)
			return
		}
	}

	err := updateLinks()
	if err != nil {
		handlerError(w, r, 500, err.Error())
		return
	}

	for i := 0; i < 15; i++ {
		toLink, ok := links[key]
		if !ok || toLink == "" || toLink == "@" {
			handlerNotFound(w, r)
			return
		}

		if toLink[0] != '@' {
			handlerRedirect(w, r, 302, toLink)
			return
		}

		key = toLink[1:]
	}
	handlerError(w, r, 508, "Too many redirects")
}
