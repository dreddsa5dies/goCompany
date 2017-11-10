package ogrnOnline

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func Test_base_parse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		data       []byte
		wantStruct base
	}{
		{
			[]byte(`{"id" : 3, "name" : "РУКОВОДИТЕЛЬ ЮРИДИЧЕСКОГО ЛИЦА", "code" : "02", "fullName" : "02 РУКОВОДИТЕЛЬ ЮРИДИЧЕСКОГО ЛИЦА"}`),
			base{3, "02 РУКОВОДИТЕЛЬ ЮРИДИЧЕСКОГО ЛИЦА", "РУКОВОДИТЕЛЬ ЮРИДИЧЕСКОГО ЛИЦА", "02"},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			var obj base
			err := json.Unmarshal(tt.data, &obj)
			if err != nil {
				t.Fatalf("ошибка парсинга base: %v", err)
			}
			if !reflect.DeepEqual(tt.wantStruct, obj) {
				t.Fatalf("base == %v, want == %v", obj, tt.wantStruct)
			}
		})
	}
}
func Test_registration_parse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		data       []byte
		wantStruct registration
	}{
		{
			[]byte(`{"number" : "781009076778151", "registrationDate" : "2003-06-26T00:00:00.000"}`),
			registration{"781009076778151", "2003-06-26T00:00:00.000"},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			var obj registration
			err := json.Unmarshal(tt.data, &obj)
			if err != nil {
				t.Fatalf("ошибка парсинга registration: %v", err)
			}
			if !reflect.DeepEqual(tt.wantStruct, obj) {
				t.Fatalf("registration == %v, want == %v", obj, tt.wantStruct)
			}
		})
	}
}
func Test_fssRegistration_parse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		data       []byte
		wantStruct fssRegistration
	}{
		{
			[]byte(`{"number":"781009076778151","fss":{"id":59,"name":"ФИЛИАЛ №15 САНКТ-ПЕТЕРБУРГСКОГО РЕГИОНАЛЬНОГО ОТДЕЛЕНИЯ ФОНДА СОЦИАЛЬНОГО СТРАХОВАНИЯ РОССИЙСКОЙ ФЕДЕРАЦИИ","code":"7815","fullName":"7815 ФИЛИАЛ №15 САНКТ-ПЕТЕРБУРГСКОГО РЕГИОНАЛЬНОГО ОТДЕЛЕНИЯ ФОНДА СОЦИАЛЬНОГО СТРАХОВАНИЯ РОССИЙСКОЙ ФЕДЕРАЦИИ"},"registrationDate":"2003-06-26T00:00:00.000"}`),
			fssRegistration{registration{"781009076778151", "2003-06-26T00:00:00.000"}, base{59, "7815 ФИЛИАЛ №15 САНКТ-ПЕТЕРБУРГСКОГО РЕГИОНАЛЬНОГО ОТДЕЛЕНИЯ ФОНДА СОЦИАЛЬНОГО СТРАХОВАНИЯ РОССИЙСКОЙ ФЕДЕРАЦИИ", "ФИЛИАЛ №15 САНКТ-ПЕТЕРБУРГСКОГО РЕГИОНАЛЬНОГО ОТДЕЛЕНИЯ ФОНДА СОЦИАЛЬНОГО СТРАХОВАНИЯ РОССИЙСКОЙ ФЕДЕРАЦИИ", "7815"}},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			var obj fssRegistration
			err := json.Unmarshal(tt.data, &obj)
			if err != nil {
				t.Fatalf("ошибка парсинга fssRegistration: %v", err)
			}
			if !reflect.DeepEqual(tt.wantStruct, obj) {
				t.Fatalf("fssRegistration == %v, want == %v", obj, tt.wantStruct)
			}
		})
	}
}
func Test_pfrRegistration_parse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		data       []byte
		wantStruct pfrRegistration
	}{
		{
			[]byte(`{"number":"087803023999","pfr":{"id":120,"name":"ГОСУДАРСТВЕННОЕ УЧРЕЖДЕНИЕ - ГЛАВНОЕ УПРАВЛЕНИЕ ПЕНСИОННОГО ФОНДА РФ №2 УПРАВЛЕНИЕ №1 МУНИЦИПАЛЬНЫЙ РАЙОН МОЖАЙСКИЙ Г.МОСКВЫ","code":"087803","fullName":"087803 ГОСУДАРСТВЕННОЕ УЧРЕЖДЕНИЕ - ГЛАВНОЕ УПРАВЛЕНИЕ ПЕНСИОННОГО ФОНДА РФ №2 УПРАВЛЕНИЕ №1 МУНИЦИПАЛЬНЫЙ РАЙОН МОЖАЙСКИЙ Г.МОСКВЫ"},"registrationDate":"2001-05-23T00:00:00.000"}`),
			pfrRegistration{registration{"087803023999", "2001-05-23T00:00:00.000"}, base{120, "087803 ГОСУДАРСТВЕННОЕ УЧРЕЖДЕНИЕ - ГЛАВНОЕ УПРАВЛЕНИЕ ПЕНСИОННОГО ФОНДА РФ №2 УПРАВЛЕНИЕ №1 МУНИЦИПАЛЬНЫЙ РАЙОН МОЖАЙСКИЙ Г.МОСКВЫ", "ГОСУДАРСТВЕННОЕ УЧРЕЖДЕНИЕ - ГЛАВНОЕ УПРАВЛЕНИЕ ПЕНСИОННОГО ФОНДА РФ №2 УПРАВЛЕНИЕ №1 МУНИЦИПАЛЬНЫЙ РАЙОН МОЖАЙСКИЙ Г.МОСКВЫ", "087803"}},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			var obj pfrRegistration
			err := json.Unmarshal(tt.data, &obj)
			if err != nil {
				t.Fatalf("ошибка парсинга pfrRegistration: %v", err)
			}
			if !reflect.DeepEqual(tt.wantStruct, obj) {
				t.Fatalf("pfrRegistration == %v, want == %v", obj, tt.wantStruct)
			}
		})
	}
}
func Test_authorizedCapital_parse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		data       []byte
		wantStruct authorizedCapital
	}{
		{
			[]byte(`{"type":{"id":3,"name":"УСТАВНЫЙ КАПИТАЛ"},"value":25000000.00}`),
			authorizedCapital{
				Type: struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				}{ID: 3, Name: "УСТАВНЫЙ КАПИТАЛ"}, Value: 25000000.00,
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			var obj authorizedCapital
			err := json.Unmarshal(tt.data, &obj)
			if err != nil {
				t.Fatalf("ошибка парсинга authorizedCapital: %v", err)
			}
			if !reflect.DeepEqual(tt.wantStruct, obj) {
				t.Fatalf("authorizedCapital == %v, want == %v", obj, tt.wantStruct)
			}
		})
	}
}
func Test_okopf_parse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		data       []byte
		wantStruct okopf
	}{
		{
			[]byte(`{"id":165,"name":"Непубличные акционерные общества","code":"12267","parent":{"id":130}}`),
			okopf{base{165, "", "Непубличные акционерные общества", "12267"}, id{130}},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			var obj okopf
			err := json.Unmarshal(tt.data, &obj)
			if err != nil {
				t.Fatalf("ошибка парсинга okopf: %v", err)
			}
			if !reflect.DeepEqual(tt.wantStruct, obj) {
				t.Fatalf("okopf == %v, want == %v", obj, tt.wantStruct)
			}
		})
	}
}
func Test_okved_parse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		data       []byte
		wantStruct okved
	}{
		{
			[]byte(`{"id":2290,"name":"Управление эксплуатацией нежилого фонда за вознаграждение или на договорной основе","code":"68.32.2","parent":{"id":2288},"description":"\r\n","url":"/оквэд2/68_32_2-управление_эксплуатацией_нежилого_фонда_за_вознаграждение_или_на_договорной_основе/","fullName":"68.32.2 Управление эксплуатацией нежилого фонда за вознаграждение или на договорной основе"}`),
			okved{base{2290, "68.32.2 Управление эксплуатацией нежилого фонда за вознаграждение или на договорной основе", "Управление эксплуатацией нежилого фонда за вознаграждение или на договорной основе", "68.32.2"}, "\r\n", 0, id{2288}, "/оквэд2/68_32_2-управление_эксплуатацией_нежилого_фонда_за_вознаграждение_или_на_договорной_основе/"},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			var obj okved
			err := json.Unmarshal(tt.data, &obj)
			if err != nil {
				t.Fatalf("ошибка парсинга okved: %v", err)
			}
			if !reflect.DeepEqual(tt.wantStruct, obj) {
				t.Fatalf("okved == %v, want == %v", obj, tt.wantStruct)
			}
		})
	}
}
func Test_closeInfo_parse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		data       []byte
		wantStruct closeInfo
	}{
		{
			[]byte(`{"date":"2004-11-02T00:00:00.000","closeReason":{"id":4}}`),
			closeInfo{"2004-11-02T00:00:00.000", id{4}},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			var obj closeInfo
			err := json.Unmarshal(tt.data, &obj)
			if err != nil {
				t.Fatalf("ошибка парсинга closeInfo: %v", err)
			}
			if !reflect.DeepEqual(tt.wantStruct, obj) {
				t.Fatalf("closeInfo == %v, want == %v", obj, tt.wantStruct)
			}
		})
	}
}
func Test_typeAdressObject_parse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		data       []byte
		wantStruct typeAddressObject
	}{
		{
			[]byte(`{"id":6,"name":"Область","shortName":"обл","code":"105","level":1}`),
			typeAddressObject{6, "Область", "обл", "105", 1},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			var obj typeAddressObject
			err := json.Unmarshal(tt.data, &obj)
			if err != nil {
				t.Fatalf("ошибка парсинга typeAdressObject: %v", err)
			}
			if !reflect.DeepEqual(tt.wantStruct, obj) {
				t.Fatalf("typeAdressObject == %v, want == %v", obj, tt.wantStruct)
			}
		})
	}
}
func Test_adressObject_parse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		data       []byte
		wantStruct addressObject
	}{
		{
			[]byte(`{"id":1011883,"name":"Свердловская","aoid":"e76abf09-3148-42f6-85db-51edb09e72b7","guid":"92b30014-4d52-4e2e-892d-928142b924bf","postalCode":"620000","level":1,"okato":"65000000000","type":{"id":6,"name":"Область","shortName":"обл","code":"105","level":1},"regionCode":"66","autoCode":"0","areaCode":"000","cityCode":"000","ctarCode":"000","placeCode":"000","streetCode":"0000","extrCode":"0000","sextCode":"000","kladrCode":"6600000000000","live":true,"typeName":"Область","typeShortName":"обл","url":"/фиас/1011883-свердловская_область/","companyCount":362553,"fullName":"Свердловская область"}`),
			addressObject{1011883, "Область", "обл", "Свердловская область", "Свердловская", "e76abf09-3148-42f6-85db-51edb09e72b7", "92b30014-4d52-4e2e-892d-928142b924bf", "620000", 1, "65000000000", typeAddressObject{6, "Область", "обл", "105", 1}, "66", "0", "000", "000", "000", "000", "0000", "0000", "000", "6600000000000", true, "/фиас/1011883-свердловская_область/", 362553},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			var obj addressObject
			err := json.Unmarshal(tt.data, &obj)
			if err != nil {
				t.Fatalf("ошибка парсинга addressObject: %v", err)
			}
			if !reflect.DeepEqual(tt.wantStruct, obj) {
				t.Fatalf("addressObject == %v, want == %v", obj, tt.wantStruct)
			}
		})
	}
}
func Test_adress_parse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		data       []byte
		wantStruct address
	}{
		{
			[]byte(`{"region":{"id":1208752,"name":"Санкт-Петербург","aoid":"aad1469e-54ff-4605-af4f-f016c75b84d2","guid":"c2deb16a-0330-4f05-821f-1d09c93331e6","postalCode":"190000","level":1,"okato":"40000000000","type":{"id":4,"name":"Город","shortName":"г","code":"103","level":1},"regionCode":"78","autoCode":"0","areaCode":"000","cityCode":"000","ctarCode":"000","placeCode":"000","streetCode":"0000","extrCode":"0000","sextCode":"000","kladrCode":"7800000000000","live":true,"typeName":"Город","typeShortName":"г","url":"/фиас/1208752-санктпетербург_город/","companyCount":791135,"fullName":"Санкт-Петербург город"},"street":{"id":1209533,"name":"Дунайский","aoid":"f9bed8f3-48d4-4baf-81fa-86cb373134c3","guid":"5bf2e2fd-64e0-4d85-a50d-c62e8e5a7d42","level":7,"type":{"id":141,"name":"Проспект","shortName":"пр-кт","code":"719","level":7},"regionCode":"78","autoCode":"0","areaCode":"000","cityCode":"000","ctarCode":"000","placeCode":"000","streetCode":"0388","extrCode":"0000","sextCode":"000","kladrCode":"78000000000038800","live":true,"typeName":"Проспект","typeShortName":"пр-кт","url":"/фиас/1209533-проспект_дунайский/","companyCount":932,"fullName":"проспект Дунайский"},"house":"13","building":"2","flat":"ЛИТА","postalIndex":"196158","fiasOnlineLink":"https://фиас.онлайн//5bf2e2fd-64e0-4d85-a50d-c62e8e5a7d42/","fullAddress":"196158, Санкт-Петербург город, проспект Дунайский, 13 2, ЛИТА","fullHouseAddress":"196158, Санкт-Петербург город, проспект Дунайский, 13 2"}`),
			address{Region: addressObject{ID: 1208752, Name: "Санкт-Петербург", Aoid: "aad1469e-54ff-4605-af4f-f016c75b84d2", GUID: "c2deb16a-0330-4f05-821f-1d09c93331e6", PostalCode: "190000", Level: 1, Okato: "40000000000", Type: typeAddressObject{ID: 4, Name: "Город", ShortName: "г", Code: "103", Level: 1}, RegionCode: "78", AutoCode: "0", AreaCode: "000", CityCode: "000", CtarCode: "000", PlaceCode: "000", StreetCode: "0000", ExtrCode: "0000", SextCode: "000", KladrCode: "7800000000000", Live: true, TypeName: "Город", TypeShortName: "г", URL: "/фиас/1208752-санктпетербург_город/", CompanyCount: 791135, FullName: "Санкт-Петербург город"}, Street: addressObject{ID: 1209533, Name: "Дунайский", Aoid: "f9bed8f3-48d4-4baf-81fa-86cb373134c3", GUID: "5bf2e2fd-64e0-4d85-a50d-c62e8e5a7d42", Level: 7, Type: typeAddressObject{ID: 141, Name: "Проспект", ShortName: "пр-кт", Code: "719", Level: 7}, RegionCode: "78", AutoCode: "0", AreaCode: "000", CityCode: "000", CtarCode: "000", PlaceCode: "000", StreetCode: "0388", ExtrCode: "0000", SextCode: "000", KladrCode: "78000000000038800", Live: true, TypeName: "Проспект", TypeShortName: "пр-кт", URL: "/фиас/1209533-проспект_дунайский/", CompanyCount: 932, FullName: "проспект Дунайский"}, House: "13", Building: "2", Flat: "ЛИТА", PostalIndex: "196158", FiasOnlineLink: "https://фиас.онлайн//5bf2e2fd-64e0-4d85-a50d-c62e8e5a7d42/", FullAddress: "196158, Санкт-Петербург город, проспект Дунайский, 13 2, ЛИТА", FullHouseAddress: "196158, Санкт-Петербург город, проспект Дунайский, 13 2"},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i+1), func(t *testing.T) {
			var obj address
			err := json.Unmarshal(tt.data, &obj)
			if err != nil {
				t.Fatalf("ошибка парсинга address: %v", err)
			}
			if !reflect.DeepEqual(tt.wantStruct, obj) {
				t.Fatalf("address == %v, want == %v", obj, tt.wantStruct)
			}
		})
	}
}
