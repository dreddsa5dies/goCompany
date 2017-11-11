package ogrnOnline

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_struct_CompanyFinanceInfo_parse(t *testing.T) {
	p := `./test_json/CompanyFinanceInfo.json`
	j, e := os.Open(p)
	if e != nil {
		t.Skipf("ошибка при открытии файла %s: %v", p, e)
	}
	defer j.Close()
	var obj []CompanyFinanceInfo
	e = json.NewDecoder(j).Decode(&obj)
	if e != nil {
		t.Errorf("ошибка парсинга: %v", e)
	}
}

func TestGenFinance(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	var (
		path = `./test_json/CompanyFinanceInfo.json`
		want = []CompanyFinanceInfo{}
	)
	res, err := GenFinance(8)
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
func Test_CompanyBaseInfo_GenFinance(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	var (
		path = `./test_json/CompanyFinanceInfo.json`
		want = []CompanyFinanceInfo{}
	)
	obj := CompanyBaseInfo{}
	obj.ID = 8
	res, err := obj.GetFinance()
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
func Test_CompanyInfo_GenFinance(t *testing.T) {
	time.Sleep(time.Millisecond * 400)
	var (
		path = `./test_json/CompanyFinanceInfo.json`
		want = []CompanyFinanceInfo{}
	)
	obj := CompanyInfo{}
	obj.ID = 8
	res, err := obj.GetFinance()
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
