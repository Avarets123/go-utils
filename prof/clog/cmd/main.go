package main

import (
	"log"
	"os"
	"path"
)

func main() {

	cdir, _ := os.Getwd()
	LOGFILE := path.Join(os.TempDir(), "clogger.log")
	file, err := os.OpenFile(LOGFILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	flags := log.Lshortfile | log.Ldate
	clog := log.New(file, "[clogger]: ", flags)

	clog.Println("Logger successfully inited!")
	clog.Printf("App runned on path: %s \n", cdir)
	clog.Println("New logg")

}
