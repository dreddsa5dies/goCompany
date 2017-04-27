package goCompany

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetDepComp Возвращает данные о зависимых компаниях от компании с идентификатором {id}
func GetDepComp(id string) ([]BaseData, error) {
	url := baseURL + id + "/зависимые/"

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
