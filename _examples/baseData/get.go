package main

import (
	"fmt"
	"os"

	gocompany "github.com/dreddsa5dies/goCompany"
)

func main() {
	// поиск по имени
	name := "Ласточка"
	resultByName, err := gocompany.GetCompanyInfo(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultByName[0].INN)

	// поиск по ИНН (тип string)
	inn := "6820002944"
	resultByInn, err := gocompany.GetCompanyInfo(inn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultByInn[0].OGRN)

	// поиск по ОГРН (тип string)
	ogrn := "1066820016416"
	resultByOgrn, err := gocompany.GetCompanyInfo(ogrn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultByOgrn[0].NAME)

	// при отсутствии информации возвращается пустой массив
	errGet := "непонятно13212"
	resultError, err := gocompany.GetCompanyInfo(errGet)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultError)
	os.Exit(0)
}
