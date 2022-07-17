package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "parse form error:%v", err)
	}

	fmt.Fprintf(w, "Post request successfull\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "name you entered is : %s\n",name)
	fmt.Fprintf(w, "address you entered is: %s\n",address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 3000\n")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}

}
