package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Access-Control-Allow-Headers","*")
		w.Header().Add("Access-Control-Allow-Origin","*")
		if r.Method == "GET" {
			fmt.Fprintf(w, "<p>GET</p>")
		}else if r.Method == "POST" {
			fmt.Fprintf(w, "<p>POST</p>")
		}else {
			fmt.Fprintf(w, fmt.Sprintf("<p>%s - %v</p>", r.Method, time.Now()))
	 }
	})
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Access-Control-Allow-Headers","*")
		w.Header().Add("Access-Control-Allow-Origin","*")
		if r.Method == "GET"{
			fmt.Fprintf(w,fmt.Sprintf("<p>%s - %v</p>", r.Method, time.Now()))
		}else {
			fmt.Fprintf(w, fmt.Sprintf("<p>%s - %v</p>", r.Method, time.Now()))
	 }
	})
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
	}
}
