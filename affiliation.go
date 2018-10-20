package goCompany

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Graph граф (карта анализируемых связей)
type Graph map[int]Branch

// Branch ветка графа
type Branch []Node

// Node вершина графа (юридическое или физическое лицо)
type Node interface {
	getID() int
	String() string
	getBranch() (Branch, error)
}

// getID возвращает ID
func (c *CompanyInfo) getID() int {
	return c.ID
}

// getID возвращает ID
func (p *PeopleInfo) getID() int {
	return p.ID
}

// String возвращает CompanyInfo в строковом представлении
func (c *CompanyInfo) String() string {
	if c.INN != "" {
		return fmt.Sprintf("[ %s : ИНН %s ]", c.Name, c.INN)
	}
	return fmt.Sprintf("[ %s : ИНН неизвестен ]", c.Name)
}

// String возвращает PeopleInfo в строковом представлении
func (p *PeopleInfo) String() string {
	if p.INN != "" {
		return fmt.Sprintf("[ %s : ИНН %s ]", p.FullName, p.INN)
	}
	return fmt.Sprintf("[ %s : ИНН неизвестен ]", p.FullName)
}

func (b Branch) sortByID() {
	if !sort.SliceIsSorted(b, func(i, j int) bool {
		return b[i].getID() < b[j].getID()
	},
	) {
		sort.Slice(b, func(i, j int) bool {
			return b[i].getID() < b[j].getID()
		})
	}
}

// ClearDouble очищает Branch от дублирующих друг друга Node
func (b *Branch) ClearDouble() {
	var clean Branch
	b.sortByID()

	for i, n := range *b {
		if i == 0 {
			clean = append(clean, n)
			continue
		}
		if n.getID() != clean[len(clean)-1].getID() {
			clean = append(clean, n)
		}
	}
	*b = clean
}

// isBankrupt возвращает true, если компанией управляет арбитражный управляющий
func isBankrupt(post string) bool {
	res, err := regexp.MatchString(`(арбитражный|временный|внешний|административный|конкурсный)`, strings.ToLower(post))
	if err != nil {
		panic(fmt.Errorf("isBankrupt: ошибка при проверке должности %s = %v", post, err))
	}
	return res
}

// isCommercial возвращает true если компания коммерческая
func isCommercial(okopf string) bool {
	res, err := regexp.MatchString(`^(1|6)`, okopf)
	if err != nil {
		panic(fmt.Errorf("isCommercial: ошибка при проверке ОКОПФ %s = %v", okopf, err))
	}
	return res
}

