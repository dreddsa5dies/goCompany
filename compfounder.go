package goCompany

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetCompFounder Возвращает данные о компаниях, в которых данный человек является учредителем
func GetCompFounder(id string) ([]BaseData, error) {
	url := baseURLPers + id + "/компании/"

	// создание массива переменных для хранения ответа
	var compsFounderJSON []BaseData

	// обращение к API
	resp, err := http.Get(url)
	if err != nil {
		return compsFounderJSON, err
	}
	defer resp.Body.Close()

	// запись ответа в переменную
	body, err := ioutil.ReadAll(resp.Body)

	// декодирование []byte в интерфейс
	err = json.Unmarshal(body, &compsFounderJSON)
	if err != nil {
		return compsFounderJSON, err
	}

	return compsFounderJSON, nil
}
