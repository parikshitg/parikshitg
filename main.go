package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Port for running the UI")
	flag.Parse()

	var err error
	if os.Getenv("PORT") != "" {
		port, err = strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			panic(err)
		}
	}

	http.Handle("/", http.HandlerFunc(Home))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("static/images"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("static/fonts"))))
	http.Handle("/download/", http.StripPrefix("/download/", http.FileServer(http.Dir("static/download"))))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
