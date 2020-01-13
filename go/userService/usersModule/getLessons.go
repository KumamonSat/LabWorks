package usersModule

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
)

type reqLesson struct {
	Group string `json:"group"`
}

type respLesson struct{
	Id string `json:"id"`
	Group string `json:"group"`
	Lesson string `json:"lesson"`
	Teacher string `json:"teacher"`
}

func GetLessons(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {

		var res= reqLesson{}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
		}

		log.Print(res.Group)

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

		rows, err := db.Query("select * from lessons where gr = $1", res.Group)
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		mess := []respLesson{}

		for rows.Next() {
			m := respLesson{}
			err := rows.Scan(&m.Id, &m.Group, &m.Lesson, &m.Teacher)
			if err != nil {
				log.Print(err)
				continue
			}
			mess = append(mess, m)
		}

		b, err := json.Marshal(mess)
		log.Print(b)
		if err != nil {
			log.Print(err)
			return
		}
		w.Write(b)
	}
}
