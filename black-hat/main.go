package main

import (
	"black-hat/metasploit"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}

	host := os.Getenv("MSF_HOST")
	pass := os.Getenv("MSF_PASS")
	user := "msf"

	mts, err := metasploit.New(host, user, pass)
	if err != nil {
		log.Fatalln(err)
	}

	sessions, err := mts.SessionList()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(sessions)

}
