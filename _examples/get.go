package main

import (
	"fmt"
	"os"

	gocompany "github.com/dreddsa5dies/goCompany"
)

func main() {
	name := "Ласточка"
	result, err := gocompany.GetByName(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
	os.Exit(0)
}
