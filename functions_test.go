package goCompany

import "testing"

func TestGetCompanyInfo(t *testing.T) {
	ok := 5990130
	result, _ := GetCompanyInfo("https://ru.rus.company/интеграция/компании/?инн=7736002426")
	if result[0].ID != ok {
		t.Fatalf("Want %v, but got %v", result, ok)
	}
}

/*
func TestGetEmployees(t *testing.T) {

}

func TestGetIndivEntrep(t *testing.T) {

}

func TestGetCompFounder(t *testing.T) {

}

func TestGetDepComp(t *testing.T) {

}

func TestGetPerson(t *testing.T) {

}

func TestGetPositions(t *testing.T) {

}

func TestGetIDData(t *testing.T) {

}

func TestGetFounders(t *testing.T) {

}

func Testresponse(t *testing.T) {

}
*/
