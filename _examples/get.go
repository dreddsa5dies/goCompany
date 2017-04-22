package main

import (
	"fmt"
	"os"

	gocompany "github.com/dreddsa5dies/goCompany"
)

func main() {
	// поиск по имени
	name := "Ласточка"
	resultByName, err := gocompany.GetByName(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultByName[0].INN)

	// поиск по ИНН (тип string)
	inn := "6820002944"
	resultByInn, err := gocompany.GetByInn(inn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultByInn[0].OGRN)

	// поиск по ОГРН (тип string)
	ogrn := "1066820016416"
	resultByOgrn, err := gocompany.GetByOgrn(ogrn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultByOgrn[0].NAME)
	os.Exit(0)
}
