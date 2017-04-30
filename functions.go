package goCompany

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

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

	// декодирование []byte в интерфейс
	err := json.Unmarshal(response(url), &companyJSON)
	if err != nil {
		return companyJSON, err
	}

	return companyJSON, nil
}

// GetEmployees Возвращает данные об сотрудниках компании с идентификатором {id}
func GetEmployees(id string) ([]Employees, error) {
	url := baseURL + id + "/сотрудники/"

	// создание массива переменных для хранения ответа
	var companyJSON []Employees

	// декодирование []byte в интерфейс
	err := json.Unmarshal(response(url), &companyJSON)
	if err != nil {
		return companyJSON, err
	}

	return companyJSON, nil
}

// GetIndivEntrep поиск индивидуального предпринимателя
func GetIndivEntrep(data string) ([]IEData, error) {
	var url string

	// блоки для проверки по регулярному выражению
	regOgrnip, _ := regexp.Compile(`([0-9]){15}`)
	regInn, _ := regexp.Compile(`([0-9]){10,12}`)

	switch {
	//если это ОГРНИП
	case regOgrnip.MatchString(data):
		url = baseURLIE + "?огрнип=" + data
	// если это ИНН
	case regInn.MatchString(data):
		url = baseURLIE + "?инн=" + data
	// или название
	default:
		url = baseURLIE + "?человек=" + data
	}

	// создание массива переменных для хранения ответа
	var indiventrepJSON []IEData

	// декодирование []byte в интерфейс
	err := json.Unmarshal(response(url), &indiventrepJSON)
	if err != nil {
		return indiventrepJSON, err
	}

	return indiventrepJSON, nil
}

// GetCompFounder Возвращает данные о компаниях, в которых данный человек является учредителем
func GetCompFounder(id string) ([]BaseData, error) {
	url := baseURLPers + id + "/компании/"

	// создание массива переменных для хранения ответа
	var compsFounderJSON []BaseData

	// декодирование []byte в интерфейс
	err := json.Unmarshal(response(url), &compsFounderJSON)
	if err != nil {
		return compsFounderJSON, err
	}

	return compsFounderJSON, nil
}

// GetDepComp Возвращает данные о зависимых компаниях от компании с идентификатором {id}
func GetDepComp(id string) ([]BaseData, error) {
	url := baseURL + id + "/зависимые/"

	// создание массива переменных для хранения ответа
	var companyJSON []BaseData

	// декодирование []byte в интерфейс
	err := json.Unmarshal(response(url), &companyJSON)
	if err != nil {
		return companyJSON, err
	}

	return companyJSON, nil
}

// GetPerson Возвращает данные о человеке с идентификатором {id}
func GetPerson(id string) (Person, error) {
	url := baseURLPers + id + "/"

	// создание массива переменных для хранения ответа
	var personJSON Person

	// декодирование []byte в интерфейс
	err := json.Unmarshal(response(url), &personJSON)
	if err != nil {
		return personJSON, err
	}

	return personJSON, nil
}

// GetPositions Возвращает данные о должностях с идентификатором {id}
func GetPositions(id string) ([]Positions, error) {
	url := baseURLPers + id + "/должности/"

	// создание массива переменных для хранения ответа
	var personsJSON []Positions

	// декодирование []byte в интерфейс
	err := json.Unmarshal(response(url), &personsJSON)
	if err != nil {
		return personsJSON, err
	}

	return personsJSON, nil
}

// GetIDData возвращает IDData
func GetIDData(id string) (IDData, error) {
	url := baseURL + id + "/"

	// создание массива переменных для хранения ответа
	var companyJSON IDData

	// декодирование []byte в интерфейс
	err := json.Unmarshal(response(url), &companyJSON)
	if err != nil {
		return companyJSON, err
	}

	return companyJSON, nil
}

// GetFounders Возвращает данные об учредителях компании с идентификатором {id}
func GetFounders(id string) ([]Founders, error) {
	url := baseURL + id + "/учредители/"

	// создание массива переменных для хранения ответа
	var companyJSON []Founders

	// декодирование []byte в интерфейс
	err := json.Unmarshal(response(url), &companyJSON)
	if err != nil {
		return companyJSON, err
	}

	return companyJSON, nil
}

func response(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Ошибка http.Get %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Ошибка ioutil.ReadAll %v", err)
	}

	return body
}
