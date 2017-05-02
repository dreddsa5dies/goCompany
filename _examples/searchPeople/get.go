package main

import (
	"fmt"
	"os"

	gocompany "github.com/dreddsa5dies/goCompany"
)

func main() {
	// поиск по ИНН (тип string)
	resultByInn, err := gocompany.GetSearchPerson("781010102204")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("ФИО %v\n", resultByInn[0].FullName)

	// поиск по ФИО
	resultByF, err := gocompany.GetSearchPerson("Журбин")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("ФИО %v\n", resultByF[0].FullName)

	os.Exit(0)
}
