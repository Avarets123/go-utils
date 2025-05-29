package ch6

import (
	"bufio"
	"fmt"
	"os"
)

func WriteWithDifferentVariants() {

	f1, err := os.Create("/tmp/f1.txt")
	if err != nil {
		panic(err)
	}

	defer f1.Close()

	writeText := "ADD TEXT"

	n, err := fmt.Fprint(f1, writeText)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote f1 file bytes %d", n)

	f2, err := os.Create("/tmp/f2.txt")
	if err != nil {
		panic(err)
	}

	defer f2.Close()

	n, err = f2.WriteString(writeText)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote f2 file bytes %d", n)

	f3, err := os.OpenFile("/tmp/f3.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer f3.Close()

	bw := bufio.NewWriter(f3)

	n, err = bw.WriteString(writeText)
	bw.Flush()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote f3 file bytes %d", n)

}
