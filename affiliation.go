package ogrnOnline

import (
	"fmt"
	"sort"
)

// Node - вершина графа (юридическое или физическое лицо)
type Node interface {
	String() string
	GetID() int
	GetBranch() ([]Node, error)
}

// Graph - граф
type Graph map[int][]Node

// GetID - метод CompanyBaseInfo возвращающий ID юридического лица
func (c *CompanyBaseInfo) GetID() int {
	return c.ID
}

// GetID - метод PeopleInfo возвращающий ID физического лица
func (p *PeopleInfo) GetID() int {
	return p.ID
}

// String - метод CompanyBaseInfo возвращающий информацию о юридическом лице в виде строки
func (c *CompanyBaseInfo) String() string {
	if c.INN != "" {
		return fmt.Sprintf(`%s (ИНН %s)`, c.Name, c.INN)
	}
	return fmt.Sprintf(`%s (ИНН неизвестен)`, c.Name)
}

// String - метод PeopleInfo возвращающий информацию о физическом лице в виде строки
func (p *PeopleInfo) String() string {
	if p.INN != "" {
		return fmt.Sprintf(`%s (ИНН %s)`, p.FullName, p.INN)
	}
	return fmt.Sprintf(`%s (ИНН неизвестен)`, p.FullName)
}

// sortNode - сортировка []Node по ID
func sortNode(sliceNode []Node) {
	if !sort.SliceIsSorted(sliceNode, func(i, j int) bool {
		return sliceNode[i].GetID() < sliceNode[j].GetID()
	},
	) {
		sort.Slice(sliceNode, func(i, j int) bool {
			return sliceNode[i].GetID() < sliceNode[j].GetID()
		})
	}
}

// clearSlice - очистка []Node от дублирующих структур
func clearSlice(sliceNode []Node) []Node {
	cleanSlice := []Node{}
	for i := range sliceNode {
		if i == 0 {
			cleanSlice = append(cleanSlice, sliceNode[i])
		}
		if sliceNode[i].GetID() != cleanSlice[len(cleanSlice)-1].GetID() {
			cleanSlice = append(cleanSlice, sliceNode[i])
		}
	}
	return cleanSlice
}

// GetBranch - метод CompanyBaseInfo возвращающий список связанных Node
func (c *CompanyBaseInfo) GetBranch() ([]Node, error) {
	aaa := []Node{}

	sub, err := GetSubCompanies(c.ID)
	if err != nil {
		return []Node{}, err
	}
	if len(sub) != 0 {
		for i := range sub {
			aaa = append(aaa, &sub[i])
		}
	}

	owners, err := GetOwners(c.ID)
	if err != nil {
		return []Node{}, err
	}
	if len(owners) != 0 {
		for i := range owners {
			if owners[i].CompanyOwner.ID != 0 {
				aaa = append(aaa, &owners[i].CompanyOwner)
			}
			if owners[i].PersonOwner.ID != 0 {
				aaa = append(aaa, &owners[i].PersonOwner)
			}
		}
	}

	workers, err := GetAssociates(c.ID)
	if err != nil {
		return []Node{}, err
	}
	if len(workers) != 0 {
		for i := range workers {
			aaa = append(aaa, &workers[i].Person)
		}
	}
	sortNode(aaa)
	return clearSlice(aaa), nil
}

// GetBranch - метод PeopleInfo возвращающий список связанных Node
func (p *PeopleInfo) GetBranch() ([]Node, error) {
	aaa := []Node{}

	jobs, err := GetJobs(p.ID)
	if err != nil {
		return []Node{}, err
	}
	for i := range jobs {
		aaa = append(aaa, &jobs[i].Company)
	}

	shares, err := GetShare(p.ID)
	if err != nil {
		return []Node{}, err
	}
	for i := range shares {
		aaa = append(aaa, &shares[i])
	}

	sortNode(aaa)
	return clearSlice(aaa), nil
}

// Построитель графа
func buildGraph(start Node, graph *Graph) error {

	nodes, err := start.GetBranch()
	if err != nil {
		return err
	}

	(*graph)[start.GetID()] = nodes

	for _, node := range nodes {
		if _, inGraph := (*graph)[node.GetID()]; !inGraph {
			err := buildGraph(node, graph)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// NewGraph - метод CompanyBaseInfo, возвращающий новый Graph для юридического лица
func (c *CompanyBaseInfo) NewGraph() (*Graph, error) {
	graph := &Graph{}
	err := buildGraph(c, graph)
	return graph, err
}

// NewGraph - метод PeopleInfo, возвращающий новый Graph для физического лица
func (p *PeopleInfo) NewGraph() (*Graph, error) {
	graph := &Graph{}
	err := buildGraph(p, graph)
	return graph, err
}
