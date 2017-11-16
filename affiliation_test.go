package ogrnOnline

import "testing"
import "reflect"

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

func TestSortNode(t *testing.T) {
	tests := []struct {
		name  string
		slice []Node
		want  []Node
	}{
		{
			"сортировка не требуется",
			[]Node{&CompanyInfo{ID: 1}, &PeopleInfo{ID: 2}, &CompanyInfo{ID: 3}, &PeopleInfo{ID: 4}, &PeopleInfo{ID: 5}},
			[]Node{&CompanyInfo{ID: 1}, &PeopleInfo{ID: 2}, &CompanyInfo{ID: 3}, &PeopleInfo{ID: 4}, &PeopleInfo{ID: 5}},
		},
		{
			"сортировка требуется",
			[]Node{&PeopleInfo{ID: 2}, &CompanyInfo{ID: 3}, &CompanyInfo{ID: 1}, &PeopleInfo{ID: 5}, &PeopleInfo{ID: 4}},
			[]Node{&CompanyInfo{ID: 1}, &PeopleInfo{ID: 2}, &CompanyInfo{ID: 3}, &PeopleInfo{ID: 4}, &PeopleInfo{ID: 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortNode(tt.slice)
			if !reflect.DeepEqual(tt.slice, tt.want) {
				t.Errorf("sortNode: got = %v, want %v", tt.slice, tt.want)
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
				t.Errorf("clearSlice: got = %v, want %v", got, tt.want)
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
		{"гуп", &CompanyInfo{ID: 1198655}, 5, 4, 8302065},
		{"захарова", &PeopleInfo{ID: 8302065}, 2, 0, 1198654},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			branch, err := tt.obj.getBranch()
			if err != nil {
				t.Skipf("getBranch: %v", err)
			}
			if l := len(branch); l != tt.wantLen {
				t.Fatalf("getBranch: len = %v, want %v", l, tt.wantLen)
			}
			if branch[tt.testIndex].getID() != tt.testValue {
				t.Errorf("getBranch: testValue = %v, want %v", branch[tt.testIndex].getID(), tt.testValue)
			}
		})
	}
}

func TestFindConnectionBetweenNode(t *testing.T) {
	graph := &Graph{
		1: []Node{&PeopleInfo{ID: 2}, &PeopleInfo{ID: 3}, &PeopleInfo{ID: 4}},
		2: []Node{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 5}},
		3: []Node{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 5}, &PeopleInfo{ID: 6}},
		4: []Node{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 6}},
		5: []Node{&PeopleInfo{ID: 2}, &PeopleInfo{ID: 3}, &PeopleInfo{ID: 7}},
		6: []Node{&PeopleInfo{ID: 3}, &PeopleInfo{ID: 4}},
		7: []Node{&PeopleInfo{ID: 5}},
		8: []Node{},
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
			connection := findConnectionBetweenNode(graph, &PeopleInfo{ID: tt.start}, &PeopleInfo{ID: tt.finish}, []Node{})
			if l := (len(connection) != 0); l != tt.want {
				t.Errorf("findConnectionBetweenNode: %d -> %d == %v, want %v", tt.start, tt.finish, l, tt.want)
			}
		})
	}
}

func TestFindConnection(t *testing.T) {
	graph := &Graph{
		1: []Node{&PeopleInfo{ID: 2}, &PeopleInfo{ID: 3}, &PeopleInfo{ID: 4}},
		2: []Node{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 5}},
		3: []Node{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 5}, &PeopleInfo{ID: 6}},
		4: []Node{&PeopleInfo{ID: 1}, &PeopleInfo{ID: 6}},
		5: []Node{&PeopleInfo{ID: 2}, &PeopleInfo{ID: 3}, &PeopleInfo{ID: 7}},
		6: []Node{&PeopleInfo{ID: 3}, &PeopleInfo{ID: 4}},
		7: []Node{&PeopleInfo{ID: 5}},
		8: []Node{},
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
			connection := start.FindConnection(graph, finish)
			if l := (len(connection) != 0); l != tt.want {
				t.Errorf("FindConnection: %d -> %d == %v, want %v", tt.start, tt.finish, l, tt.want)
			}
		})
	}
}
