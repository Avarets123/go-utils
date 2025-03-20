package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Arguments not passed!")
		return
	}

	file := arguments[1]

	path := os.Getenv("PATH")

	pathSplited := filepath.SplitList(path)

	for _, dir := range pathSplited {

		fullPath := filepath.Join(dir, file)

		fileInfo, err := os.Stat(fullPath)

		if err == nil {
			mode := fileInfo.Mode()

			if mode.IsRegular() {
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					return

				}
			}
		}

	}

}
