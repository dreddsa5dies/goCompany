package goCompany

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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

// GetPerson Возвращает данные о человеке с идентификатором {id}
func GetPerson(id string) (Person, error) {
	url := baseURLPers + id + "/"

	// создание массива переменных для хранения ответа
	var personJSON Person

	// обращение к API
	resp, err := http.Get(url)
	if err != nil {
		return personJSON, err
	}
	defer resp.Body.Close()

	// запись ответа в переменную
	body, err := ioutil.ReadAll(resp.Body)

	// декодирование []byte в интерфейс
	err = json.Unmarshal(body, &personJSON)
	if err != nil {
		return personJSON, err
	}

	return personJSON, nil
}
