package ogrnOnline

import (
	"reflect"
	"testing"
)

func TestCompanyBaseInfo_GetID(t *testing.T) {
	tests := []struct {
		name string
		c    *CompanyBaseInfo
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetID(); got != tt.want {
				t.Errorf("CompanyBaseInfo.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyBaseInfo_String(t *testing.T) {
	tests := []struct {
		name string
		c    *CompanyBaseInfo
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("CompanyBaseInfo.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeopleInfo_String(t *testing.T) {
	tests := []struct {
		name string
		p    *PeopleInfo
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("PeopleInfo.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeopleInfo_GetID(t *testing.T) {
	tests := []struct {
		name string
		p    *PeopleInfo
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.GetID(); got != tt.want {
				t.Errorf("PeopleInfo.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildGraph(t *testing.T) {
	type args struct {
		start Node
		graph *Graph
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := buildGraph(tt.args.start, tt.args.graph); (err != nil) != tt.wantErr {
				t.Errorf("buildGraph() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCompanyBaseInfo_GetBranch(t *testing.T) {
	tests := []struct {
		name    string
		c       *CompanyBaseInfo
		want    []Node
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetBranch()
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyBaseInfo.GetBranch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyBaseInfo.GetBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeopleInfo_GetBranch(t *testing.T) {
	tests := []struct {
		name    string
		p       *PeopleInfo
		want    []Node
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetBranch()
			if (err != nil) != tt.wantErr {
				t.Errorf("PeopleInfo.GetBranch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PeopleInfo.GetBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyBaseInfo_NewGraph(t *testing.T) {
	tests := []struct {
		name    string
		c       *CompanyBaseInfo
		want    *Graph
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.NewGraph()
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyBaseInfo.NewGraph() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyBaseInfo.NewGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeopleInfo_NewGraph(t *testing.T) {
	tests := []struct {
		name    string
		p       *PeopleInfo
		want    *Graph
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.NewGraph()
			if (err != nil) != tt.wantErr {
				t.Errorf("PeopleInfo.NewGraph() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PeopleInfo.NewGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
