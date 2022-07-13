package main

import(
	//"fmt"
	"net/http"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request ){
	w.Write([]byte("<h1>Hello World</h1>"))
}

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
