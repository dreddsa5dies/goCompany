package ogrnOnline

import (
	"fmt"
	"net/url"
	"time"
)

// По соглашению API допускает не более 3 обращений в секунду
func ExampleFindCompany() {
	time.Sleep(time.Millisecond * 500)

	query := url.Values{}
	query.Add("инн", "7707083893")

	req, _ := FindCompany(query)

	fmt.Println(req[0].ShortName)
	// Output:
	// ПАО СБЕРБАНК
}
func ExampleFindPeople() {
	time.Sleep(time.Millisecond * 500)

	query := url.Values{}
	query.Add("фамилия", "клибанец")
	query.Add("имя", "андрей")

	req, _ := FindPeople(query)

	fmt.Printf("%s %s %s", req[0].SurName, req[0].FirstName, req[0].MiddleName)
	// Output:
	// КЛИБАНЕЦ АНДРЕЙ АЛЕКСАНДРОВИЧ
}
func ExampleFindBusinessman() {
	time.Sleep(time.Millisecond * 500)

	query := url.Values{}
	query.Add("огрнип", "314272211800010")

	req, _ := FindBusinessman(query)

	fmt.Println(req[0].Person.FullName)
	// Output:
	// НИКУЛИН ИЛЬЯ АЛЕКСАНДРОВИЧ
}
func ExampleGetCompany() {
	time.Sleep(time.Millisecond * 500)
	cmp, _ := GetCompany(86225)
	fmt.Printf("Уставной капитал Сбербанка составляет %.2f", cmp.AuthorizedCapital.Value)
	// Output:
	// Уставной капитал Сбербанка составляет 67760844000.00
}
func ExampleGetOwners() {
	time.Sleep(time.Millisecond * 500)
	own, _ := GetOwners(86225)
	fmt.Printf("Единственный акционер Сбербанка %s", own[0].CompanyOwner.Name)
	// Output:
	// Единственный акционер Сбербанка ЦЕНТРАЛЬНЫЙ БАНК РОССИЙСКОЙ ФЕДЕРАЦИИ
}
func ExampleGetAssociates() {
	time.Sleep(time.Millisecond * 500)
	ass, _ := GetAssociates(86225)
	fmt.Printf("%s %s, дайте денег )))", ass[0].Person.FirstName, ass[0].Person.MiddleName)
	// Output:
	// ГЕРМАН ОСКАРОВИЧ, дайте денег )))
}
