// Package goCompany это Go (golang) пакет для использования ОГРН онлайн(https://ru.rus.company/) API
package goCompany

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	baseURL = "https://ru.rus.company/интеграция/компании/"
)

// BaseData структура для базовой информации формата:
// {
// "id" : "7030",
// "name" : "АКЦИОНЕРНОЕ ОБЩЕСТВО "ИНСТИТУТ "СТРОЙПРОЕКТ"",
// "shortName" : "АО "ИНСТИТУТ "СТРОЙПРОЕКТ"",
// "ogrn" : "1027810258673",
// "ogrnDate" : "2002-11-12T00:00:00.000",
// "inn" : "7826688390",
// "kpp" : "781001001"
// }
type BaseData struct {
	ID        int    `json:"id"`
	OGRN      string `json:"ogrn"`
	NAME      string `json:"name"`
	SHORTNAME string `json:"shortName"`
	OGRNDATE  string `json:"ogrnDate"`
	INN       string `json:"inn"`
	KPP       string `json:"kpp"`
	CloseInfo `json:"closeInfo"`
	URL       string `json:"url"`
}

// GetBaseData возвращает массив BaseData
func GetBaseData(data string) ([]BaseData, error) {
	var url string

	// блоки для проверки по регулярному выражению
	regOgrn, _ := regexp.Compile(`([0-9]){13}`)
	regInn, _ := regexp.Compile(`([0-9]){10,12}`)

	switch {
	//если это ОГРН
	case regOgrn.MatchString(data):
		url = baseURL + "?огрн=" + data
	// если это ИНН
	case regInn.MatchString(data):
		url = baseURL + "?инн=" + data
	// или название
	default:
		url = baseURL + "?наименование=" + data + "&стр=5"
	}

	// создание массива переменных для хранения ответа
	var companyJSON []BaseData

	// обращение к API
	resp, err := http.Get(url)
	if err != nil {
		return companyJSON, err
	}
	defer resp.Body.Close()

	// запись ответа в переменную
	body, err := ioutil.ReadAll(resp.Body)

	// декодирование []byte в интерфейс
	err = json.Unmarshal(body, &companyJSON)
	if err != nil {
		return companyJSON, err
	}

	return companyJSON, nil
}
