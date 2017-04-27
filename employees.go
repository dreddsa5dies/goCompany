package goCompany

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Employees - структура данных об сотрудниках компании
type Employees struct {
	ID          int `json:"id"`
	BaseData    `json:"company"`
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

// GetEmployees Возвращает данные об сотрудниках компании с идентификатором {id}
func GetEmployees(id string) ([]Employees, error) {
	url := baseURL + id + "/сотрудники/"

	// создание массива переменных для хранения ответа
	var companyJSON []Employees

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
