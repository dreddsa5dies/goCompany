package main

import (
	"fmt"
	"os"

	gocompany "github.com/dreddsa5dies/goCompany"
)

func main() {
	resultDepComp, err := gocompany.GetDepComp("7030")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultDepComp)
	os.Exit(0)
}
