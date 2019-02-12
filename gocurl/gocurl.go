package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func get(link string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body) //Read all Respond from web
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents)) // Print all contents

}

/*func push(link string) {

}

func post(link string) {

}*/

func main() {
	method := flag.String("x", "GET", "HTTP METHOD. Defualt to GET")
	flag.Parse()

	link := flag.Arg(0)       //Get first Arg
	u, err := url.Parse(link) //Parse URL
	if err != nil {
		panic(err)
	}

	if *method == "GET" {
		get(u.String()) //Conver URL to String
	}

	//	if *method == "PUSH" {
	//		push(link)
	//	}
	//	if *method == "POST" {
	//		post(link)
	//	}
	//var link string
	//fmt.Printf("%s\n", *method)
	//fmt.Printf("%s\n", method3)

}
