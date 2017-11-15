package ogrnOnline

import (
	"net/url"
	"reflect"
	"testing"
)

func Test_createURL(t *testing.T) {
	type args struct {
		path  string
		query url.Values
	}
	tests := []struct {
		name string
		args args
		want *url.URL
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createURL(tt.args.path, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidQuery(t *testing.T) {
	type args struct {
		q         url.Values
		typeQuery int
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
			if err := isValidQuery(tt.args.q, tt.args.typeQuery); (err != nil) != tt.wantErr {
				t.Errorf("isValidQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getDataFromServer(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDataFromServer(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDataFromServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindCompany(t *testing.T) {
	type args struct {
		query url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    []CompanyBaseInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindCompany(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindCompany() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindPeople(t *testing.T) {
	type args struct {
		query url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    []PeopleInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindPeople(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindPeople() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPeople() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindBusinessman(t *testing.T) {
	type args struct {
		query url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    []PeopleBusinessmanInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindBusinessman(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindBusinessman() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindBusinessman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCompany(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    CompanyInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCompany(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCompany() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyBaseInfo_GetCompany(t *testing.T) {
	tests := []struct {
		name    string
		c       *CompanyBaseInfo
		want    CompanyInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetCompany()
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyBaseInfo.GetCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyBaseInfo.GetCompany() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOwners(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    []CompanyOwnerInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetOwners(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOwners() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOwners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyBaseInfo_GetOwners(t *testing.T) {
	tests := []struct {
		name    string
		c       *CompanyBaseInfo
		want    []CompanyOwnerInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetOwners()
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyBaseInfo.GetOwners() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyBaseInfo.GetOwners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAssociates(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    []CompanyAssociateInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAssociates(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAssociates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAssociates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyBaseInfo_GetAssociates(t *testing.T) {
	tests := []struct {
		name    string
		c       *CompanyBaseInfo
		want    []CompanyAssociateInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetAssociates()
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyBaseInfo.GetAssociates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyBaseInfo.GetAssociates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSubCompanies(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    []CompanyBaseInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSubCompanies(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubCompanies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubCompanies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyBaseInfo_GetSubCompanies(t *testing.T) {
	tests := []struct {
		name    string
		c       *CompanyBaseInfo
		want    []CompanyBaseInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetSubCompanies()
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyBaseInfo.GetSubCompanies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyBaseInfo.GetSubCompanies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRepresentativeOffices(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    []CompanyBranchInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRepresentativeOffices(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepresentativeOffices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRepresentativeOffices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyBaseInfo_GetRepresentativeOffices(t *testing.T) {
	tests := []struct {
		name    string
		c       *CompanyBaseInfo
		want    []CompanyBranchInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetRepresentativeOffices()
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyBaseInfo.GetRepresentativeOffices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyBaseInfo.GetRepresentativeOffices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBranches(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    []CompanyBranchInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBranches(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBranches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBranches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyBaseInfo_GetBranches(t *testing.T) {
	tests := []struct {
		name    string
		c       *CompanyBaseInfo
		want    []CompanyBranchInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetBranches()
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyBaseInfo.GetBranches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyBaseInfo.GetBranches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenFinance(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    []CompanyFinanceInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenFinance(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenFinance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenFinance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompanyBaseInfo_GetFinance(t *testing.T) {
	tests := []struct {
		name    string
		c       *CompanyBaseInfo
		want    []CompanyFinanceInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetFinance()
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyBaseInfo.GetFinance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyBaseInfo.GetFinance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPeople(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    PeopleInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPeople(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPeople() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPeople() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetJobs(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    []CompanyAssociateInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetJobs(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeopleInfo_GetJobs(t *testing.T) {
	tests := []struct {
		name    string
		p       *PeopleInfo
		want    []CompanyAssociateInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetJobs()
			if (err != nil) != tt.wantErr {
				t.Errorf("PeopleInfo.GetJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PeopleInfo.GetJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetShare(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    []CompanyBaseInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetShare(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetShare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetShare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeopleInfo_GetShare(t *testing.T) {
	tests := []struct {
		name    string
		p       *PeopleInfo
		want    []CompanyBaseInfo
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetShare()
			if (err != nil) != tt.wantErr {
				t.Errorf("PeopleInfo.GetShare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PeopleInfo.GetShare() = %v, want %v", got, tt.want)
			}
		})
	}
}
