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
		return []Node{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %s: %v", c.ID, err)
	}
	for i := range sub {
		nodes = append(nodes, &sub[i])
	}

	owners, err := GetOwners(c.ID)
	if err != nil {
		return []Node{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %s: %v", c.ID, err)
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
		return []Node{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %s: %v", c.ID, err)
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
		return []Node{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %s: %v", p.ID, err)
	}
	for i := range jobs {
		nodes = append(nodes, &jobs[i].Company)
	}

	shares, err := GetShare(p.ID)
	if err != nil {
		return []Node{}, fmt.Errorf("getBranch: ошибка обработки компании с ID %s: %v", p.ID, err)
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

// findConnectionBetweenNode обходит граф и ищет связи между Node
func findConnectionBetweenNode(graph *Graph, start, finish Node, connection []Node) []Node {
	connection = append(connection, start)

	if start.getID() == finish.getID() {
		return connection
	}

	if _, inGraph := (*graph)[start.getID()]; !inGraph {
		return []Node{}
	}

	for _, node := range (*graph)[start.getID()] {
		if !inSlice(node, connection) {
			testConnection := findConnectionBetweenNode(graph, node, finish, connection)
			if len(testConnection) != 0 {
				return testConnection
			}
		}
	}
	return []Node{}
}

// FindConnection обходит граф и ищет связи с finishNode
func (c *CompanyInfo) FindConnection(graph *Graph, finishNode Node) []Node {
	return findConnectionBetweenNode(graph, c, finishNode, []Node{})
}

// FindConnection обходит граф и ищет связи с finishNode
func (p *PeopleInfo) FindConnection(graph *Graph, finishNode Node) []Node {
	return findConnectionBetweenNode(graph, p, finishNode, []Node{})
}
