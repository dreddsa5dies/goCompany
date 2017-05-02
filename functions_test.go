package goCompany

import "testing"

func TestGetCompanyInfo(t *testing.T) {
	okOGRN := 5990130
	resultOGRN, err := GetCompanyInfo("7736002426")
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
	result, _ := GetIndivEntrep("7528374")
	if result[0].ID != ok {
		t.Fatalf("Want %v, but got %v", result, ok)
	}
}

func TestGetCompFounder(t *testing.T) {
	ok := 5545071
	result, _ := GetCompFounder("2191023")
	if result[0].ID != ok {
		t.Fatalf("Want %v, but got %v", result, ok)
	}
}

func TestGetDepComp(t *testing.T) {
	ok := 1425227
	result, _ := GetDepComp("7030")
	if result[0].ID != ok {
		t.Fatalf("Want %v, but got %v", result, ok)
	}
}

func TestGetPerson(t *testing.T) {
	ok := "АЛЕКСЕЙ"
	result, _ := GetPerson("2191023")
	if result.FirstName != ok {
		t.Fatalf("Want %v, but got %v", result, ok)
	}
}

func TestGetPositions(t *testing.T) {
	ok := 147863776
	result, _ := GetPositions("2191023")
	if result[0].ID != ok {
		t.Fatalf("Want %v, but got %v", result, ok)
	}
}

func TestGetIDData(t *testing.T) {
	ok := 7030
	result, _ := GetIDData("7030")
	if result.ID != ok {
		t.Fatalf("Want %v, but got %v", result, ok)
	}
}

func TestGetFounders(t *testing.T) {
	ok := 253175464
	result, _ := GetFounders("7030")
	if result[0].ID != ok {
		t.Fatalf("Want %v, but got %v", result, ok)
	}
}

func TestResponse(t *testing.T) {
	ok := []byte("[ ]")
	result := response("https://ru.rus.company/интеграция/компании/")
	if result[0] != ok[0] {
		t.Fatalf("Want %v, but got %v", result[0], ok[0])
	}
}
