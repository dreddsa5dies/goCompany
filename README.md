[![Build Status](https://travis-ci.org/dreddsa5dies/goCompany.svg?branch=master)](https://travis-ci.org/dreddsa5dies/goCompany) [![Go Report Card](https://goreportcard.com/badge/github.com/dreddsa5dies/goCompany)](https://goreportcard.com/report/github.com/dreddsa5dies/goCompany) [![GORef](https://godoc.org/github.com/dreddsa5dies/goCompany?status.svg)](https://godoc.org/github.com/dreddsa5dies/goCompany) [![CircleCI Status](https://circleci.com/gh/dreddsa5dies/goCompany.svg?style=shield)](https://circleci.com/gh/dreddsa5dies/goCompany) ![License](https://img.shields.io/badge/License-MIT-blue.svg) [![codecov](https://codecov.io/gh/dreddsa5dies/goCompany/branch/master/graph/badge.svg)](https://codecov.io/gh/dreddsa5dies/goCompany)  

![IMAGE](img/goCompany.png)

### Go (golang) пакет для использования [ОГРН онлайн](https://ru.rus.company/) API

## Особенности
* Поиск компании по Наименованию
* Поиск компании по ОГРН
* Поиcк компании по ИНН
* Получение данных из базы по ID о компании
* Получение данных из базы по ID о учредителях
* Получение данных из базы по ID о сотрудниках
* Получение данных из базы по ID о зависимых организациях
* Получение данных из базы по ID о человеке
* Получение данных из базы по ID о должностях человека
* Получение данных из базы по ID о компаниях, в которых данный человек является учредителем
* Поиск ИП по ID
* Поиск ИП по ОГРНИП
* Поиcк ИП по ИНН  

## TODO
* Increase coverage with tests

## Установка
```bash
go get github.com/dreddsa5dies/goCompany
```

## Пример использования
Примеры использования смотрите в [примерах](https://github.com/dreddsa5dies/goCompany/tree/master/_examples)  
Документация по Client [![GORef](https://godoc.org/github.com/dreddsa5dies/goCompany?status.svg)](https://godoc.org/github.com/dreddsa5dies/goCompany)
```Go
package main

import (
	"fmt"
	"os"

	gocompany "github.com/dreddsa5dies/goCompany"
)

func main() {
	name := "Ласточка"
	result, err := gocompany.GetCompanyInfo(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)
	os.Exit(0)
}
```

## Лицензия
Проект под лицензией MIT. Прочитайте LICENSE файл.

## Вклад в проект
Добро пожаловать, следуйте следующим шагам:

- Форкните (fork) проект на github.com.
- Создайте новую ветку.
- Зафиксируйте (commit) изменения в Вашей ветке.
- Отправьте pull request.
