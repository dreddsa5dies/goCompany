package goCompany

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func init() {
	v.Set("baseURL", "https://ru.rus.company/интеграция")

	v.Add("apiURL", "/компании/")
	v.Add("apiURL", "/люди/")
	v.Add("apiURL", "/ип/")
	v.Add("apiURL", "/сотрудники/")
	v.Add("apiURL", "/компании/")
	v.Add("apiURL", "/зависимые/")
	v.Add("apiURL", "/должности/")
	v.Add("apiURL", "/учредители/")

	v.Add("apiKEY", "?огрн=")
	v.Add("apiKEY", "?инн=")
	v.Add("apiKEY", "?наименование=")
	v.Add("apiKEY", "?огрнип=")
	v.Add("apiKEY", "?человек=")
	v.Add("apiKEY", "?фамилия=")
	v.Add("apiKEY", "имя=")
	v.Add("apiKEY", "отчество=")
}

// GetCompanyInfo возвращает массив CompanyInfo
func GetCompanyInfo(data string) ([]CompanyInfo, error) {
	var url string

	switch {
	//если это ОГРН
	case regOgrn.MatchString(data):
		url = v.Get("baseURL") + v["apiURL"][0] + v["apiKEY"][0] + data
	// если это ИНН
	case regInn.MatchString(data):
		url = v.Get("baseURL") + v["apiURL"][0] + v["apiKEY"][1] + data
	// или название
	default:
		url = v.Get("baseURL") + v["apiURL"][0] + v["apiKEY"][2] + data + "&стр=5"
	}

	var companyJSON []CompanyInfo

	err := json.Unmarshal(response(url), &companyJSON)
	if err != nil {
		return nil, err
	}

	return companyJSON, nil
}

// GetEmployees Возвращает данные об сотрудниках компании с идентификатором {id}
func GetEmployees(id string) ([]Employees, error) {
	url := v.Get("baseURL") + v["apiURL"][0] + id + v["apiURL"][3]

	var companyJSON []Employees

	err := json.Unmarshal(response(url), &companyJSON)
	if err != nil {
		return nil, err
	}

	return companyJSON, nil
}

// GetIndivEntrep поиск индивидуального предпринимателя
func GetIndivEntrep(data string) ([]IEData, error) {
	var url string

	switch {
	//если это ОГРНИП
	case regOgrnip.MatchString(data):
		url = v.Get("baseURL") + v["apiURL"][2] + v["apiKEY"][3] + data
	// если это ИНН
	case regInn.MatchString(data):
		url = v.Get("baseURL") + v["apiURL"][2] + v["apiKEY"][1] + data
	// или название
	default:
		url = v.Get("baseURL") + v["apiURL"][2] + v["apiKEY"][4] + data
	}

	var indiventrepJSON []IEData

	err := json.Unmarshal(response(url), &indiventrepJSON)
	if err != nil {
		return nil, err
	}

	return indiventrepJSON, nil
}

// GetCompFounder Возвращает данные о компаниях, в которых данный человек является учредителем
func GetCompFounder(id string) ([]CompanyInfo, error) {
	url := v.Get("baseURL") + v["apiURL"][1] + id + v["apiURL"][0]

	var compsFounderJSON []CompanyInfo

	err := json.Unmarshal(response(url), &compsFounderJSON)
	if err != nil {
		return nil, err
	}

	return compsFounderJSON, nil
}

// GetDepComp Возвращает данные о зависимых компаниях от компании с идентификатором {id}
func GetDepComp(id string) ([]CompanyInfo, error) {
	url := v.Get("baseURL") + v["apiURL"][0] + id + v["apiURL"][5]

	var companyJSON []CompanyInfo

	err := json.Unmarshal(response(url), &companyJSON)
	if err != nil {
		return nil, err
	}

	return companyJSON, nil
}

// GetPerson Возвращает данные о человеке с идентификатором {id}
func GetPerson(id string) (Person, error) {
	url := v.Get("baseURL") + v["apiURL"][1] + id + "/"

	var personJSON Person

	err := json.Unmarshal(response(url), &personJSON)
	if err != nil {
		return personJSON, err
	}

	return personJSON, nil
}

// GetPositions Возвращает данные о должностях с идентификатором {id}
func GetPositions(id string) ([]Positions, error) {
	url := v.Get("baseURL") + v["apiURL"][1] + id + v["apiURL"][6]

	var personsJSON []Positions

	err := json.Unmarshal(response(url), &personsJSON)
	if err != nil {
		return nil, err
	}

	return personsJSON, nil
}

// GetIDData возвращает IDData
func GetIDData(id string) (IDData, error) {
	url := v.Get("baseURL") + v["apiURL"][0] + id + "/"

	var companyJSON IDData

	err := json.Unmarshal(response(url), &companyJSON)
	if err != nil {
		return companyJSON, err
	}

	return companyJSON, nil
}

// GetFounders Возвращает данные об учредителях компании с идентификатором {id}
func GetFounders(id string) ([]Founders, error) {
	url := v.Get("baseURL") + v["apiURL"][0] + id + v["apiURL"][7]

	var companyJSON []Founders

	err := json.Unmarshal(response(url), &companyJSON)
	if err != nil {
		return nil, err
	}

	return companyJSON, nil
}

// GetSearchPerson - поиск человека по ФИО или ИНН
// Порядок ввода ФИО - Фамилия, Имя, Отчество
func GetSearchPerson(data string) ([]PersonOwner, error) {
	var url string

	switch {
	// если это ИНН
	case regInn.MatchString(data):
		url = v.Get("baseURL") + v["apiURL"][1] + v["apiKEY"][1] + data
	// или ФИО
	default:
		fio := strings.Split(data, " ")
		// v["apiKEY"][5] - Ф, [6] - И, [7] - О
		switch {
		case len(fio) == 1:
			url = v.Get("baseURL") + v["apiURL"][1] + v["apiKEY"][5] + fio[0]
		case len(fio) == 2:
			url = v.Get("baseURL") + v["apiURL"][1] + v["apiKEY"][5] + fio[0] + "&" + v["apiKEY"][6] + fio[1]
		case len(fio) == 3:
			url = v.Get("baseURL") + v["apiURL"][1] + v["apiKEY"][5] + fio[0] + "&" + v["apiKEY"][6] + fio[1] + "&" + v["apiKEY"][7] + fio[2]
		default:
			fmt.Printf("Неккоректный ввод")
		}
	}

	var personJSON []PersonOwner

	err := json.Unmarshal(response(url), &personJSON)
	if err != nil {
		return nil, err
	}

	return personJSON, nil
}

func response(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Ошибка http.Get %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка ioutil.ReadAll %v", err)
	}

	return body
}
