package ch6

import (
	"os"
	"path"
)

func DirSize(dirPath string) int64 {

	dirs, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	var size int64
	for _, d := range dirs {
		if d.IsDir() {
			size += DirSize(path.Join(dirPath, d.Name()))
		}

		finfo, err := d.Info()
		if err != nil {
			panic(err)
		}
		size += finfo.Size()

	}

	return size

}
