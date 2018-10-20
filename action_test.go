package goCompany

import (
	"reflect"
	"testing"
)

func TestIsActive(t *testing.T) {
	tests := []struct {
		name string
		id   int
		want bool
	}{
		{"10104415:true", 10104415, true},
		{"6879538:false", 6879538, false},
		{"7030:true", 7030, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			company, err := GetCompanyFullInfo(tt.id)
			if err != nil {
				t.Skipf("IsActive: непредвиденная ошибка в запросе: %v", err)
			}
			if got := company.IsActive(); got != tt.want {
				t.Errorf("IsActive: got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Struct_SortByID(t *testing.T) {
	tests := []struct {
		name  string
		slice sorterByID
		want  sorterByID
	}{
		{"SliceCompanyInfo: сортровка не требуется", SliceCompanyInfo{CompanyInfo{ID: 1}, CompanyInfo{ID: 2}, CompanyInfo{ID: 3}}, SliceCompanyInfo{CompanyInfo{ID: 1}, CompanyInfo{ID: 2}, CompanyInfo{ID: 3}}},
		{"SlicePeopleInfo: сортровка не требуется", SlicePeopleInfo{PeopleInfo{ID: 1}, PeopleInfo{ID: 2}, PeopleInfo{ID: 3}}, SlicePeopleInfo{PeopleInfo{ID: 1}, PeopleInfo{ID: 2}, PeopleInfo{ID: 3}}},
		{"SliceCompanyInfo: сортровка требуется", SliceCompanyInfo{CompanyInfo{ID: 2}, CompanyInfo{ID: 3}, CompanyInfo{ID: 1}}, SliceCompanyInfo{CompanyInfo{ID: 1}, CompanyInfo{ID: 2}, CompanyInfo{ID: 3}}},
		{"SlicePeopleInfo: сортровка требуется", SlicePeopleInfo{PeopleInfo{ID: 3}, PeopleInfo{ID: 2}, PeopleInfo{ID: 1}}, SlicePeopleInfo{PeopleInfo{ID: 1}, PeopleInfo{ID: 2}, PeopleInfo{ID: 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slice.sortByID()
			if !reflect.DeepEqual(tt.slice, tt.want) {
				t.Errorf("sortByID: slice = %v, want %v", tt.slice, tt.want)
			}
		})
	}
}

func TestExtractOwners(t *testing.T) {
	owners, err := GetOwners(7030)
	if err != nil {
		t.Skipf("ExtractOwners: неожиданная ошибка %v", err)
	}

	companyOwners, peopleOwners := owners.ExtractOwners()

	if len(companyOwners) != 1 || companyOwners[0].ID != 10237185 {
		t.Errorf("ExtractOwners: len(companyOwners) = %d, want 1. ID = %d, want 10237185", len(companyOwners), companyOwners[0].ID)
	}

	if len(peopleOwners) != 5 || peopleOwners[2].ID != 2191025 {
		t.Errorf("ExtractOwners: len(peopleOwners) = %d, want 5. ID = %d, want 2191025", len(peopleOwners), peopleOwners[2].ID)
	}
}

func TestExtractCompany(t *testing.T) {
	jobs, err := GetJobs(2191023)
	if err != nil {
		t.Skipf("ExtractCompany: неожиданная ошибка %v", err)
	}

	slice := jobs.ExtractCompany()
	test := []int{10104415, 6879538, 7030, 7030, 10104415}

	for i := range slice {
		if slice[i].ID != test[i] {
			t.Error()
		}
	}
}

func TestExtractPeople(t *testing.T) {
	workers, err := GetAssociates(32357)
	if err != nil {
		t.Skipf("ExtractPeople: неожиданная ошибка %v", err)
	}

	peoples := workers.ExtractPeople()
	if len(peoples) != 1 || peoples[0].ID != 49758 {
		t.Error()
	}
}

func TestClearInactive(t *testing.T) {
	slice, err := GetSubCompanies(7030)
	if err != nil {
		t.Skipf("ClearInactive: неожиданная ошибка %v", err)
	}
	l := len(slice)

	slice.ClearInactive()
	if len(slice) != l-2 {
		t.Errorf("ClearInactive: len = %d, want %d", len(slice), l-2)
	}

	test := []int{7029, 1425227, 1710862, 2881529, 4721632, 8204373, 9929617, 10104415}

	for i := range slice {
		if slice[i].ID != test[i] {
			t.Errorf("ClearInactive: test %d, ID = %d, want %d", i, slice[i].ID, test[i])
		}
	}
}

func Test_Struct_ClearDouble(t *testing.T) {
	tests := []struct {
		name  string
		slice clearnerDouble
		want  clearnerDouble
	}{
		{
			"нет лишних структур",
			&SliceCompanyInfo{CompanyInfo{ID: 3}, CompanyInfo{ID: 2}, CompanyInfo{ID: 1}},
			&SliceCompanyInfo{CompanyInfo{ID: 1}, CompanyInfo{ID: 2}, CompanyInfo{ID: 3}},
		},
		{
			"есть лишние структуры",
			&SlicePeopleInfo{PeopleInfo{ID: 3}, PeopleInfo{ID: 2}, PeopleInfo{ID: 3}, PeopleInfo{ID: 1}},
			&SlicePeopleInfo{PeopleInfo{ID: 1}, PeopleInfo{ID: 2}, PeopleInfo{ID: 3}},
		},
		/*{
			"пустой слайс",
			&SlicePeopleInfo{},
			&SlicePeopleInfo{},
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slice.ClearDouble()
			if !reflect.DeepEqual(tt.slice, tt.want) {
				t.Errorf("ClearDouble: got = %v, want %v", tt.slice, tt.want)
			}
		})
	}
}
