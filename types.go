// Package goCompany это Go (golang) пакет для использования ОГРН онлайн(https://ru.rus.company/) API
package goCompany

import (
	"net/url"
	"regexp"
)

var (
	// блоки для проверки по регулярному выражению
	regOgrnip, _ = regexp.Compile(`([0-9]){15}`)
	regOgrn, _   = regexp.Compile(`([0-9]){13}`)
	regInn, _    = regexp.Compile(`([0-9]){10,12}`)
	v            = url.Values{}
)

// CompanyInfo структура для базовой информации формата:
// {
// "id" : "7030",
// "name" : "АКЦИОНЕРНОЕ ОБЩЕСТВО "ИНСТИТУТ "СТРОЙПРОЕКТ"",
// "shortName" : "АО "ИНСТИТУТ "СТРОЙПРОЕКТ"",
// "ogrn" : "1027810258673",
// "ogrnDate" : "2002-11-12T00:00:00.000",
// "inn" : "7826688390",
// "kpp" : "781001001"
// }
type CompanyInfo struct {
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

// Employees - структура данных об сотрудниках компании
type Employees struct {
	ID          int `json:"id"`
	CompanyInfo `json:"company"`
	PersonOwner `json:"person"`
	Post        `json:"post"`
	PostName    string `json:"postName"`
	Phone       string `json:"phone"`
}

// Post - структура для описания должности в компании
type Post struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	FullName string `json:"fullName"`
}

// Founders - структура данных об учредителях компании
type Founders struct {
	ID          int `json:"id"`
	CompanyInfo `json:"company"`
	PersonOwner `json:"personOwner"`
	Price       float64 `json:"price"`
	OwnerRussia bool    `json:"ownerRussia"`
}

// PersonOwner - данные о владельцах
type PersonOwner struct {
	ID              int    `json:"id"`
	FirstName       string `json:"firstName"`
	MiddleName      string `json:"middleName"`
	SurName         string `json:"surName"`
	INN             string `json:"inn"`
	URL             string `json:"url"`
	FullName        string `json:"fullName"`
	FullNameWithInn string `json:"fullNameWithInn"`
}

// Person - данные о человеке по идентификатору {id}
type Person struct {
	ID              int    `json:"id"`
	FirstName       string `json:"firstName"`
	MiddleName      string `json:"middleName"`
	SurName         string `json:"surName"`
	URL             string `json:"url"`
	FullName        string `json:"fullName"`
	FullNameWithInn string `json:"fullNameWithInn"`
}

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
	MainOkved1        MainOkved `json:"mainOkved1"`
	Okved1            Okved     `json:"okved1"`
	MainOkved2        MainOkved `json:"mainOkved2"`
	Okved2            Okved     `json:"okved2"`
	PfrRegistration   `json:"pfrRegistration"`
	FssRegistration   `json:"fssRegistration"`
	Fns               `json:"fns"`
	ArrAssignee       []CoType `json:"assignee"`
	ArrPredecessor    []CoType `json:"predecessor"`
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
	Region           Reg    `json:"region"`
	Street           Reg    `json:"street"`
	House            string `json:"house"`
	Building         string `json:"building"`
	Flat             string `json:"flat"`
	PostalIndex      string `json:"postalIndex"`
	FullAddress      string `json:"fullAddress"`
	FullHouseAddress string `json:"fullHouseAddress"`
}

// Reg - Регион или Улица
type Reg struct {
	ID            int     `json:"id"`
	NAME          string  `json:"name"`
	AOID          string  `json:"aoid"`
	GUID          string  `json:"guid"`
	PostalCode    string  `json:"postalCode"`
	Level         int     `json:"level"`
	OKATO         string  `json:"okato"`
	RegionType    RegType `json:"type"`
	RegionCode    string  `json:"regionCode"`
	AutoCode      string  `json:"autoCode"`
	AreaCode      string  `json:"areaCode"`
	CityCode      string  `json:"cityCode"`
	CtarCode      string  `json:"ctarCode"`
	PlaceCode     string  `json:"placeCode"`
	StreetCode    string  `json:"streetCode"`
	ExtrCode      string  `json:"extrCode"`
	SextCode      string  `json:"sextCode"`
	KladrCode     string  `json:"kladrCode"`
	Live          bool    `json:"live"`
	TypeName      string  `json:"typeName"`
	TypeShortName string  `json:"typeShortName"`
	URL           string  `json:"url"`
	CompanyCount  int     `json:"companyCount"`
	FullName      string  `json:"fullName"`
}

// RegType - Тип региона или Улицы
type RegType struct {
	ID        int    `json:"id"`
	NAME      string `json:"name"`
	ShortName string `json:"shortName"`
	Code      string `json:"code"`
	Level     int    `json:"level"`
}

// MainOkved - Основной вид деятельности
type MainOkved struct {
	ID           int    `json:"id"`
	NAME         string `json:"name"`
	Code         string `json:"code"`
	Parent       `json:"parent"`
	Description  string `json:"description"`
	CompanyCount int    `json:"companyCount"`
	URL          string `json:"url"`
	FullName     string `json:"fullName"`
}

// Okved - сборник кодов, присвоенных видам деятельности компаний
type Okved []struct {
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
	OrgData          `json:"pfr"`
}

// FssRegistration - регистрация в Фонде социального страхования (ФСС)
type FssRegistration struct {
	RegistrationDate string `json:"registrationDate"`
	Number           string `json:"number"`
	OrgData          `json:"fss"`
}

// OrgData - данные по ПФР или ФСС
type OrgData struct {
	ID       int    `json:"id"`
	NAME     string `json:"name"`
	Code     string `json:"code"`
	FullName string `json:"fullName"`
}

// Fns - ФНС (Федеральная налоговая служба)
type Fns struct {
	ID         int    `json:"id"`
	NAME       string `json:"name"`
	Code       string `json:"code"`
	AddressFns string `json:"address"`
	FullName   string `json:"fullName"`
}

// CoType - предшественник, правопреемник
type CoType struct {
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

// IEData структура данных о индивидуальном предпринимателе
// пример по ссылке https://ru.rus.company/интеграция/?человек=7528374
type IEData struct {
	ID             int `json:"id"`
	PersonOwner    `json:"person"`
	LastUpdateDate string `json:"lastUpdateDate"`
	MainOkved      `json:"mainOkved1"`
	Okved          `json:"okved1"`
	CloseInfo      `json:"closeInfo"`
	Ogrn           string `json:"ogrn"`
	OgrnDate       string `json:"ogrnDate"`
	Citizenship    `json:"citizenship"`
	Fns            `json:"fns"`
	TypeIE         OrgData  `json:"type"`
	License        []string `json:"license"`
}

// Citizenship - гражданин РФ
type Citizenship struct {
	Russian bool `json:"russian"`
}

// Positions - данные о должностях по идентификатору {id}
type Positions struct {
	ID          int `json:"id"`
	CompanyInfo `json:"company"`
	PersonOwner `json:"person"`
	Post        `json:"post"`
	PostName    string `json:"postName"`
	Phone       string `json:"phone"`
}
