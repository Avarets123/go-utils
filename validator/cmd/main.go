package main

import "validator/pkg/validator"

type Test2 struct {
	Id      string  `valid:"string,required"`
	Amount  float64 `valid:"float,required"`
	Comment string  `valid:"string,required"`
}

type Test struct {
	Str     string  `valid:"string"`
	Int     int     `valid:"int,required"`
	Fl      float64 `valid:"float"`
	Sll     []int   `valid:"int,required"`
	AddInfo Test2   `valid:"struct,required"`
}

func main() {

	valid := validator.New()

	err := valid.ValidStruct(Test{
		Str: "dsds",
		Fl:  12,
	})

	err.LogError()

}
