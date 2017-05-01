package goCompany

import "testing"

func TestGetCompanyInfo(t *testing.T) {
	ok := 5990130
	result, _ := GetCompanyInfo("7736002426")
	if result[0].ID != ok {
		t.Fatalf("Want %v, but got %v", result, ok)
	}
}

func TestGetEmployees(t *testing.T) {
	ok := 149735376
	result, _ := GetEmployees("32357")
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
