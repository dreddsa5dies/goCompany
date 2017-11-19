package ogrnOnline

import (
	"reflect"
	"testing"
)

func TestInterface(t *testing.T) {
	var c interface{} = new(CompanyInfo)
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
		{"CompanyInfo", &CompanyInfo{ID: 1}, 1},
		{"PeopleInfo", &PeopleInfo{ID: 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if id := tt.obj.getID(); id != tt.id {
				t.Errorf("getID: obj.ID = %v, want %v", id, tt.id)
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
		{"CompanyInfo", &CompanyInfo{Name: "CompanyInfo", INN: "123456789"}, "[ CompanyInfo : ИНН 123456789 ]"},
		{"PeopleInfo", &PeopleInfo{FullName: "PeopleInfo", INN: "123456789"}, "[ PeopleInfo : ИНН 123456789 ]"},

		{"CompanyInfo", &CompanyInfo{Name: "CompanyInfo", INN: ""}, "[ CompanyInfo : ИНН неизвестен ]"},
		{"PeopleInfo", &PeopleInfo{FullName: "PeopleInfo", INN: ""}, "[ PeopleInfo : ИНН неизвестен ]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if prn := tt.obj.String(); prn != tt.prn {
				t.Errorf("String: %v, want %v", prn, tt.prn)
			}
		})
	}
}

func Test_Node_SortByID(t *testing.T) {
	tests := []struct {
		name  string
		slice Branch
		want  Branch
	}{
		{
			"сортировка не требуется",
			Branch{&CompanyInfo{ID: 1}, &PeopleInfo{ID: 2}, &CompanyInfo{ID: 3}, &PeopleInfo{ID: 4}, &PeopleInfo{ID: 5}},
			Branch{&CompanyInfo{ID: 1}, &PeopleInfo{ID: 2}, &CompanyInfo{ID: 3}, &PeopleInfo{ID: 4}, &PeopleInfo{ID: 5}},
		},
		{
			"сортировка требуется",
			Branch{&PeopleInfo{ID: 2}, &CompanyInfo{ID: 3}, &CompanyInfo{ID: 1}, &PeopleInfo{ID: 5}, &PeopleInfo{ID: 4}},
			Branch{&CompanyInfo{ID: 1}, &PeopleInfo{ID: 2}, &CompanyInfo{ID: 3}, &PeopleInfo{ID: 4}, &PeopleInfo{ID: 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slice.sortByID()
			if !reflect.DeepEqual(tt.slice, tt.want) {
				t.Errorf("sortNode: got = %v, want %v", tt.slice, tt.want)
			}
		})
	}
}

func TestClearDouble(t *testing.T) {
	tests := []struct {
		name  string
		slice Branch
		want  Branch
	}{
		{
			"clear",
			Branch{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 1}, &PeopleInfo{ID: 2}, &PeopleInfo{ID: 3}, &PeopleInfo{ID: 3}, &PeopleInfo{ID: 4}, &PeopleInfo{ID: 5}},
			Branch{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 2}, &PeopleInfo{ID: 3}, &PeopleInfo{ID: 4}, &PeopleInfo{ID: 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.slice.ClearDouble()
			if !reflect.DeepEqual(tt.slice, tt.want) {
				t.Errorf("clearSlice: got = %v, want %v", tt.slice, tt.want)
			}
		})
	}
}

func TestIsBankrupt(t *testing.T) {
	tests := []struct {
		name string
		post string
		want bool
	}{
		{"арбитражный", "Арбитражный управляющий", true},
		{"временный", "ВРЕМЕННЫЙ УПРАВЛЯЮЩИЙ", true},
		{"внешний", "внешний управляющий", true},
		{"конкурсный", "конкурсный УПРАВЛЯЮЩИЙ", true},
		{"директор", "директор", false},
		{"генеральный директор", "генеральный директор", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if b := isBankrupt(tt.post); b != tt.want {
				t.Errorf("isBankrupt: got = %v, want %v", b, tt.want)
			}
		})
	}
}

