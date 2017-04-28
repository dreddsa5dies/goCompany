package goCompany

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
)

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
	TypeIE         Fss      `json:"type"`
	License        []string `json:"license"`
}

// Citizenship - гражданин РФ
type Citizenship struct {
	Russian bool `json:"russian"`
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

	// обращение к API
	resp, err := http.Get(url)
	if err != nil {
		return indiventrepJSON, err
	}
	defer resp.Body.Close()

	// запись ответа в переменную
	body, err := ioutil.ReadAll(resp.Body)

	// декодирование []byte в интерфейс
	err = json.Unmarshal(body, &indiventrepJSON)
	if err != nil {
		return indiventrepJSON, err
	}

	return indiventrepJSON, nil
}
