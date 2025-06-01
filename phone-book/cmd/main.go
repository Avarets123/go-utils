package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"phone-book/internal/entry"
	"phone-book/pkg/ch7"
	"regexp"
	"runtime"
	"strings"
)

func isName(str string) bool {
	re := regexp.MustCompile("^[A-Z][a-z]*$")
	return re.Match([]byte(str))
}

func isInt(str string) bool {
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match([]byte(str))
}

func isCard(card string) bool {
	return regexp.MustCompile(`^[0-9]{10,18}$`).Match([]byte(card))
}
func isCardExire(card string) bool {
	return regexp.MustCompile(`^[0-9]{2}/[0-9]{4}$`).Match([]byte(card))
}

func isValidHolder(card string) bool {
	return regexp.MustCompile(`^[A-Z|a-z|А-Я|а-я]+$`).Match([]byte(card))
}

func systemLogs() {
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.Compiler)
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.Version())
	fmt.Println(runtime.NumCPU())

	fmt.Println(os.Hostname())

}

func isValidInput(input string) bool {
	inpSlic := strings.Split(input, ",")

	name := inpSlic[0]
	surname := inpSlic[1]
	phone := inpSlic[2]

	if !isName(name) {
		return false
	}

	if !isName(surname) {
		return false
	}

	return isInt(phone)

}

func main() {

	// systemLogs()

	ch7.UseCtx()

	return

	args := os.Args

	exeFile := path.Base(args[0])

	if len(args) == 1 {
		fmt.Printf("Usage: %s search|list|inser|delete <args> \n", exeFile)
		os.Exit(0)
	}

	phoneBooks := entry.InitPhoneBook()

	switch args[1] {

	case "search":
		{
			if len(args) == 2 {
				fmt.Printf("Usage: search <phone> \n")
				os.Exit(0)
			}

			book := phoneBooks.Search(args[2])
			if book == nil {
				fmt.Println("Record in phone-book not found!")
				return
			}

			b, err := json.MarshalIndent(book, "", "\t")
			if err != nil {
				panic(err)
			}

			fmt.Println(string(b))
			return

		}

	case "list":
		{
			// for _, v := range phoneBooks.List() {
			// 	fmt.Printf("%+v \n", v)
			// }

			b, err := json.MarshalIndent(phoneBooks.List(), "", "\t")
			if err != nil {
				panic(err)
			}

			fmt.Println(string(b))

			return

		}
	case "idx-list":
		{
			// fmt.Println(phoneBooks.IdxList())

			b, err := json.MarshalIndent(phoneBooks.IdxList(), "", "\t")
			if err != nil {
				panic(err)
			}

			fmt.Println(string(b))

			return
		}

	case "insert":
		{
			if len(args) != 5 {
				fmt.Println("Usage: insert <name> <surname> <phone>")
				return
			}

			err := phoneBooks.Insert(args[2], args[3], args[4])
			fmt.Println(err)

			return
		}
	case "delete":
		{

			if len(args) != 3 {
				fmt.Println("Usage: delete <phone>")
				return
			}

			err := phoneBooks.Delete(args[2])
			fmt.Println(err)
			return

		}

	default:
		{
			fmt.Println("Passed invalid command")

		}

	}

}
