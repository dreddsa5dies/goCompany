package ogrnOnline

import (
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestIsValidQuery(t *testing.T) {
	type args struct {
		query     url.Values
		typeQuery int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ошибка в параметре", args{url.Values{"огрн": []string{"1051633025256"}, "ошибка": []string{"1658064460"}}, typeQueryCompany}, true},

		{"юридическое лицо. ok", args{url.Values{"огрн": []string{"1051633025256"}, "инн": []string{"1658064460"}}, typeQueryCompany}, false},
		{"юридическое лицо. ошибка в ОГРН", args{url.Values{"огрн": []string{"105133025256"}, "инн": []string{"1658064460"}}, typeQueryCompany}, true},
		{"юридическое лицо. ошибка в ИНН", args{url.Values{"огрн": []string{"1051633025256"}, "инн": []string{"16580664460"}}, typeQueryCompany}, true},

		{"физическое лицо. ok", args{url.Values{"инн": []string{"732812398429"}}, typeQueryPeople}, false},
		{"физическое лицо. ошибка в ИНН", args{url.Values{"инн": []string{"732814352398429"}}, typeQueryPeople}, true},

		{"предприниматель. ok", args{url.Values{"огрнип": []string{"314272211800010"}, "инн": []string{"272508402480"}}, typeQueryBusinessman}, false},
		{"предприниматель. ошибка в ОГРНИП", args{url.Values{"огрнип": []string{"3142722110"}, "инн": []string{"272508402480"}}, typeQueryBusinessman}, true},
		{"предприниматель. ошибка в ИНН", args{url.Values{"огрнип": []string{"314272211800010"}, "инн": []string{"272523442308402480"}}, typeQueryBusinessman}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := isValidQuery(tt.args.query, tt.args.typeQuery); (err != nil) != tt.wantErr {
				t.Errorf("isValidQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateURL(t *testing.T) {
	tests := []struct {
		name  string
		path  string
		query url.Values
		ok    bool
	}{
		{"ok", "/интеграция/ип/", url.Values{"огрнип": []string{"314272211800010"}}, true},
		{"ошибка", "/интегвыырация/ип/", url.Values{"огрнип": []string{"314272211800010"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(time.Millisecond * pauseForRequest)
			resp, err := http.Get(createURL(tt.path, tt.query).String())
			if err != nil {
				t.Skipf("непредвиденная ошибка в запросе: %v", err)
			}
			if (resp.StatusCode == 200) != tt.ok {
				t.Errorf("StatusCode = %d", resp.StatusCode)
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
		wantLen int
		wantErr bool
	}{
		{"ok", args{url.Values{"огрн": []string{"1051633025256"}, "инн": []string{"1658064460"}}}, 1, false},
		{"ok", args{url.Values{"наименование": []string{"цементно-огнеупорный завод"}}}, 4, false},
		{"ошибка", args{url.Values{"огрн": []string{"1051633025256"}, "инн": []string{"16580664460"}}}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindCompany(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if l := len(got); l != tt.wantLen {
				t.Errorf("FindCompany() = %v, wantLen %v", l, tt.wantLen)
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
		wantLen int
		wantErr bool
	}{
		{"ok", args{url.Values{"инн": []string{"732812398429"}}}, 1, false},
		{"ok", args{url.Values{"фамилия": []string{"ципорин"}, "имя": []string{"андрей"}}}, 2, false},
		{"ошибка", args{url.Values{"инн": []string{"73281238429"}}}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindPeople(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindPeople() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if l := len(got); l != tt.wantLen {
				t.Errorf("FindPeople() = %v, wantLen %v", l, tt.wantLen)
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
		wantLen int
		wantErr bool
	}{
		{"ok", args{url.Values{"огрнип": []string{"314272211800010"}, "инн": []string{"272508402480"}}}, 1, false},
		{"ok", args{url.Values{"огрнип": []string{"31422211800010"}, "инн": []string{"272508402480"}}}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindBusinessman(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindBusinessman() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if l := len(got); l != tt.wantLen {
				t.Errorf("FindBusinessman() = %v, wantLen %v", l, tt.wantLen)
			}
		})
	}
}

func TestGetCompany(t *testing.T) {
	got, err := GetCompany(7030)
	if err != nil {
		t.Skipf("GetCompany(): %v", err)
	}
	if got.OGRN != "1027810258673" {
		t.Errorf("GetCompany(): OGRN = %v, want %v", got.OGRN, "1027810258673")
	}
}

func TestGetOwners(t *testing.T) {
	got, err := GetOwners(7030)
	if err != nil {
		t.Skipf("GetOwners(): %v", err)
	}
	if l := len(got); l != 6 {
		t.Fatalf("GetOwners(): len = %v, want 6", l)
	}
	if got[3].Price != 800.0 {
		t.Errorf("GetOwners(): Price = %v, want 800.00", got[3].Price)
	}
}

func TestGetAssociates(t *testing.T) {
	got, err := GetAssociates(32357)
	if err != nil {
		t.Skipf("GetAssociates(): %v", err)
	}
	if l := len(got); l != 1 {
		t.Fatalf("GetAssociates(): len = %v, want 1", l)
	}
	if got[0].Person.ID != 49758 {
		t.Errorf("GetAssociates(): Person.ID = %v, want 49758", got[0].Person.ID)
	}
}

func TestGetSubCompanies(t *testing.T) {
	got, err := GetSubCompanies(1198655)
	if err != nil {
		t.Skipf("GetSubCompanies(): %v", err)
	}
	if l := len(got); l != 6 {
		t.Fatalf("GetSubCompanies(): len = %v, want 6", l)
	}
	if got[5].OGRN != "1156679003909" {
		t.Errorf("GetSubCompanies(): OGRN = %v, want 1156679003909", got[5].OGRN)
	}
}

func TestGetRepresentativeOffices(t *testing.T) {
	got, err := GetRepresentativeOffices(4527642)
	if err != nil {
		t.Skipf("GetRepresentativeOffices(): %v", err)
	}
	if l := len(got); l != 1 {
		t.Fatalf("GetRepresentativeOffices(): len = %v, want 1", l)
	}
	if got[0].ID != 287 {
		t.Errorf("GetRepresentativeOffices(): ID = %v, want 287", got[0].ID)
	}
}

func TestGetBranches(t *testing.T) {
	got, err := GetBranches(4527642)
	if err != nil {
		t.Skipf("GetBranches(): %v", err)
	}
	if l := len(got); l != 3 {
		t.Fatalf("GetBranches(): len = %v, want 3", l)
	}
	if got[1].ID != 289 {
		t.Errorf("GetBranches(): ID = %v, want 287", got[1].ID)
	}
}

func TestGenFinance(t *testing.T) {
	got, err := GenFinance(8)
	if err != nil {
		t.Skipf("GenFinance(): %v", err)
	}
	if l := len(got); l != 4 {
		t.Fatalf("GenFinance(): len = %v, want 3", l)
	}
	if got[1].Company.ID != 8 {
		t.Errorf("GenFinance(): ID = %v, want 8", got[1].Company.ID)
	}
}

func TestGetPeople(t *testing.T) {
	got, err := GetPeople(2191023)
	if err != nil {
		t.Skipf("GetPeople(): %v", err)
	}
	if got.FirstName != "АЛЕКСЕЙ" {
		t.Errorf("GetPeople(): FirstName = %s, want АЛЕКСЕЙ", got.FirstName)
	}
}

func TestGetJobs(t *testing.T) {
	got, err := GetJobs(2191023)
	if err != nil {
		t.Skipf("GetJobs(): %v", err)
	}
	if l := len(got); l != 5 {
		t.Fatalf("GetJobs(): len = %v, want 5", l)
	}
	if got[1].Company.OGRN != "1107847114760" {
		t.Errorf("GetJobs(): OGRN = %s, want 1107847114760", got[1].Company.OGRN)
	}
}
func TestGetShare(t *testing.T) {
	got, err := GetShare(2191023)
	if err != nil {
		t.Skipf("GetShare(): %v", err)
	}
	if l := len(got); l != 9 {
		t.Fatalf("GetShare(): len = %v, want 9", l)
	}
	if got[1].OGRN != "1037851007424" {
		t.Errorf("GetShare(): OGRN = %s, want 10271037851007424810258673", got[1].OGRN)
	}
}
