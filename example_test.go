package ogrnOnline

import (
	"fmt"
	"net/url"
)

func ExampleWork() {
	find, err := FindPeople(url.Values{"инн": []string{"740413956177"}})
	if err != nil {
		fmt.Printf("FindPeople() error = %v\n", err)
		return
	}
	if len(find) == 0 || len(find) > 1 {
		fmt.Println("ошибка в количестве найденных лиц")
		return
	}
	people := find[0]

	graph, err := people.NewGraph()
	if err != nil {
		fmt.Printf("NewGraph() error = %v\n", err)
		return
	}
	test, err := GetPeople(6382336)
	if err != nil {
		fmt.Printf("GetPeople() error = %v\n", err)
		return
	}
	connection := people.FindConnection(graph, &test)
	fmt.Printf("%v\n", connection)

	// Output:
	// [ЧАРИЕВА ЕЛЕНА АЛЕКСАНДРОВНА (ИНН 740413956177) ЗАКРЫТОЕ АКЦИОНЕРНОЕ ОБЩЕСТВО "УПРАВЛЕНИЕ БИЗНЕС РЕСУРСАМИ" (ИНН 6659158336) ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ "ЮРИДИЧЕСКАЯ ФИРМА "ЕЛАКС" (ИНН 6658188546) НЕКОММЕРЧЕСКОЕ ПАРТНЕРСТВО "КООРДИНАЦИОННЫЙ ЦЕНТР РАЗВИТИЯ ЭКСПЕРТНОЙ ДЕЯТЕЛЬНОСТИ" (ИНН 5902989536) ЮГАС ДМИТРИЙ ВАСИЛЬЕВИЧ (ИНН 594700034243)]
}
