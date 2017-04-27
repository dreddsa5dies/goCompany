package main

import (
	"fmt"
	"os"

	gocompany "github.com/dreddsa5dies/goCompany"
)

func main() {
	resultPerson, err := gocompany.GetPerson("7030")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultPerson)
	os.Exit(0)
}
