// Package goCompany это Go (golang) пакет для использования ОГРН онлайн(https://ru.rus.company/) API
package goCompany

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	baseURL = "https://ru.rus.company/"
)

// SaveData структура для базовой информации формата:
// {
// "id" : "7030",
// "name" : "АКЦИОНЕРНОЕ ОБЩЕСТВО "ИНСТИТУТ "СТРОЙПРОЕКТ"",
// "shortName" : "АО "ИНСТИТУТ "СТРОЙПРОЕКТ"",
// "ogrn" : "1027810258673",
// "ogrnDate" : "2002-11-12T00:00:00.000",
// "inn" : "7826688390",
// "kpp" : "781001001"
// }
type SaveData struct {
	ID        int    `json:"id"`
	OGRN      string `json:"ogrn"`
	NAME      string `json:"name"`
	SHORTNAME string `json:"shortName"`
	OGRNDATE  string `json:"ogrnDate"`
	INN       string `json:"inn"`
	KPP       string `json:"kpp"`
	URL       string `json:"url"`
}

// GetByName возвращает массив SaveData
func GetByName(nameCompany string) ([]SaveData, error) {
	url := baseURL + "/интеграция/компании/?наименование=" + nameCompany + "&стр=5"

	// создание массива переменных для хранения ответа
	var companyJSON []SaveData

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

// GetByOgrn возвращает массив SaveData
func GetByOgrn(ogrn string) ([]SaveData, error) {
	url := baseURL + "/интеграция/компании/?огрн=" + ogrn

	var companyJSON []SaveData

	resp, err := http.Get(url)
	if err != nil {
		return companyJSON, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &companyJSON)
	if err != nil {
		return companyJSON, err
	}

	return companyJSON, nil
}

// GetByInn возвращает массив SaveData
func GetByInn(inn string) ([]SaveData, error) {
	url := baseURL + "/интеграция/компании/?инн=" + inn

	var companyJSON []SaveData

	resp, err := http.Get(url)
	if err != nil {
		return companyJSON, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &companyJSON)
	if err != nil {
		return companyJSON, err
	}

	return companyJSON, nil
}
