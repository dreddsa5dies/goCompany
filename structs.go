package ogrnOnline

// CompanyBaseInfo - базовая информация о юридическом лице
type CompanyBaseInfo struct {
	ID          int       `json:"id"`                    // ID
	OGRN        string    `json:"ogrn"`                  // ОГРН
	Name        string    `json:"name"`                  // Полное наименование
	ShortName   string    `json:"shortName"`             // Сокращенное наименование
	OgrnDate    string    `json:"ogrnDate"`              // Дата присвоения ОГРН
	INN         string    `json:"inn"`                   // ИНН
	KPP         string    `json:"kpp"`                   // КПП
	URL         string    `json:"url"`                   // URL
	Description string    `json:"description,omitempty"` // Описание
	CloseInfo   closeInfo `json:"closeInfo,omitempty"`   // Cведения о прекращении деятельности
}

// CompanyInfo - полная информация о юридическом лице
type CompanyInfo struct {
	CompanyBaseInfo
	Address             address           `json:"address"`                       // Адрес
	Email               string            `json:"email,omitempty"`               // Email
	AuthorizedCapital   authorizedCapital `json:"authorizedCapital,omitempty"`   // Уставной капитал
	LastUpdateDate      string            `json:"lastUpdateDate"`                // Дата последнего обновления информации
	Assignee            []CompanyBaseInfo `json:"assignee"`                      // Правопреемники
	Predecessor         []CompanyBaseInfo `json:"predecessor"`                   // Правопредшественники
	MainOkved           okved             `json:"mainOkved,omitempty"`           // Главный ОКВЭД
	MainOkved1          okved             `json:"mainOkved1,omitempty"`          // Главный ОКВЭД 1
	MainOkved2          okved             `json:"mainOkved2,omitempty"`          // Главный ОКВЭД 2
	Okved               []okved           `json:"okved"`                         // Список ОКВЭД
	Okved1              []okved           `json:"okved1"`                        // Список ОКВЭД 1
	Okved2              []okved           `json:"okved2"`                        // Список ОКВЭД 2
	Fns                 fns               `json:"fns,omitempty"`                 // Cведения об учете в налоговом органе
	PfrRegistration     pfrRegistration   `json:"pfrRegistration,omitempty"`     // Сведения о регистрации в ПФР
	FssRegistration     fssRegistration   `json:"fssRegistration,omitempty"`     // Сведения о регистрации в ФСС
	StockRegisterHolder CompanyBaseInfo   `json:"stockRegisterHolder,omitempty"` // Сведения о регистраторе ценных бумаг
}

// CompanyOwnerInfo - информация об участнике юридического лица
type CompanyOwnerInfo struct {
	ID           int             `json:"id"`                     // ID
	Company      CompanyBaseInfo `json:"company"`                // Информация о юридическом лице
	PersonOwner  PeopleInfo      `json:"personOwner,omitempty"`  // Участник - физическое лицо
	CompanyOwner CompanyBaseInfo `json:"companyOwner,omitempty"` // Участник - юридическое лицо
	Price        float64         `json:"price"`                  // Номинальная стоимость доли
	OwnerRussia  bool            `json:"ownerRussia"`            // Участник - РФ
	Part         string          `json:"part"`                   // Доля
}

// CompanyAssociateInfo - информация об управляющих
type CompanyAssociateInfo struct {
	ID            int             `json:"id"`            // ID
	Company       CompanyBaseInfo `json:"company"`       // Информация о юридическом лице
	Person        PeopleInfo      `json:"person"`        // Информация об управляющем
	Post          base            `json:"post"`          // Информация о должности
	PostName      string          `json:"postName"`      // Наименование должности
	Phone         string          `json:"phone"`         // Телефон
	Unreliability []interface{}   `json:"unreliability"` //
}

// CompanyBranchInfo - филиал / представительство юридического лица
type CompanyBranchInfo struct {
	ID      int     `json:"id"`      // ID
	Company id      `json:"company"` // Информация о юридическом лице
	Branch  bool    `json:"branch"`  // Филиал
	Agency  bool    `json:"agency"`  // Представительство
	Name    string  `json:"name"`    // Наименование
	Address address `json:"address"` // Адрес
}

// PeopleInfo - полная информация о физическом лице
type PeopleInfo struct {
	ID              int           `json:"id"`              // ID
	FirstName       string        `json:"firstName"`       // Имя
	MiddleName      string        `json:"middleName"`      // Отчество
	SurName         string        `json:"surName"`         // Фамилия
	Inn             string        `json:"inn"`             // ИНН
	FullName        string        `json:"fullName"`        // Полное имя
	FullNameWithInn string        `json:"fullNameWithInn"` // Полное имя с ИНН
	URL             string        `json:"url"`             // URL
	OtherNames      []interface{} `json:"otherNames"`      // Другие имена
}