func TestIsCommercial(t *testing.T) {
	tests := []struct {
		name  string
		okopf string
		want  bool
	}{
		{"АО", "12267", true},
		{"ГУП", "65242", true},
		{"ООО", "12300", true},
		{"объединения работодателей", "20612", false},
		{"пустая строка", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if b := isCommercial(tt.okopf); b != tt.want {
				t.Errorf("isCommercial: got = %v, want %v", b, tt.want)
			}
		})
	}
}

func TestGetBranch(t *testing.T) {
	tests := []struct {
		name    string
		obj     Node
		wantLen int
	}{
		{"гуп", &CompanyInfo{ID: 1198655}, 5},
		{"захарова", &PeopleInfo{ID: 8302065}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			branch, err := tt.obj.getBranch()
			if err != nil {
				t.Skipf("getBranch: %v", err)
			}
			if l := len(branch); l != tt.wantLen {
				t.Errorf("getBranch: len = %v, want %v", l, tt.wantLen)
			}
		})
	}
}

func TestConnectionIsExist(t *testing.T) {
	graph := &Graph{
		1: Branch{&PeopleInfo{ID: 2}, &PeopleInfo{ID: 3}, &PeopleInfo{ID: 4}},
		2: Branch{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 5}},
		3: Branch{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 5}, &PeopleInfo{ID: 6}},
		4: Branch{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 6}},
		5: Branch{&PeopleInfo{ID: 2}, &PeopleInfo{ID: 3}, &PeopleInfo{ID: 7}},
		6: Branch{&PeopleInfo{ID: 3}, &PeopleInfo{ID: 4}},
		7: Branch{&PeopleInfo{ID: 5}},
		8: Branch{},
	}
	tests := []struct {
		name   string
		start  int
		finish int
		want   bool
	}{
		{"1 -> 7", 1, 7, true},
		{"5 -> 1", 5, 1, true},
		{"2 -> 4", 2, 4, true},
		{"1 -> 8", 1, 8, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := &PeopleInfo{ID: tt.start}
			finish := &PeopleInfo{ID: tt.finish}
			if connect := start.ConnectionIsExist(graph, finish); connect != tt.want {
				t.Errorf("ConnectionIsExist: %d -> %d == %v, want %v", tt.start, tt.finish, connect, tt.want)
			}
		})
	}
}

func TestFindAllConnections(t *testing.T) {
	graph := &Graph{
		1:  Branch{&PeopleInfo{ID: 2}, &PeopleInfo{ID: 3}, &PeopleInfo{ID: 4}, &PeopleInfo{ID: 5}},
		2:  Branch{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 9}},
		3:  Branch{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 4}, &PeopleInfo{ID: 6}, &PeopleInfo{ID: 7}},
		4:  Branch{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 3}},
		5:  Branch{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 6}, &PeopleInfo{ID: 8}},
		6:  Branch{&PeopleInfo{ID: 3}, &PeopleInfo{ID: 5}},
		7:  Branch{&PeopleInfo{ID: 3}, &PeopleInfo{ID: 12}},
		8:  Branch{&PeopleInfo{ID: 5}},
		9:  Branch{&PeopleInfo{ID: 2}, &PeopleInfo{ID: 10}, &PeopleInfo{ID: 11}},
		10: Branch{&PeopleInfo{ID: 9}},
		11: Branch{&PeopleInfo{ID: 9}},
		12: Branch{&PeopleInfo{ID: 7}},
	}
	tests := []struct {
		name   string
		start  int
		finish int
		want   int
	}{
		{"1 -> 7", 1, 7, 3},
		{"5 -> 1", 5, 1, 3},
		{"2 -> 4", 2, 4, 3},
		{"1 -> 8", 1, 8, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := &PeopleInfo{ID: tt.start}
			finish := &PeopleInfo{ID: tt.finish}
			connections := start.FindAllConnections(graph, finish)
			if l := len(connections); l != tt.want {
				t.Errorf("FindAllConnections: %d -> %d len = %v, want %v", tt.start, tt.finish, l, tt.want)
			}
		})
	}
}
