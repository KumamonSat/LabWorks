package main

import (
	"./lessonModule"
	"./userModule"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/user/add", userModule.AddUser).Methods("GET")
	//r.HandleFunc("/user/get", userModule.GetUser).Methods("GET")
	//r.HandleFunc("/user/del", userModule.UserDelete).Methods("GET")
	//r.HandleFunc("/user/upd", userModule.UserUpdate).Methods("GET")
	//
	//r.HandleFunc("/group/add", addGroup).Methods("GET")
	//r.HandleFunc("/group/get", getGroup).Methods("GET")
	//r.HandleFunc("/group/del", delGroup).Methods("GET")
	//r.HandleFunc("/group/upd", updGroup).Methods("GET")
	//
	r.HandleFunc("/lesson/add", lessonModule.AddLesson).Methods("GET")
	//r.HandleFunc("/lesson/get", getLesson).Methods("GET")
	//r.HandleFunc("/lesson/del", delLesson).Methods("GET")
	//r.HandleFunc("/lesson/upd", updLesson).Methods("GET")

	log.Print("Server was started")
	err := http.ListenAndServe("127.0.0.1:8081", r)
	if err != nil{
		log.Print(err)
	}
}


