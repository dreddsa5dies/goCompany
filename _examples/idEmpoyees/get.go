package main

import (
	"fmt"
	"os"

	gocompany "github.com/dreddsa5dies/goCompany"
)

func main() {
	resultEmp, err := gocompany.GetEmployees("7030")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultEmp)
	os.Exit(0)
}
