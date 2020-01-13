package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/login", AuthCheck).Methods("POST")

	log.Print("Server is listening")
	err := http.ListenAndServe("127.0.0.1:8080", r)
	if err != nil{
		log.Print(err)
	}
}

type reqUser struct {
	Email string `json:"email"`
	//Password string `json:"password"`
}

type respUser struct {
	Id int `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status string `json:"status"`
	Group string `json:"group"`
}

func AuthCheck(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {

		var res = reqUser{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
		}

		err = json.Unmarshal(body, &res)
		if err != nil {
			log.Print(err)
		}
		//log.Print(res.Password)


		db, err := sql.Open("sqlite3", "../labworks.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		rows, err := db.Query("select * from users where email = $1", res.Email)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		mess := []respUser{}

		for rows.Next(){
			m := respUser{}
			err := rows.Scan(&m.Id, &m.Email, &m.Username, &m.Password, &m.Status, &m.Group)
			if err != nil{
				log.Print(err)
				continue
			}
			mess = append(mess, m)
		}


		if len(mess) == 0 {
			log.Print("я насрал")

			m := respUser {Id:0,Email:"null",Username:"null",Password:"null",Status:"null",Group:"null"}
			mess := append(mess,m)

			b, err := json.Marshal(mess)
			if err != nil {
				log.Print(err)
				return
			}
			w.Write(b)
		} else if len(mess) >= 1{
			b, err := json.Marshal(mess)
			if err != nil {
				log.Print(err)
				return
			}
			w.Write(b)
		}
	}
}