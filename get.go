package ogrnOnline

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

const (
	typeQueryCompany = iota
	typeQueryPeople
	typeQueryBusinessman

	host = `https://огрн.онлайн`

	pauseForRequest = 400
)

// isValidQuery проверяет запрос на верность ключевых слов и параметров
func isValidQuery(query url.Values, typeQuery int) error {
	// testType определяет допустимые ключевые слова
	testType := map[int]map[string]bool{
		typeQueryCompany:     map[string]bool{"огрн": true, "инн": true, "кпп": true, "наименование": true, "стр": true},
		typeQueryPeople:      map[string]bool{"фамилия": true, "имя": true, "отчество": true, "инн": true, "стр": true},
		typeQueryBusinessman: map[string]bool{"человек": true, "огрнип": true, "инн": true},
	}

	// Регулярными выражениями проверяется использование допустимых символов в параметрах
	ogrnBusinessman := regexp.MustCompile(`^[0-9]{15}$`)
	ogrnCompany := regexp.MustCompile(`^[0-9]{13}$`)
	innBusinessman := regexp.MustCompile(`^[0-9]{12}$`)
	innCompany := regexp.MustCompile(`^[0-9]{10}$`)

	if typeQuery != typeQueryCompany && typeQuery != typeQueryPeople && typeQuery != typeQueryBusinessman {
		panic(fmt.Errorf("isValidQuery: неверный параметр typeQuery: %d", typeQuery))
	}

	for options := range query {
		if !testType[typeQuery][options] {
			return fmt.Errorf("isValidQuery: неверный параметр: %s", options)
		}

		switch {
		case options == "огрн" && typeQuery == typeQueryCompany:
			if !ogrnCompany.MatchString(query[options][0]) {
				return fmt.Errorf("isValidQuery: недопустимое значение ОГРН: %s", query[options][0])
			}
		case options == "инн" && typeQuery == typeQueryCompany:
			if !innCompany.MatchString(query[options][0]) {
				return fmt.Errorf("isValidQuery: недопустимое значение ИНН: %s", query[options][0])
			}
		case options == "огрнип" && typeQuery == typeQueryBusinessman:
			if !ogrnBusinessman.MatchString(query[options][0]) {
				return fmt.Errorf("isValidQuery: недопустимое значение ОГРНИП: %s", query[options][0])
			}
		case options == "инн" && (typeQuery == typeQueryBusinessman || typeQuery == typeQueryPeople):
			if !innBusinessman.MatchString(query[options][0]) {
				return fmt.Errorf("isValidQuery: недопустимое значение ИНН: %s", query[options][0])
			}
		}
	}
	return nil
}

// createURL формирует URL на основе пути и запроса
func createURL(path string, query url.Values) *url.URL {
	ur, err := url.Parse(host)
	if err != nil {
		panic(fmt.Errorf("createURL: ошибка парсинга хоста: %v", err))
	}
	ur.Path = path
	if query != nil {
		ur.RawQuery = query.Encode()
	}
	return ur
}

// getDataFromServer - запрос к API
func getDataFromServer(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(fmt.Errorf("getDataFromServer: ошибка запроса к серверу: %v", err))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("getDataFromServer: ошибка чтения ответа сервера: %v", err))
	}
	time.Sleep(time.Millisecond * pauseForRequest) // Пауза для предотвращения перегрузки сервера
	return body
}

// companyNotExist возвращает true если компания прекратила деятельность
func companyNotExist(c *CompanyInfo) bool {
	return c.CloseInfo.Date != ""
}

// FindCompany осуществляет поиск юридического лица по заданным параметрам
func FindCompany(query url.Values) ([]CompanyInfo, error) {
	var (
		path        = `/интеграция/компании/`
		result      = []CompanyInfo{}
		cleanResult = []CompanyInfo{}
	)

	if err := isValidQuery(query, typeQueryCompany); err != nil {
		return []CompanyInfo{}, fmt.Errorf("FindCompany: %v", err)
	}

	if err := json.Unmarshal(getDataFromServer(createURL(path, query).String()), &result); err != nil {
		return []CompanyInfo{}, fmt.Errorf("FindCompany: %v", err)
	}
	for _, company := range result {
		if !companyNotExist(&company) {
			cleanResult = append(cleanResult, company)
		}
	}
	return cleanResult, nil
}

