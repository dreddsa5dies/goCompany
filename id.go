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
	MainOkved         `json:"mainOkved1"`
	ArrOkved          []Okved `json:"okved1"`
	PfrRegistration   `json:"pfrRegistration"`
	FssRegistration   `json:"fssRegistration"`
	Fns               `json:"fns"`
	ArrAssignee       []Assignee    `json:"assignee"`
	ArrPredecessor    []Predecessor `json:"predecessor"`
}

// Okopf - Общероссийский классификатор организационно-правовых форм
type Okopf struct {
	ID       int    `json:"id"`
	CODE     string `json:"code"`
	NAME     string `json:"name"`
	FULLNAME string `json:"fullname"`
	Parent   `json:"parent"`
}

// Parent - ID родителя
type Parent struct {
	ID int `json:"id"`
}

// AuthorizedCapital - Уставной капитал
type AuthorizedCapital struct {
	TypeCapital `json:"type"`
	Value       float64 `json:"value"`
}

// TypeCapital - Тип уставного капиталаы
type TypeCapital struct {
	ID   int    `json:"id"`
	NAME string `json:"name"`
}

// Address - Адрес
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

// Region - Регион
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

// RegionType - Тип региона
type RegionType struct {
	ID        int    `json:"id"`
	NAME      string `json:"name"`
	ShortName string `json:"shortName"`
	Code      string `json:"code"`
	Level     int    `json:"level"`
}

// Street - Улица
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

// StreetType - Тип улицы
type StreetType struct {
	ID        int    `json:"id"`
	NAME      string `json:"name"`
	ShortName string `json:"shortName"`
	Code      string `json:"code"`
	Level     int    `json:"level"`
}

// MainOkved - Основной вид деятельности
type MainOkved struct {
	ID       int    `json:"id"`
	NAME     string `json:"name"`
	Code     string `json:"code"`
	Parent   `json:"parent"`
	URL      string `json:"url"`
	FullName string `json:"fullName"`
}

// Okved - сборник кодов, присвоенных видам деятельности компаний
type Okved struct {
	ID          int    `json:"id"`
	NAME        string `json:"name"`
	Code        string `json:"code"`
	Parent      `json:"parent"`
	Description string `json:"description"`
	URL         string `json:"url"`
	FullName    string `json:"fullName"`
}

// PfrRegistration - регистрация в ПФР (Пенсионный фонд России)
type PfrRegistration struct {
	RegistrationDate string `json:"registrationDate"`
	Number           string `json:"number"`
	Pfr              `json:"pfr"`
}

// Pfr - данные по ПФР
type Pfr struct {
	ID       int    `json:"id"`
	NAME     string `json:"name"`
	Code     string `json:"code"`
	FullName string `json:"fullName"`
}

// FssRegistration - регистрация в Фонде социального страхования (ФСС)
type FssRegistration struct {
	RegistrationDate string `json:"registrationDate"`
	Number           string `json:"number"`
	Fss              `json:"fss"`
}

// Fss - ФСС
type Fss struct {
	ID       int    `json:"id"`
	NAME     string `json:"name"`
	Code     string `json:"code"`
	FullName string `json:"fullName"`
}

// Assignee - правопреемник
type Assignee struct {
	ID        int    `json:"id"`
	OGRN      string `json:"ogrn"`
	NAME      string `json:"name"`
	ShortName string `json:"shortName"`
	OgrnDate  string `json:"ogrnDate"`
	INN       string `json:"inn"`
	KPP       string `json:"kpp"`
	URL       string `json:"url"`
}

// Predecessor - предшественник
type Predecessor struct {
	ID        int    `json:"id"`
	OGRN      string `json:"ogrn"`
	NAME      string `json:"name"`
	ShortName string `json:"shortName"`
	OgrnDate  string `json:"ogrnDate"`
	INN       string `json:"inn"`
	KPP       string `json:"kpp"`
	CloseInfo `json:"closeInfo"`
	URL       string `json:"url"`
}

// CloseInfo - информация о закрытии
type CloseInfo struct {
	Date        string `json:"date"`
	CloseReason `json:"closeReason"`
}

// CloseReason - причина закрытия
type CloseReason struct {
	ID int `json:"id"`
}

// Fns - ФНС (Федеральная налоговая служба)
type Fns struct {
	ID         int    `json:"id"`
	NAME       string `json:"name"`
	Code       string `json:"code"`
	AddressFns string `json:"address"`
	FullName   string `json:"fullName"`
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
