package ch6

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func ReadLineByLine(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(file)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			fmt.Println(err)
			break
		}

		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(line)

	}

}

func ReadWordByWord(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		reg := regexp.MustCompile(`[^\\|]+`)

		words := reg.FindAllString(line, -1)

		for _, word := range words {
			fmt.Println(word)
		}

	}

}

func BufferedReadFile(filepath string, readSize int) {

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, readSize)
	r := bufio.NewReader(file)

	for {
		n, err := r.Read(buffer)
		if err == io.EOF {
			fmt.Println(err)
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}
		buffer = buffer[:n]
		fmt.Printf("Wrote count %d \n", n)
		fmt.Println(string(buffer))
	}

}
