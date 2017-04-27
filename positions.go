package goCompany

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Positions - данные о должностях по идентификатору {id}
type Positions struct {
	ID          int `json:"id"`
	BaseData    `json:"company"`
	PersonOwner `json:"person"`
	Post        `json:"post"`
	PostName    string `json:"postName"`
	Phone       string `json:"phone"`
}

// GetPositions Возвращает данные о должностях с идентификатором {id}
func GetPositions(id string) ([]Positions, error) {
	url := baseURLPers + id + "/должности/"

	// создание массива переменных для хранения ответа
	var personsJSON []Positions

	// обращение к API
	resp, err := http.Get(url)
	if err != nil {
		return personsJSON, err
	}
	defer resp.Body.Close()

	// запись ответа в переменную
	body, err := ioutil.ReadAll(resp.Body)

	// декодирование []byte в интерфейс
	err = json.Unmarshal(body, &personsJSON)
	if err != nil {
		return personsJSON, err
	}

	return personsJSON, nil
}
