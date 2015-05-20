package main

import(
	"net/http"
	"log"
	"github.com/webapi/handlers"
)

func main(){
	http.HandleFunc("/files", handlers.FilesHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}
