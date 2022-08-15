package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",func(res http.ResponseWriter,req * http.Request){
		log.Println("get req")
		log.Println("get req2")
		log.Println("get req3")
		res.Write([]byte("hello\n"))
	})
	http.ListenAndServe(":10101",nil)
}
