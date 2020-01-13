package main

import "log"

func main() {
	//auth := true
	//status := "admin"

	auth("admin", true)
}

func auth(status string, auth bool) {
	if status == "admin" && auth == True {
		log.Print(true)
	} else {
		log.Print(false)
	}
}