// FindPeople осуществляет поиск физического лица по заданным параметрам
func FindPeople(query url.Values) ([]PeopleInfo, error) {
	var (
		path   = `/интеграция/люди/`
		result = []PeopleInfo{}
	)
	if err := isValidQuery(query, typeQueryPeople); err != nil {
		return []PeopleInfo{}, fmt.Errorf("FindPeople: %v", err)
	}
	if err := json.Unmarshal(getDataFromServer(createURL(path, query).String()), &result); err != nil {
		return []PeopleInfo{}, fmt.Errorf("FindPeople: %v", err)
	}
	return result, nil
}

// FindBusinessman осуществляет поиск индивидуального предпринимателя по заданным параметрам
func FindBusinessman(query url.Values) ([]PeopleBusinessmanInfo, error) {
	var (
		path   = `/интеграция/ип/`
		result = []PeopleBusinessmanInfo{}
	)
	if err := isValidQuery(query, typeQueryBusinessman); err != nil {
		return []PeopleBusinessmanInfo{}, fmt.Errorf("FindBusinessman: %v", err)
	}
	if err := json.Unmarshal(getDataFromServer(createURL(path, query).String()), &result); err != nil {
		return []PeopleBusinessmanInfo{}, fmt.Errorf("FindBusinessman: %v", err)
	}
	return result, nil
}

// GetCompanyFullInfo возвращает полную информацию о юридическом лице на основе его id
func GetCompanyFullInfo(id int) (CompanyFullInfo, error) {
	var (
		path   = `/интеграция/компании/`
		result = CompanyFullInfo{}
		param  = "/"
	)
	if err := json.Unmarshal(getDataFromServer(createURL(fmt.Sprintf(`%s%d%s`, path, id, param), nil).String()), &result); err != nil {
		return CompanyFullInfo{}, fmt.Errorf("GetCompanyFullInfo: %v", err)
	}
	return result, nil
}

// GetCompanyFullInfo возвращает полную информацию о юридическом лице
func (c *CompanyInfo) GetCompany() (CompanyFullInfo, error) {
	return GetCompanyFullInfo(c.ID)
}

// GetOwners возвращает список участников юридического лица на основе его id
func GetOwners(id int) ([]CompanyOwnerInfo, error) {
	var (
		path        = `/интеграция/компании/`
		result      = []CompanyOwnerInfo{}
		cleanResult = []CompanyOwnerInfo{}
		param       = "/учредители/"
	)
	if err := json.Unmarshal(getDataFromServer(createURL(fmt.Sprintf(`%s%d%s`, path, id, param), nil).String()), &result); err != nil {
		return []CompanyOwnerInfo{}, fmt.Errorf("GetOwners: %v", err)
	}
	for _, owner := range result {
		if !companyNotExist(&owner.CompanyOwner) {
			cleanResult = append(cleanResult, owner)
		}
	}
	return cleanResult, nil
}

// GetOwners возвращает список участников юридического лица
func (c *CompanyInfo) GetOwners() ([]CompanyOwnerInfo, error) {
	return GetOwners(c.ID)
}

// GetAssociates возвращает список управляющих юридического лица на основе его id
func GetAssociates(id int) ([]CompanyAssociateInfo, error) {
	var (
		path   = `/интеграция/компании/`
		result = []CompanyAssociateInfo{}
		param  = "/сотрудники/"
	)
	if err := json.Unmarshal(getDataFromServer(createURL(fmt.Sprintf(`%s%d%s`, path, id, param), nil).String()), &result); err != nil {
		return []CompanyAssociateInfo{}, fmt.Errorf("GetAssociates: %v", err)
	}
	return result, nil
}

// GetAssociates возвращает список управляющих юридического лица
func (c *CompanyInfo) GetAssociates() ([]CompanyAssociateInfo, error) {
	return GetAssociates(c.ID)
}

// GetSubCompanies возвращает список зависимых компаний юридического лица на основе его id
func GetSubCompanies(id int) ([]CompanyInfo, error) {
	var (
		path        = `/интеграция/компании/`
		result      = []CompanyInfo{}
		cleanResult = []CompanyInfo{}
		param       = "/зависимые/"
	)
	if err := json.Unmarshal(getDataFromServer(createURL(fmt.Sprintf(`%s%d%s`, path, id, param), nil).String()), &result); err != nil {
		return []CompanyInfo{}, fmt.Errorf("GetSubCompanies: %v", err)
	}
	for _, company := range result {
		if !companyNotExist(&company) {
			cleanResult = append(cleanResult, company)
		}
	}
	return cleanResult, nil
}

// GetSubCompanies возвращает список зависимых компаний юридического лица
func (c *CompanyInfo) GetSubCompanies() ([]CompanyInfo, error) {
	return GetSubCompanies(c.ID)
}

