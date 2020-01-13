package main

import (
	"./usersModule"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"mux"
	"net/http"
)




func main()  {
	r := mux.NewRouter()
	r.HandleFunc("/get/lessons", usersModule.GetLessons).Methods("POST")

	log.Print("Server is listening")
	http.ListenAndServe("127.0.0.1:8082", r)
}