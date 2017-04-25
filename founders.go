package goCompany

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Founders - структура данных об учредителях компании
type Founders struct {
	ID          int `json:"id"`
	BaseData    `json:"company"`
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

// GetFounders Возвращает данные об учредителях компании с идентификатором {id}
func GetFounders(id string) ([]Founders, error) {
	url := baseURL + id + "/учредители/"

	// создание массива переменных для хранения ответа
	var companyJSON []Founders

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
