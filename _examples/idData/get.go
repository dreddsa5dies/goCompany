package main

import (
	"fmt"
	"os"

	gocompany "github.com/dreddsa5dies/goCompany"
)

func main() {
	resultByID, err := gocompany.GetIDData("7030")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultByID)
	os.Exit(0)
}
