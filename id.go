// Package goCompany это Go (golang) пакет для использования ОГРН онлайн(https://ru.rus.company/) API
package goCompany

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// IDData структура для информации по ID компании
// пример по ссылке https://ru.rus.company/интеграция/компании/7030/
type IDData struct {
	ID                int    `json:"id"`
	OGRN              string `json:"ogrn"`
	NAME              string `json:"name"`
	SHORTNAME         string `json:"shortName"`
	OGRNDATE          string `json:"ogrnDate"`
	INN               string `json:"inn"`
	KPP               string `json:"kpp"`
	URL               string `json:"url"`
	Okopf             `json:"okopf"`
	LastUpdateDate    string `json:"lastUpdateDate"`
	Email             string `json:"email"`
	AuthorizedCapital `json:"authorizedCapital"`
}

// Okopf
type Okopf struct {
	ID       string `json:"id"`
	CODE     string `json:"code"`
	NAME     string `json:"name"`
	FULLNAME string `json:"fullname"`
	Parent   `json:"parent"`
}

// Parent
type Parent struct {
	ID string `json:"id"`
}

// AuthorizedCapital
type AuthorizedCapital struct {
	TypeCapital `json:"type"`
	Value       string `json:"value"`
}

// Address
type Address struct {
	Region
}

// Region
type Region struct {
}

// GetIDData возвращает массив IDData
func GetIDData(id string) ([]IDData, error) {
	url := baseURL + id + "/"

	// создание массива переменных для хранения ответа
	var companyJSON []IDData

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
