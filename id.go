// Package goCompany это Go (golang) пакет для использования ОГРН онлайн(https://ru.rus.company/) API
package goCompany

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// IDData структура данных из ЕГРЮЛ о компании с идентификатором {id}
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
	Address           `json:"address"`
	MainOkved1        `json:"mainOkved1"`
}

// Okopf
type Okopf struct {
	ID       int    `json:"id"`
	CODE     string `json:"code"`
	NAME     string `json:"name"`
	FULLNAME string `json:"fullname"`
	Parent   `json:"parent"`
}

// Parent
type Parent struct {
	ID int `json:"id"`
}

// AuthorizedCapital
type AuthorizedCapital struct {
	TypeCapital `json:"type"`
	Value       float64 `json:"value"`
}

// TypeCapital
type TypeCapital struct {
	ID   int    `json:"id"`
	NAME string `json:"name"`
}

// Address
type Address struct {
	Region           `json:"region"`
	Street           `json:"street"`
	House            string `json:"house"`
	Building         string `json:"building"`
	Flat             string `json:"flat"`
	PostalIndex      string `json:"postalIndex"`
	FullAddress      string `json:"fullAddress"`
	FullHouseAddress string `json:"fullHouseAddress"`
}

// Region
type Region struct {
	ID            int    `json:"id"`
	NAME          string `json:"name"`
	AOID          string `json:"aoid"`
	GUID          string `json:"guid"`
	PostalCode    string `json:"postalCode"`
	Level         int    `json:"level"`
	OKATO         string `json:"okato"`
	RegionType    `json:"type"`
	RegionCode    string `json:"regionCode"`
	AutoCode      string `json:"autoCode"`
	AreaCode      string `json:"areaCode"`
	CityCode      string `json:"cityCode"`
	CtarCode      string `json:"ctarCode"`
	PlaceCode     string `json:"placeCode"`
	StreetCode    string `json:"streetCode"`
	ExtrCode      string `json:"extrCode"`
	SextCode      string `json:"sextCode"`
	KladrCode     string `json:"kladrCode"`
	Live          bool   `json:"live"`
	TypeName      string `json:"typeName"`
	TypeShortName string `json:"typeShortName"`
	URL           string `json:"url"`
	CompanyCount  int    `json:"companyCount"`
	FullName      string `json:"fullName"`
}

// RegionType
type RegionType struct {
	ID        int    `json:"id"`
	NAME      string `json:"name"`
	ShortName string `json:"shortName"`
	Code      string `json:"code"`
	Level     int    `json:"level"`
}

// Street
type Street struct {
	ID            int    `json:"id"`
	NAME          string `json:"name"`
	AOID          string `json:"aoid"`
	GUID          string `json:"guid"`
	Level         int    `json:"level"`
	StreetType    `json:"type"`
	RegionCode    string `json:"regionCode"`
	AutoCode      string `json:"autoCode"`
	AreaCode      string `json:"areaCode"`
	CityCode      string `json:"cityCode"`
	CtarCode      string `json:"ctarCode"`
	PlaceCode     string `json:"placeCode"`
	StreetCode    string `json:"streetCode"`
	ExtrCode      string `json:"extrCode"`
	SextCode      string `json:"sextCode"`
	KladrCode     string `json:"kladrCode"`
	Live          bool   `json:"live"`
	TypeName      string `json:"typeName"`
	TypeShortName string `json:"typeShortName"`
	URL           string `json:"url"`
	CompanyCount  int    `json:"companyCount"`
	FullName      string `json:"fullName"`
}

// StreetType
type StreetType struct {
	ID        int    `json:"id"`
	NAME      string `json:"name"`
	ShortName string `json:"shortName"`
	Code      string `json:"code"`
	Level     int    `json:"level"`
}

// MainOkved1
type MainOkved1 struct {
	ID           int    `json:"id"`
	NAME         string `json:"name"`
	Code         string `json:"code"`
	Parent       `json:"parent"`
	CompanyCount int      `json:"companyCount"`
	URL          string   `json:"url"`
	FullName     string   `json:"fullName"`
	Okved1       []Okved1 `json:"okved1"`
}

// Okved1
type Okved1 struct {
	ID           int    `json:"id"`
	NAME         string `json:"name"`
	Code         string `json:"code"`
	Parent       `json:"parent"`
	CompanyCount int    `json:"companyCount"`
	URL          string `json:"url"`
	FullName     string `json:"fullName"`
}

// GetIDData возвращает IDData
func GetIDData(id string) (IDData, error) {
	url := baseURL + id + "/"

	// создание массива переменных для хранения ответа
	var companyJSON IDData

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
