package ogrnOnline

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_getDataFromServer(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	tests := []struct {
		num int
		url string
	}{
		{1, `https://xn--c1aubj.xn--80asehdb/%D0%B8%D0%BD%D1%82%D0%B5%D0%B3%D1%80%D0%B0%D1%86%D0%B8%D1%8F/%D0%BB%D1%8E%D0%B4%D0%B8/?%D1%84%D0%B0%D0%BC%D0%B8%D0%BB%D0%B8%D1%8F=%D1%86%D0%B8%D0%BF%D0%BE%D1%80%D0%B8%D0%BD&%D0%B8%D0%BC%D1%8F=%D0%B0%D0%BB%D0%B5%D0%BA%D1%81%D0%B0%D0%BD%D0%B4%D1%80`},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.num), func(t *testing.T) {
			if resp := getDataFromServer(tt.url); resp == nil {
				t.Error("resp == nil")
			}
		})
	}
}
func Test_createURL(t *testing.T) {
	t.Parallel()
	type args struct {
		path  string
		query url.Values
	}
	tests := []struct {
		args args
		want *url.URL
	}{
		{
			args{"/интеграция/компании/7030/учредители/", nil},
			func() *url.URL {
				ur, _ := url.Parse(host)
				ur.Path = "/интеграция/компании/7030/учредители/"
				return ur
			}(),
		},
		{
			args{"/интеграция/компании/", url.Values{"инн": []string{"7736002426"}}},
			func() *url.URL {
				ur, _ := url.Parse(host)
				ur.Path = "/интеграция/компании/"
				ur.RawQuery = url.Values{"инн": []string{"7736002426"}}.Encode()
				return ur
			}(),
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if got := createURL(tt.args.path, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_isValidQuery(t *testing.T) {
	t.Parallel()
	type args struct {
		u url.Values
		t int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"typeQueryCompany. ok", args{url.Values{"огрн": []string{"0000000000000"}}, typeQueryCompany}, false},
		{"typeQueryCompany. err_1", args{url.Values{"о1грн": []string{""}}, typeQueryCompany}, true},
		{"typeQueryCompany. err_2", args{url.Values{"огрн": []string{""}, "test": []string{}}, typeQueryCompany}, true},
		{"typeQueryPeople. ok", args{url.Values{"инн": []string{"000000000000"}}, typeQueryPeople}, false},
		{"typeQueryPeople. err_1", args{url.Values{"о1грн": []string{""}}, typeQueryPeople}, true},
		{"typeQueryPeople. err_2", args{url.Values{"огрн": []string{""}, "test": []string{}}, typeQueryPeople}, true},
		{"typeQueryBusinessman. ok", args{url.Values{"огрнип": []string{"000000000000000"}}, typeQueryBusinessman}, false},
		{"typeQueryBusinessman. err_1", args{url.Values{"о1грн": []string{""}}, typeQueryBusinessman}, true},
		{"typeQueryBusinessman. err_2", args{url.Values{"огрн": []string{""}, "test": []string{}}, typeQueryBusinessman}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := isValidQuery(tt.args.u, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("err:= isValidQuery() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
func TestFindCompany(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	var (
		path  = `./test_json/CompanyBaseInfo.json`
		query = url.Values{"инн": []string{"7707083893"}}
		want  = []CompanyBaseInfo{}
	)
	res, err := FindCompany(query)
	if err != nil {
		t.Errorf("ошибка функции: %v", err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Skipf("ошибка открытия файла: %v", err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&want)
	if err != nil {
		t.Skipf("ошибка парсинга файла: %v", err)
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("res == %v, want %v", res, want)
	}
}
func TestFindPeople(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	var (
		path  = `./test_json/PeopleInfo.json`
		query = url.Values{"фамилия": []string{"ЖУРБИН"}, "имя": []string{"АЛЕКСЕЙ"}}
		want  = []PeopleInfo{}
	)
	res, err := FindPeople(query)
	if err != nil {
		t.Errorf("ошибка функции: %v", err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Skipf("ошибка открытия файла: %v", err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&want)
	if err != nil {
		t.Skipf("ошибка парсинга файла: %v", err)
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("res == %v, want %v", res, want)
	}
}
func TestFindBusinessman(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	var (
		path  = `./test_json/PeopleBusinessmanInfo.json`
		query = url.Values{"огрнип": []string{"314272211800010"}}
		want  = []PeopleBusinessmanInfo{}
	)
	res, err := FindBusinessman(query)
	if err != nil {
		t.Errorf("ошибка функции: %v", err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Skipf("ошибка открытия файла: %v", err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&want)
	if err != nil {
		t.Skipf("ошибка парсинга файла: %v", err)
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("res == %v, want %v", res, want)
	}
}
func TestGetCompany(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	var (
		path = `./test_json/CompanyInfo.json`
		want = CompanyInfo{}
	)
	res, err := GetCompany(7030)
	if err != nil {
		t.Errorf("ошибка функции: %v", err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Skipf("ошибка открытия файла: %v", err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&want)
	if err != nil {
		t.Skipf("ошибка парсинга файла: %v", err)
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("res == %v, want %v", res, want)
	}
}
func TestGetOwners(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	var (
		path = `./test_json/CompanyOwnersInfo.json`
		want = []CompanyOwnerInfo{}
	)
	res, err := GetOwners(7030)
	if err != nil {
		t.Errorf("ошибка функции: %v", err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Skipf("ошибка открытия файла: %v", err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&want)
	if err != nil {
		t.Skipf("ошибка парсинга файла: %v", err)
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("res == %v, want %v", res, want)
	}
}
func TestGetAssociates(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	var (
		path = `./test_json/CompanyAssociateInfo.json`
		want = []CompanyAssociateInfo{}
	)
	res, err := GetAssociates(32357)
	if err != nil {
		t.Errorf("ошибка функции: %v", err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Skipf("ошибка открытия файла: %v", err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&want)
	if err != nil {
		t.Skipf("ошибка парсинга файла: %v", err)
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("res == %v, want %v", res, want)
	}
}
func TestGetSubCompanies(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	var (
		path = `./test_json/CompanySubCompaniesInfo.json`
		want = []CompanyBaseInfo{}
	)
	res, err := GetSubCompanies(7030)
	if err != nil {
		t.Errorf("ошибка функции: %v", err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Skipf("ошибка открытия файла: %v", err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&want)
	if err != nil {
		t.Skipf("ошибка парсинга файла: %v", err)
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("res == %v, want %v", res, want)
	}
}
func TestGetRepresentativeOffices(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	var (
		path = `./test_json/CompanyRepresentativeOfficesInfo.json`
		want = []CompanyBranchInfo{}
	)
	res, err := GetRepresentativeOffices(4527642)
	if err != nil {
		t.Errorf("ошибка функции: %v", err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Skipf("ошибка открытия файла: %v", err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&want)
	if err != nil {
		t.Skipf("ошибка парсинга файла: %v", err)
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("res == %v, want %v", res, want)
	}
}
func TestGetBranches(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	var (
		path = `./test_json/CompanyBranchesInfo.json`
		want = []CompanyBranchInfo{}
	)
	res, err := GetBranches(4527642)
	if err != nil {
		t.Errorf("ошибка функции: %v", err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Skipf("ошибка открытия файла: %v", err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&want)
	if err != nil {
		t.Skipf("ошибка парсинга файла: %v", err)
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("res == %v, want %v", res, want)
	}
}
func TestGetPeople(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	var (
		path = `./test_json/PeopleInfo.json`
		want = PeopleInfo{}
	)
	res, err := GetPeople(2191023)
	if err != nil {
		t.Errorf("ошибка функции: %v", err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Skipf("ошибка открытия файла: %v", err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&want)
	if err != nil {
		t.Skipf("ошибка парсинга файла: %v", err)
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("res == %v, want %v", res, want)
	}
}
func TestGetJobs(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	var (
		path = `./test_json/PeopleJobsInfo.json`
		want = []CompanyAssociateInfo{}
	)
	res, err := GetJobs(2191023)
	if err != nil {
		t.Errorf("ошибка функции: %v", err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Skipf("ошибка открытия файла: %v", err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&want)
	if err != nil {
		t.Skipf("ошибка парсинга файла: %v", err)
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("res == %v, want %v", res, want)
	}
}
func TestGetShare(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	var (
		path = `./test_json/PeopleShareInfo.json`
		want = []CompanyBaseInfo{}
	)
	res, err := GetShare(2191023)
	if err != nil {
		t.Errorf("ошибка функции: %v", err)
	}
	file, err := os.Open(path)
	if err != nil {
		t.Skipf("ошибка открытия файла: %v", err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&want)
	if err != nil {
		t.Skipf("ошибка парсинга файла: %v", err)
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("res == %v, want %v", res, want)
	}
}
