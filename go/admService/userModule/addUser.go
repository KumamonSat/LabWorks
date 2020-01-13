package userModule

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type reqUser struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status string `json:"status"`
	Group string `json:"group"`
}


func AddUser(w http.ResponseWriter, r *http.Request){
	var res = reqUser{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Print(err)
	}

	db, err := sql.Open("sqlite3", "../labworks.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into users (email, username, password, status, gr) values ($1, $2, $3, $4, $5)",
	res.Email, res.Username, res.Password, res.Status, res.Group)
	if err != nil{
		panic(err)
	}

	last, _ := result.LastInsertId()
	str := strconv.FormatInt(last, 10)

	message := fmt.Sprintf("LastInsertedID:"+ str)

 	w.Write([]byte(message))

	log.Printf("Added new user with id: " + str)
}