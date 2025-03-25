package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var path string = os.Getenv("PATH")

func main() {

	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Arguments not passed!")
		return
	}

	files := arguments[1:]

	filesExistsMap := map[string][]string{}

	pathSplited := filepath.SplitList(path)

	for _, dir := range pathSplited {

		for _, file := range files {
			fPath := execIsExists(dir, file)
			if fPath == "" {
				continue
			}

			filesValue := filesExistsMap[file]
			filesValue = append(filesValue, fPath)
			filesExistsMap[file] = filesValue
			continue
		}

	}

	fmt.Println(filesExistsMap)

}

func execIsExists(dir, file string) string {
	fullPath := filepath.Join(dir, file)

	fileInfo, err := os.Stat(fullPath)

	if err == nil {
		mode := fileInfo.Mode()

		if mode.IsRegular() {
			if mode&0111 != 0 {
				return fullPath

			}
		}
	}

	return ""

}