// GetRepresentativeOffices возвращает список представительсв юридического лица на основе его id
func GetRepresentativeOffices(id int) ([]CompanyBranchInfo, error) {
	var (
		path   = `/интеграция/компании/`
		result = []CompanyBranchInfo{}
		param  = "/представительства/"
	)
	if err := json.Unmarshal(getDataFromServer(createURL(fmt.Sprintf(`%s%d%s`, path, id, param), nil).String()), &result); err != nil {
		return []CompanyBranchInfo{}, fmt.Errorf("GetRepresentativeOffices: %v", err)
	}
	return result, nil
}

// GetRepresentativeOffices возвращает список представительсв юридического лица
func (c *CompanyInfo) GetRepresentativeOffices() ([]CompanyBranchInfo, error) {
	return GetRepresentativeOffices(c.ID)
}

// GetBranches возвращает список филиалов юридического лица на основе его id
func GetBranches(id int) ([]CompanyBranchInfo, error) {
	var (
		path   = `/интеграция/компании/`
		result = []CompanyBranchInfo{}
		param  = "/филиалы/"
	)
	if err := json.Unmarshal(getDataFromServer(createURL(fmt.Sprintf(`%s%d%s`, path, id, param), nil).String()), &result); err != nil {
		return []CompanyBranchInfo{}, fmt.Errorf("GetBranches: %v", err)
	}
	return result, nil
}

// GetBranches возвращает список филиалов юридического лица на основе его id
func (c *CompanyInfo) GetBranches() ([]CompanyBranchInfo, error) {
	return GetBranches(c.ID)
}

// GenFinance возвращает бухгалтерские балансы юридиеского лица за предшествующие годы на основе его id
func GenFinance(id int) ([]CompanyFinanceInfo, error) {
	var (
		path   = `/интеграция/компании/`
		result = []CompanyFinanceInfo{}
		param  = "/финансы/"
	)
	if err := json.Unmarshal(getDataFromServer(createURL(fmt.Sprintf(`%s%d%s`, path, id, param), nil).String()), &result); err != nil {
		return []CompanyFinanceInfo{}, fmt.Errorf("GenFinance: %v", err)
	}
	return result, nil
}

// GenFinance возвращает бухгалтерские балансы юридиеского лица за предшествующие годы
func (c *CompanyInfo) GetFinance() ([]CompanyFinanceInfo, error) {
	return GenFinance(c.ID)
}

// GetPeople полную информацию о физическом лице на основе его id
func GetPeople(id int) (PeopleInfo, error) {
	var (
		path   = `/интеграция/люди/`
		result = PeopleInfo{}
		param  = "/"
	)
	if err := json.Unmarshal(getDataFromServer(createURL(fmt.Sprintf(`%s%d%s`, path, id, param), nil).String()), &result); err != nil {
		return PeopleInfo{}, fmt.Errorf("GetPeople: %v", err)
	}
	return result, nil
}

// GetJobs возвращает места работы физического лица на основе его id
func GetJobs(id int) ([]CompanyAssociateInfo, error) {
	var (
		path        = `/интеграция/люди/`
		result      = []CompanyAssociateInfo{}
		cleanResult = []CompanyAssociateInfo{}
		param       = "/должности/"
	)
	if err := json.Unmarshal(getDataFromServer(createURL(fmt.Sprintf(`%s%d%s`, path, id, param), nil).String()), &result); err != nil {
		return []CompanyAssociateInfo{}, fmt.Errorf("GetJobs: %v", err)
	}
	for _, work := range result {
		if !companyNotExist(&work.Company) {
			cleanResult = append(cleanResult, work)
		}
	}
	return cleanResult, nil
}

// GetJobs возвращает места работы физического лица
func (p *PeopleInfo) GetJobs() ([]CompanyAssociateInfo, error) {
	return GetJobs(p.ID)
}

// GetShare возвращает список компаний c участием физического лица на основе его id
func GetShare(id int) ([]CompanyInfo, error) {
	var (
		path        = `/интеграция/люди/`
		result      = []CompanyInfo{}
		cleanResult = []CompanyInfo{}
		param       = "/компании/"
	)
	if err := json.Unmarshal(getDataFromServer(createURL(fmt.Sprintf(`%s%d%s`, path, id, param), nil).String()), &result); err != nil {
		return []CompanyInfo{}, fmt.Errorf("GetShare: %v", err)
	}
	for _, share := range result {
		if !companyNotExist(&share) {
			cleanResult = append(cleanResult, share)
		}
	}
	return cleanResult, nil
}

// GetShare - метод PeopleInfo, возвращающий список компаний c участием физического лица
func (p *PeopleInfo) GetShare() ([]CompanyInfo, error) {
	return GetShare(p.ID)
}
