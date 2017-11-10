package ogrnOnline

import (
	"encoding/json"
	"os"
	"testing"
)

func Test_struct_CompanyBaseInfo_parse(t *testing.T) {
	p := `./test_json/CompanyBaseInfo.json`
	j, e := os.Open(p)
	if e != nil {
		t.Skipf("ошибка при открытии файла %s: %v", p, e)
	}
	defer j.Close()
	var obj []CompanyBaseInfo
	e = json.NewDecoder(j).Decode(&obj)
	if e != nil {
		t.Errorf("ошибка парсинга: %v", e)
	}
}
func Test_struct_CompanyInfo_parse(t *testing.T) {
	p := `./test_json/CompanyInfo.json`
	j, e := os.Open(p)
	if e != nil {
		t.Skipf("ошибка при открытии файла %s: %v", p, e)
	}
	defer j.Close()
	var obj CompanyInfo
	e = json.NewDecoder(j).Decode(&obj)
	if e != nil {
		t.Errorf("ошибка парсинга: %v", e)
	}
}
func Test_struct_CompanyOwnersInfo_parse(t *testing.T) {
	p := `./test_json/CompanyOwnersInfo.json`
	j, e := os.Open(p)
	if e != nil {
		t.Skipf("ошибка при открытии файла %s: %v", p, e)
	}
	defer j.Close()
	var obj []CompanyOwnerInfo
	e = json.NewDecoder(j).Decode(&obj)
	if e != nil {
		t.Errorf("ошибка парсинга: %v", e)
	}
}
func Test_struct_CompanyAssociateInfo_parse(t *testing.T) {
	p := `./test_json/CompanyAssociateInfo.json`
	j, e := os.Open(p)
	if e != nil {
		t.Skipf("ошибка при открытии файла %s: %v", p, e)
	}
	defer j.Close()
	var obj []CompanyAssociateInfo
	e = json.NewDecoder(j).Decode(&obj)
	if e != nil {
		t.Errorf("ошибка парсинга: %v", e)
	}
}
func Test_struct_CompanySubCompaniesInfo_parse(t *testing.T) {
	p := `./test_json/CompanySubCompaniesInfo.json`
	j, e := os.Open(p)
	if e != nil {
		t.Skipf("ошибка при открытии файла %s: %v", p, e)
	}
	defer j.Close()
	var obj []CompanyInfo
	e = json.NewDecoder(j).Decode(&obj)
	if e != nil {
		t.Errorf("ошибка парсинга: %v", e)
	}
}
func Test_struct_CompanyBranchesInfo_parse(t *testing.T) {
	p := `./test_json/CompanyBranchesInfo.json`
	j, e := os.Open(p)
	if e != nil {
		t.Skipf("ошибка при открытии файла %s: %v", p, e)
	}
	defer j.Close()
	var obj []CompanyBranchInfo
	e = json.NewDecoder(j).Decode(&obj)
	if e != nil {
		t.Errorf("ошибка парсинга: %v", e)
	}
}
func Test_struct_PeopleInfo_parse(t *testing.T) {
	p := `./test_json/PeopleInfo.json`
	j, e := os.Open(p)
	if e != nil {
		t.Skipf("ошибка при открытии файла %s: %v", p, e)
	}
	defer j.Close()
	var obj PeopleInfo
	e = json.NewDecoder(j).Decode(&obj)
	if e != nil {
		t.Errorf("ошибка парсинга: %v", e)
	}
}
func Test_struct_PeopleJobsInfo_parse(t *testing.T) {
	p := `./test_json/PeopleJobsInfo.json`
	j, e := os.Open(p)
	if e != nil {
		t.Skipf("ошибка при открытии файла %s: %v", p, e)
	}
	defer j.Close()
	var obj []CompanyAssociateInfo
	e = json.NewDecoder(j).Decode(&obj)
	if e != nil {
		t.Errorf("ошибка парсинга: %v", e)
	}
}
func Test_struct_PeopleShareInfo_parse(t *testing.T) {
	p := `./test_json/PeopleShareInfo.json`
	j, e := os.Open(p)
	if e != nil {
		t.Skipf("ошибка при открытии файла %s: %v", p, e)
	}
	defer j.Close()
	var obj []CompanyOwnerInfo
	e = json.NewDecoder(j).Decode(&obj)
	if e != nil {
		t.Errorf("ошибка парсинга: %v", e)
	}
}
func Test_struct_PeopleBusinessmanInfo_parse(t *testing.T) {
	p := `./test_json/PeopleBusinessmanInfo.json`
	j, e := os.Open(p)
	if e != nil {
		t.Skipf("ошибка при открытии файла %s: %v", p, e)
	}
	defer j.Close()
	var obj []PeopleBusinessmanInfo
	e = json.NewDecoder(j).Decode(&obj)
	if e != nil {
		t.Errorf("ошибка парсинга: %v", e)
	}
}
