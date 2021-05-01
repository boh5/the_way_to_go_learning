package main

import (
	"io"
	"log"
	"net/http"
)

const form = `<html><body><form action="#" method="post" name="bar">
		<input type="text" name="in"/>
		<input type="submit" value="Submit"/>
	</form></html></body>`

type HandleFunc func(w http.ResponseWriter, r *http.Request)

func main() {
	http.HandleFunc("/test1", logPanics(SimpleServer))
	http.HandleFunc("/test2", logPanics(FormServer))
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}

func SimpleServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<h1>Hello, world</h1>")
}

func FormServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch req.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, req.FormValue("in"))
	}
}

func logPanics(handleFunc HandleFunc) HandleFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[%v] caught panic %v", req.RemoteAddr, r)
			}
		}()
		handleFunc(w, req)
	}
}
