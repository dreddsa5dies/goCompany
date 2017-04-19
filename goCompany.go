// urlScrab
package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// набор флагов
var (
	golog   bool
	company string
)

// инициализация флагов
func init() {
	flag.BoolVar(&golog, "log", false, "Создание лог-файла")

	flag.StringVar(&company, "company", "Соболь", "Имя компании")
}

func main() {
	// разбор флагов
	flag.Parse()

	if golog {
		// создание файла log для записи ошибок
		fLog, err := os.OpenFile(`./.log`, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0640)
		if err != nil {
			log.Fatalln(err)
		}
		defer fLog.Close()

		// запись ошибок и инфы в файл и вывод
		log.SetOutput(io.MultiWriter(fLog, os.Stdout))
	}

	// GET /интеграция/компании/
	// Поиск компании по ИНН, ОГРН или по названию. С помощью этой команды можно получить идентификатор нужной компании и далее с помощью следующих команд получить детальную информацию.
	url := "https://ru.rus.company/интеграция/компании/?наименование=" + company + "&стр=5"
	log.Printf("url = %v", url)

	// обращение к API
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	log.Printf("status response = %v", resp.Status)

	// запись ответа в переменную
	body, err := ioutil.ReadAll(resp.Body)

	type SaveData struct {
		ID        int    `json:"id"`
		OGRN      string `json:"ogrn"`
		NAME      string `json:"name"`
		SHORTNAME string `json:"shortName"`
		OGRNDATE  string `json:"ogrnDate"`
		INN       string `json:"inn"`
		KPP       string `json:"kpp"`
		URL       string `json:"url"`
	}

	// создание массива переменных для хранения ответа
	var companyJSON []SaveData

	// декодирование []byte в интерфейс
	err = json.Unmarshal(body, &companyJSON)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// создание файла отчета в формате csv
	file, err := os.OpenFile("./reports.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalln(err)

	}
	defer file.Close()
	log.Printf("Создание файла отчета %v", file.Name())

	// заголовок
	getFile, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	if getFile.Size() <= 1 {
		// заголовок
		file.WriteString("ID;INN;KPP;NAME;OGRN;OGRNDATE;SHOTNAME;URL")
		file.WriteString("\n")
	}

	for i := 0; i < len(companyJSON); i++ {
		log.Printf("Запись данных о %v", companyJSON[i].SHORTNAME)
		file.WriteString(strconv.Itoa(companyJSON[i].ID) + ";")
		file.WriteString(companyJSON[i].INN + ";")
		file.WriteString(companyJSON[i].KPP + ";")
		file.WriteString(companyJSON[i].NAME + ";")
		file.WriteString(companyJSON[i].OGRN + ";")
		file.WriteString(companyJSON[i].OGRNDATE + ";")
		file.WriteString(companyJSON[i].SHORTNAME + ";")
		file.WriteString("https://ru.rus.company/интеграция/" + companyJSON[i].URL)
		file.WriteString("\n")
	}
}