// getBranch возвращает список связанных Node
func (c *CompanyInfo) getBranch() (Branch, error) {
	var (
		nodes     Branch
		companies SliceCompanyInfo
		peoples   SlicePeopleInfo
	)

	sub, err := GetSubCompanies(c.ID)
	if err != nil {
		return Branch{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %d: %v", c.ID, err)
	}
	companies = append(companies, sub...)

	owners, err := GetOwners(c.ID)
	if err != nil {
		return Branch{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %d: %v", c.ID, err)
	}
	cc, pp := owners.ExtractOwners()
	companies = append(companies, cc...)
	peoples = append(peoples, pp...)

	workers, err := GetAssociates(c.ID)
	if err != nil {
		return Branch{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %d: %v", c.ID, err)
	}
	for _, work := range workers {
		if !isBankrupt(work.PostName) {
			peoples = append(peoples, work.Person)
		}
	}

	peoples.ClearDouble()
	for i := range peoples {
		nodes = append(nodes, &peoples[i])
	}

	companies.ClearDouble()
	companies.ClearInactive()
	for i := range companies {
		full, err := GetCompanyFullInfo(companies[i].ID)
		if err != nil {
			return Branch{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %d: %v", c.ID, err)
		}
		if isCommercial(full.OKOPF.Code) {
			nodes = append(nodes, &companies[i])
		}
	}

	return nodes, nil
}

// getBranch возвращает список связанных Node
func (p *PeopleInfo) getBranch() (Branch, error) {
	var (
		nodes     Branch
		companies SliceCompanyInfo
	)

	shares, err := GetShare(p.ID)
	if err != nil {
		return []Node{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %d: %v", p.ID, err)
	}
	companies = append(companies, shares...)

	jobs, err := GetJobs(p.ID)
	if err != nil {
		return []Node{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %d: %v", p.ID, err)
	}
	for _, work := range jobs {
		if !isBankrupt(work.PostName) {
			companies = append(companies, work.Company)
		}
	}

	companies.ClearDouble()
	companies.ClearInactive()
	for i := range companies {
		full, err := GetCompanyFullInfo(companies[i].ID)
		if err != nil {
			return Branch{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %d: %v", p.ID, err)
		}
		if isCommercial(full.OKOPF.Code) {
			nodes = append(nodes, &companies[i])
		}
	}

	return nodes, nil
}

// Построитель графа
func buildGraph(start Node, graph *Graph) error {
	nodes, err := start.getBranch()
	if err != nil {
		return fmt.Errorf("buildGraph: %v", err)
	}
	(*graph)[start.getID()] = nodes
	for _, node := range nodes {
		if _, inGraph := (*graph)[node.getID()]; !inGraph {
			err := buildGraph(node, graph)
			if err != nil {
				return fmt.Errorf("buildGraph: %v", err)
			}
		}
	}
	return nil
}

// NewGraph возвращает Graph
func (c *CompanyInfo) NewGraph() (*Graph, error) {
	graph := &Graph{}
	err := buildGraph(c, graph)
	if err != nil {
		return &Graph{}, fmt.Errorf("NewGraph: %v", err)
	}
	return graph, nil
}

// NewGraph возвращает Graph
func (p *PeopleInfo) NewGraph() (*Graph, error) {
	graph := &Graph{}
	err := buildGraph(p, graph)
	if err != nil {
		return &Graph{}, fmt.Errorf("NewGraph: %v", err)
	}
	return graph, nil
}

// inSlice проверяет наличие Node в []Node
func inSlice(node Node, sliceNode []Node) bool {
	for _, elem := range sliceNode {
		if node.getID() == elem.getID() {
			return true
		}
	}
	return false
}

// findConnectionBetweenNodes возвращает true если связь между Node есть
func findConnectionBetweenNodes(graph *Graph, start, finish Node, connection []Node) bool {
	// Если стартовой Node нет в Graph, возвращается false
	if _, inGraph := (*graph)[start.getID()]; !inGraph {
		return false
	}

	// Стартовая Node добавляется в список просмотренных
	connection = append(connection, start)

	// При совпадении стартовой и финишной Node возвращется результат true
	if start.getID() == finish.getID() {
		return true
	}

	//	Для каждой Node, с которой есть связь у стартовой Node (если она уже не была проверена ранее)
	//	рекурсивно запускается функция findConnectionBetweenNodes.
	//	Положтельный результат влечет возврат true, отсутствие результата - false.
	for _, node := range (*graph)[start.getID()] {
		if !inSlice(node, connection) {
			if findConnectionBetweenNodes(graph, node, finish, connection) {
				return true
			}
		}
	}

	return false
}

// findAllConnectionBetweenNode обходит граф и ищет связи между Node
func findAllConnectionBetweenNode(graph *Graph, start, finish Node, connection []Node) [][]Node {
	// Если стартовой Node нет в Graph, возвращается пустой [][]Node
	if _, inGraph := (*graph)[start.getID()]; !inGraph {
		return [][]Node{}
	}

	// Стартовая Node добавляется в список просмотренных
	connection = append(connection, start)

	// При совпадении стартовой и финишной Node возвращется текущая цепочка Node, обернутая в [][]Node
	if start.getID() == finish.getID() {
		return [][]Node{connection}
	}

	// Создается [][]Node для всех найденных связей
	connections := [][]Node{}

	//	Для каждой Node, с которой есть связь у стартовой Node (если она уже не была проверена ранее)
	//	рекурсивно запускается функция findAllConnectionBetweenNode.
	//	Положтельный результат вкладывадывается в общий слайс.
	for _, node := range (*graph)[start.getID()] {
		if !inSlice(node, connection) {
			newConnection := findAllConnectionBetweenNode(graph, node, finish, connection)
			if len(newConnection) != 0 {
				for _, connect := range newConnection {
					connections = append(connections, connect)
				}
			}
		}
	}

	// Накопленый слайс возвращается
	return connections
}

// ConnectionIsExist обходит граф и ищет связь с finishNode. По результату возвращает bool
func (c *CompanyInfo) ConnectionIsExist(graph *Graph, finishNode Node) bool {
	return findConnectionBetweenNodes(graph, c, finishNode, []Node{})
}

// ConnectionIsExist обходит граф и ищет связь с finishNode. По результату возвращает bool
func (p *PeopleInfo) ConnectionIsExist(graph *Graph, finishNode Node) bool {
	return findConnectionBetweenNodes(graph, p, finishNode, []Node{})
}

// FindAllConnections обходит граф и возвращает все возможные связи между Node
func (c *CompanyInfo) FindAllConnections(graph *Graph, finishNode Node) [][]Node {
	return findAllConnectionBetweenNode(graph, c, finishNode, []Node{})
}

// FindAllConnections обходит граф и возвращает все возможные связи между Node
func (p *PeopleInfo) FindAllConnections(graph *Graph, finishNode Node) [][]Node {
	return findAllConnectionBetweenNode(graph, p, finishNode, []Node{})
}
