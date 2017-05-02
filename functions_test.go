package goCompany

import "testing"

func TestGetCompanyInfo(t *testing.T) {
	okOGRN := 5990130
	resultOGRN, err := GetCompanyInfo("1027700416699")
	if err != nil {
		t.Fatal(err)
	}
	if resultOGRN[0].ID != okOGRN {
		t.Fatalf("Want %v, but got %v", resultOGRN, okOGRN)
	}

	okNAME := 731085
	resultNAME, err := GetCompanyInfo("СТРОЙПРОЕКТ")
	if err != nil {
		t.Fatal(err)
	}
	if resultNAME[0].ID != okNAME {
		t.Fatalf("Want %v, but got %v", resultNAME[0].ID, okNAME)
	}

	ok := 212722
	result, err := GetCompanyInfo("1658064460")
	if err != nil {
		t.Fatal(err)
	}
	if result[0].ID != ok {
		t.Fatalf("Want %v, but got %v", result[0].ID, ok)
	}
}

func TestGetEmployees(t *testing.T) {
	ok := 149735376
	result, err := GetEmployees("32357")
	if err != nil {
		t.Fatal(err)
	}
	if result[0].ID != ok {
		t.Fatalf("Want %v, but got %v", result, ok)
	}
}

func TestGetIndivEntrep(t *testing.T) {
	ok := 1
	result, err := GetIndivEntrep("7528374")
	if err != nil {
		t.Fatal(err)
	}
	if result[0].ID != ok {
		t.Fatalf("Want %v, but got %v", result[0].ID, ok)
	}

	okOGRNIP := 1
	resultOGRNIP, err := GetIndivEntrep("314272211800010")
	if err != nil {
		t.Fatal(err)
	}
	if resultOGRNIP[0].ID != okOGRNIP {
		t.Fatalf("Want %v, but got %v", resultOGRNIP[0].ID, okOGRNIP)
	}

	okINN := 1
	resultINN, err := GetIndivEntrep("272508402480")
	if err != nil {
		t.Fatal(err)
	}
	if resultINN[0].ID != ok {
		t.Fatalf("Want %v, but got %v", resultINN[0].ID, okINN)
	}
}

func TestGetCompFounder(t *testing.T) {
	ok := 5545071
	result, err := GetCompFounder("2191023")
	if err != nil {
		t.Fatal(err)
	}
	if result[0].ID != ok {
		t.Fatalf("Want %v, but got %v", result[0].ID, ok)
	}
}

func TestGetDepComp(t *testing.T) {
	ok := 1425227
	result, err := GetDepComp("7030")
	if err != nil {
		t.Fatal(err)
	}
	if result[0].ID != ok {
		t.Fatalf("Want %v, but got %v", result[0].ID, ok)
	}
}

func TestGetPerson(t *testing.T) {
	ok := "АЛЕКСЕЙ"
	result, err := GetPerson("2191023")
	if err != nil {
		t.Fatal(err)
	}
	if result.FirstName != ok {
		t.Fatalf("Want %v, but got %v", result, ok)
	}
}

func TestGetPositions(t *testing.T) {
	ok := 147863776
	result, err := GetPositions("2191023")
	if err != nil {
		t.Fatal(err)
	}
	if result[0].ID != ok {
		t.Fatalf("Want %v, but got %v", result[0].ID, ok)
	}
}

func TestGetIDData(t *testing.T) {
	ok := 7030
	result, err := GetIDData("7030")
	if err != nil {
		t.Fatal(err)
	}
	if result.ID != ok {
		t.Fatalf("Want %v, but got %v", result, ok)
	}
}

func TestGetFounders(t *testing.T) {
	ok := 253175464
	result, err := GetFounders("7030")
	if err != nil {
		t.Fatal(err)
	}
	if result[0].ID != ok {
		t.Fatalf("Want %v, but got %v", result[0].ID, ok)
	}
}

func TestGetSearchPerson(t *testing.T) {
	ok := 2191023
	result, err := GetSearchPerson("781010102204")
	if err != nil {
		t.Fatal(err)
	}
	if result[0].ID != ok {
		t.Fatalf("Want %v, but got %v", result[0].ID, ok)
	}

	okF := 14340350
	resultF, err := GetSearchPerson("ЖУРБИН")
	if err != nil {
		t.Fatal(err)
	}
	if resultF[0].ID != okF {
		t.Fatalf("Want %v, but got %v", resultF[0].ID, okF)
	}

	okFI := 2191023
	resultFI, err := GetSearchPerson("ЖУРБИН Алексей")
	if err != nil {
		t.Fatal(err)
	}
	if resultFI[0].ID != okFI {
		t.Fatalf("Want %v, but got %v", resultFI[0].ID, okFI)
	}

	okFIO := 473312
	resultFIO, err := GetSearchPerson("ЖУРБИН Алексей Николаевич")
	if err != nil {
		t.Fatal(err)
	}
	if resultFIO[0].ID != okFIO {
		t.Fatalf("Want %v, but got %v", resultFI[0].ID, okFIO)
	}
}

func TestResponse(t *testing.T) {
	ok := []byte("[ ]")
	result := response("https://ru.rus.company/интеграция/компании/")
	if result[0] != ok[0] {
		t.Fatalf("Want %v, but got %v", result[0], ok[0])
	}
}
