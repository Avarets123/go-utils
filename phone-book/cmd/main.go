package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name, Surname, Phone string
}

type PhoneBooks struct {
	Entries *[]Entry
}

func main() {

	args := os.Args

	exeFile := path.Base(args[0])

	if len(args) == 1 {
		fmt.Printf("Usage: %s search|list <args> \n", exeFile)
		os.Exit(0)
	}

	phoneBooks := initPhoneBooks()

	switch args[1] {

	case "search":
		{
			if len(args) == 2 {
				fmt.Printf("Usage: %s search <surname> \n", exeFile)
				os.Exit(0)
			}

			if phBook := phoneBooks.search(args[2]); phBook != nil {
				fmt.Println(*phBook)
				return
			}

			fmt.Println("Record in phone-book not found!")
			return

		}

	case "list":
		{
			for _, v := range *phoneBooks.list() {
				fmt.Println(v)
			}

		}

	default:
		{
			fmt.Println("Passed invalid command")

		}

	}

}

func initPhoneBooks() *PhoneBooks {
	entries := &[]Entry{
		{
			Name:    "Md",
			Surname: "Magomedov",
			Phone:   "+7988",
		},
		{
			Name:    "Os",
			Surname: "Gasanov",
			Phone:   "+7999",
		},
		{
			Name:    "Ah",
			Surname: "Abdulaev",
			Phone:   "+7928",
		},
	}

	return &PhoneBooks{
		Entries: entries,
	}

}

func (pb *PhoneBooks) list() *[]Entry {
	return pb.Entries
}

func (pb *PhoneBooks) search(surname string) *Entry {

	for _, v := range *pb.Entries {

		if v.Surname == surname {
			return &v
		}

	}

	return nil

}
