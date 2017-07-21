package main

import (
	"log"
	"net/http"
	"os"
)

var target, port, bind string

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, target, 302)
}

func main() {
	target = os.Getenv("REDIRECT_TARGET")
	if target == "" {
		panic("REDIRECT_TARGET required")
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	bind = os.Getenv("BIND")

	http.HandleFunc("/", redirect)
	listen := bind + ":" + port
	print("listening on: " + listen + "\n")
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
