package main

import (
	"fmt"
	"os"

	"strconv"

	gocompany "github.com/dreddsa5dies/goCompany"
)

func main() {
	// поиск по ID (число)
	id := 7528374
	resultByID, err := gocompany.GetIndivEntrep(strconv.Itoa(id))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Полное наименование ОКВЕД %v\n", resultByID[0].Okved[0].FullName)

	// поиск по ИНН (тип string)
	resultByInn, err := gocompany.GetIndivEntrep(resultByID[0].INN)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Отчество %v\n", resultByInn[0].MiddleName)

	// поиск по ОГРНИП (тип string)
	ogrnip := "314272211800010"
	resultByOgrnip, err := gocompany.GetIndivEntrep(ogrnip)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("ОГРН %v\n", resultByOgrnip[0].Ogrn)

	// при отсутствии информации возвращается пустой массив или ошибку
	errGet := "непонятно13212"
	resultError, err := gocompany.GetIndivEntrep(errGet)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(resultError)

	os.Exit(0)
}
