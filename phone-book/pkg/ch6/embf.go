package ch6

import (
	_ "embed"
	"fmt"
)

//go:embed embf.go
var emb string

func PrintEmb() {

	fmt.Println(emb)

}
