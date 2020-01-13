package lessonModule

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
	Group string `json:"group"`
	Lesson string `json:"lesson"`
	Teacher string `json:"teacher"`
}


func AddLesson(w http.ResponseWriter, r *http.Request){
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

	result, err := db.Exec("insert into lessons (gr, lesson, teacher) values ($1, $2, $3)",
		res.Group, res.Lesson, res.Teacher)
	if err != nil{
		panic(err)
	}

	last, _ := result.LastInsertId()
	str := strconv.FormatInt(last, 10)

	message := fmt.Sprintf("LastInsertedID:"+ str)

	w.Write([]byte(message))

	log.Printf("Added new user with id: " + str)
}
