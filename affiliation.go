package ogrnOnline

import (
	"fmt"
	"sort"
)

// Graph граф (карта анализируемых связей)
type Graph map[int][]Node

// Node вершина графа (юридическое или физическое лицо)
type Node interface {
	String() string
	getID() int
	getBranch() ([]Node, error)
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

// sortNode сортирует []Node по ID
func sortNode(sliceNode []Node) {
	if !sort.SliceIsSorted(sliceNode, func(i, j int) bool {
		return sliceNode[i].getID() < sliceNode[j].getID()
	},
	) {
		sort.Slice(sliceNode, func(i, j int) bool {
			return sliceNode[i].getID() < sliceNode[j].getID()
		})
	}
}

// clearSlice очищает []Node от дублирующих друг друга Node
func clearSlice(sliceNode []Node) []Node {
	cleanSlice := []Node{}
	for i := range sliceNode {
		if i == 0 {
			cleanSlice = append(cleanSlice, sliceNode[i])
		}
		if sliceNode[i].getID() != cleanSlice[len(cleanSlice)-1].getID() {
			cleanSlice = append(cleanSlice, sliceNode[i])
		}
	}
	return cleanSlice
}

// getBranch возвращает список связанных Node
func (c *CompanyInfo) getBranch() ([]Node, error) {
	nodes := []Node{}

	sub, err := GetSubCompanies(c.ID)
	if err != nil {
		return []Node{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %d: %v", c.ID, err)
	}
	for i := range sub {
		nodes = append(nodes, &sub[i])
	}

	owners, err := GetOwners(c.ID)
	if err != nil {
		return []Node{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %d: %v", c.ID, err)
	}
	for i := range owners {
		if owners[i].CompanyOwner.ID != 0 {
			nodes = append(nodes, &owners[i].CompanyOwner)
		}
		if owners[i].PersonOwner.ID != 0 {
			nodes = append(nodes, &owners[i].PersonOwner)
		}
	}

	workers, err := GetAssociates(c.ID)
	if err != nil {
		return []Node{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %d: %v", c.ID, err)
	}
	for i := range workers {
		nodes = append(nodes, &workers[i].Person)
	}

	sortNode(nodes)
	return clearSlice(nodes), nil
}

// getBranch возвращает список связанных Node
func (p *PeopleInfo) getBranch() ([]Node, error) {
	nodes := []Node{}

	jobs, err := GetJobs(p.ID)
	if err != nil {
		return []Node{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %d: %v", p.ID, err)
	}
	for i := range jobs {
		nodes = append(nodes, &jobs[i].Company)
	}

	shares, err := GetShare(p.ID)
	if err != nil {
		return []Node{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %d: %v", p.ID, err)
	}
	for i := range shares {
		nodes = append(nodes, &shares[i])
	}

	sortNode(nodes)
	return clearSlice(nodes), nil
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
	return graph, fmt.Errorf("NewGraph: %v", err)
}

// NewGraph возвращает Graph
func (p *PeopleInfo) NewGraph() (*Graph, error) {
	graph := &Graph{}
	err := buildGraph(p, graph)
	return graph, fmt.Errorf("NewGraph: %v", err)
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

	/*
		Для каждой Node, с которой есть связь у стартовой Node (если она уже не была проверена ранее)
		рекурсивно запускается функция findConnectionBetweenNodes.
		Положтельный результат влечет возврат true, отсутствие результата - false.
	*/
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

	/*
		Для каждой Node, с которой есть связь у стартовой Node (если она уже не была проверена ранее)
		рекурсивно запускается функция findAllConnectionBetweenNode.
		Положтельный результат вкладывадывается в общий слайс.
	*/
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
