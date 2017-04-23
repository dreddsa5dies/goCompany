package main

import (
	"fmt"
	"os"

	gocompany "github.com/dreddsa5dies/goCompany"
)

func main() {
	// поиск по имени
	id := "7030"
	resultByID, err := gocompany.GetIDData(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultByID)
	os.Exit(0)
}
