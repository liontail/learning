package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Respond from localhost:8000") //write string
}

func v1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Enter v1 Section") // write string
}

func edit(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Enter v1 edit Section") //write string
}

func main() {

	http.HandleFunc("/", hello)        //Request call func hello
	http.HandleFunc("/v1/", v1)        //Request path v1 call func v1
	http.HandleFunc("/v1/edit/", edit) //Request edit v1
	fmt.Println("Your server is running http://localhost:8000")
	//http.ListenAndServe(":8000", nil) //Running server
	s := &http.Server{
		Addr:           ":8080",
		Handler:        http.DefaultServeMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
