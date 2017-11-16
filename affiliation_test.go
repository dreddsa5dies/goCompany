package ogrnOnline

import "testing"
import "reflect"

func TestInterface(t *testing.T) {
	var c interface{} = new(CompanyBaseInfo)
	var p interface{} = new(PeopleInfo)

	if _, ok := c.(Node); !ok {
		t.Error()
	}
	if _, ok := p.(Node); !ok {
		t.Error()
	}
}

func TestGetID(t *testing.T) {
	tests := []struct {
		name string
		obj  Node
		id   int
	}{
		{"CompanyBaseInfo", &CompanyBaseInfo{ID: 1}, 1},
		{"PeopleInfo", &PeopleInfo{ID: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if id := tt.obj.GetID(); id != tt.id {
				t.Errorf("GetID() = %v, want %v", id, tt.id)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name string
		obj  Node
		prn  string
	}{
		{"CompanyBaseInfo", &CompanyBaseInfo{Name: "CompanyBaseInfo", INN: "123456789"}, "CompanyBaseInfo (ИНН 123456789)"},
		{"PeopleInfo", &PeopleInfo{FullName: "PeopleInfo", INN: "123456789"}, "PeopleInfo (ИНН 123456789)"},

		{"CompanyBaseInfo", &CompanyBaseInfo{Name: "CompanyBaseInfo", INN: ""}, "CompanyBaseInfo (ИНН неизвестен)"},
		{"PeopleInfo", &PeopleInfo{FullName: "PeopleInfo", INN: ""}, "PeopleInfo (ИНН неизвестен)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if prn := tt.obj.String(); prn != tt.prn {
				t.Errorf("String() = %v, want %v", prn, tt.prn)
			}
		})
	}
}

func TestSortNode(t *testing.T) {
	tests := []struct {
		name  string
		slice []Node
		want  []Node
	}{
		{
			"сортировка не требуется",
			[]Node{&CompanyBaseInfo{ID: 1}, &PeopleInfo{ID: 2}, &CompanyBaseInfo{ID: 3}, &PeopleInfo{ID: 4}, &PeopleInfo{ID: 5}},
			[]Node{&CompanyBaseInfo{ID: 1}, &PeopleInfo{ID: 2}, &CompanyBaseInfo{ID: 3}, &PeopleInfo{ID: 4}, &PeopleInfo{ID: 5}},
		},
		{
			"сортировка требуется",
			[]Node{&PeopleInfo{ID: 2}, &CompanyBaseInfo{ID: 3}, &CompanyBaseInfo{ID: 1}, &PeopleInfo{ID: 5}, &PeopleInfo{ID: 4}},
			[]Node{&CompanyBaseInfo{ID: 1}, &PeopleInfo{ID: 2}, &CompanyBaseInfo{ID: 3}, &PeopleInfo{ID: 4}, &PeopleInfo{ID: 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortNode(tt.slice)
			if !reflect.DeepEqual(tt.slice, tt.want) {
				t.Errorf("sortNode(): got %v, want %v", tt.slice, tt.want)
			}
		})
	}
}

func TestClearSlice(t *testing.T) {
	tests := []struct {
		name  string
		slice []Node
		want  []Node
	}{
		{
			"clear",
			[]Node{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 1}, &PeopleInfo{ID: 2}, &PeopleInfo{ID: 3}, &PeopleInfo{ID: 3}, &PeopleInfo{ID: 4}, &PeopleInfo{ID: 5}},
			[]Node{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 2}, &PeopleInfo{ID: 3}, &PeopleInfo{ID: 4}, &PeopleInfo{ID: 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clearSlice(tt.slice)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clearSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBranch(t *testing.T) {
	tests := []struct {
		name      string
		obj       Node
		wantLen   int
		testIndex int
		testValue int
	}{
		{"гуп", &CompanyBaseInfo{ID: 1198655}, 7, 6, 8302065},
		{"захарова", &PeopleInfo{ID: 8302065}, 2, 0, 1198654},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			branch, err := tt.obj.GetBranch()
			if err != nil {
				t.Skipf("GetBranch(): %v", err)
			}
			if l := len(branch); l != tt.wantLen {
				t.Fatalf("GetBranch(): len = %v, want %v", l, tt.wantLen)
			}
			if branch[tt.testIndex].GetID() != tt.testValue {
				t.Errorf("GetBranch(): testValue = %v, want %v", branch[tt.testIndex].GetID(), tt.testValue)
			}
		})
	}
}