// PeopleBusinessmanInfo - информация об индивидуальном предпринимателе
type PeopleBusinessmanInfo struct {
	ID             int           `json:"id"`             // ID
	Person         PeopleInfo    `json:"person"`         // Bнформация о самом физическом лице
	LastUpdateDate string        `json:"lastUpdateDate"` // Дата последнего обновления сведений
	MainOkved      okved         `json:"mainOkved"`      // Главный ОКВЭД 1
	MainOkved1     okved         `json:"mainOkved1"`     // Главный ОКВЭД 1
	Okved          []okved       `json:"okved"`          // Список ОКВЭД
	Okved1         []okved       `json:"okved1"`         // Список ОКВЭД 1
	Okved2         []okved       `json:"okved2"`         // Список ОКВЭД 2
	CloseInfo      closeInfo     `json:"closeInfo"`      // Информация о прекращении деятельности
	OGRN           string        `json:"ogrn"`           // ОГРНИП
	OgrnDate       string        `json:"ogrnDate"`       // Дата присвоения ОГРНИП
	Citizenship    citizenship   `json:"citizenship"`    // Гражданство
	Fns            fns           `json:"fns"`            // Информация об учете в налоговом органе
	Type           base          `json:"type"`           // Тип
	License        []interface{} `json:"license"`        // Информация о лицензиях
}

// base - базовая структура, содержащая часто используемы поля
type base struct {
	ID       int    `json:"id"`       // ID
	FullName string `json:"fullName"` // Полное наименование
	Name     string `json:"name"`     // Сокращенное наименование
	Code     string `json:"code"`     // Код
}

// registration - базовая стуктура сведений о регистрации
type registration struct {
	Number           string `json:"number"`           // Номер
	RegistrationDate string `json:"registrationDate"` // Дата регистрации
}

// authorizedCapital - уставной капитал
type authorizedCapital struct {
	Type struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"type"`
	Value float64 `json:"value"`
}

// okopf - ОКОПФ (Общероссийский классификатор организационно-правовых форм)
type okopf struct {
	base
	Parent id `json:"parent"` //
}

// okved - ОКВЭД (Общероссийский классификатор видов экономической деятельности)
type okved struct {
	base
	Description  string `json:"description,omitempty"`  // Описание
	CompanyCount int    `json:"companyCount,omitempty"` //
	Parent       id     `json:"parent"`                 //
	URL          string `json:"url"`                    // URL
}

// pfrRegistration - сведения о регистрации в ПФР
type pfrRegistration struct {
	registration
	Pfr base `json:"pfr"`
}

// pfrRegistration - сведения о регистрации в ФСС
type fssRegistration struct {
	registration
	Fss base `json:"fss"`
}

// fns - сведения об учете в налоговом органе
type fns struct {
	base
	Address string `json:"address"` // Адрес
}

// closeInfo - сведения о прекращении деятельности
type closeInfo struct {
	Date        string `json:"date"`        // Дата прекращения деятельности
	CloseReason id     `json:"closeReason"` // Причина прекращения деятельности
}

// id - объект, содержащий только поле id
type id struct {
	ID int `json:"id"` //ID
}

// address - адрес
type address struct {
	Region           addressObject `json:"region"`      // Регион
	Street           addressObject `json:"street"`      // Улица
	House            string        `json:"house"`       // Дом
	Building         string        `json:"building"`    // Строение
	Flat             string        `json:"flat"`        // Квартира
	PostalIndex      string        `json:"postalIndex"` // Почтовый индекс
	FiasOnlineLink   string        `json:"fiasOnlineLink"`
	FullAddress      string        `json:"fullAddress"`
	FullHouseAddress string        `json:"fullHouseAddress"`
}

// addressObject - адресный объект (улица, город, регион и пр.)
type addressObject struct {
	ID            int               `json:"id"`            // ID
	TypeName      string            `json:"typeName"`      // Тип полного наименование
	TypeShortName string            `json:"typeShortName"` // Тип сокращенного наименование
	FullName      string            `json:"fullName"`      // Полное наименование
	Name          string            `json:"name"`          // Сокращенное наименование
	Aoid          string            `json:"aoid"`
	GUID          string            `json:"guid"`
	PostalCode    string            `json:"postalCode"` // Почтовый индекс
	Level         int               `json:"level"`
	Okato         string            `json:"okato"`      // ОКАТО (Общероссийский классификатор административно-территориальных образований)
	Type          typeAddressObject `json:"type"`       // Тип адресного объекта
	RegionCode    string            `json:"regionCode"` // Код региона
	AutoCode      string            `json:"autoCode"`
	AreaCode      string            `json:"areaCode"`
	CityCode      string            `json:"cityCode"`
	CtarCode      string            `json:"ctarCode"`
	PlaceCode     string            `json:"placeCode"`
	StreetCode    string            `json:"streetCode"` // Код улицы
	ExtrCode      string            `json:"extrCode"`
	SextCode      string            `json:"sextCode"`
	KladrCode     string            `json:"kladrCode"`
	Live          bool              `json:"live"`
	URL           string            `json:"url"`
	CompanyCount  int               `json:"companyCount"`
}

// typeAddressObject - тип адресного объекта
type typeAddressObject struct {
	ID        int    `json:"id"`        // ID
	Name      string `json:"name"`      // Полное наименование
	ShortName string `json:"shortName"` // Сокращенное наименование
	Code      string `json:"code"`      // Код
	Level     int    `json:"level"`     // Уровень
}

// citizenship - гражданство
type citizenship struct {
	Russian bool `json:"russian"` // Гражданство РФ
}
