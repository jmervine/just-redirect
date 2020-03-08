package main

import (
	"log"
	"net/http"
	"os"
)

var target, port, bind string

func config() {
    var ok bool
	if target, ok = os.LookupEnv("REDIRECT_TARGET"); !ok {
		log.Fatal("REDIRECT_TARGET required")
	}

	if port, ok = os.LookupEnv("PORT"); !ok {
		port = "3000"
	}

	if bind, ok = os.LookupEnv("BIND"); !ok {
        bind = "0.0.0.0"
    }
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, target, 302)
}

func main() {
    log.SetFlags(0)
    config()

	http.HandleFunc("/", redirect)
	listen := bind + ":" + port
	print("listening on: " + listen + "\n")
	log.Fatal(http.ListenAndServe(listen, nil))
}
