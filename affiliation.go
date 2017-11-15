package ogrnOnline

import (
	"fmt"
	"time"
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

// String - метод CompanyBaseInfo возвращающий информацию о юридическом лице в виде строки
func (c *CompanyBaseInfo) String() string {
	return fmt.Sprintf(`%s (ИНН %s)`, c.Name, c.INN)
}

// String - метод PeopleInfo возвращающий информацию о физическом лице в виде строки
func (p *PeopleInfo) String() string {
	return fmt.Sprintf(`%s (ИНН %s)`, p.FullName, p.INN)
}

// GetID - метод PeopleInfo возвращающий ID физического лица
func (p *PeopleInfo) GetID() int {
	return p.ID
}

// GetBranch - метод CompanyBaseInfo возвращающий список связанных Node
func (c *CompanyBaseInfo) GetBranch() ([]Node, error) {
	aaa := []Node{}

	time.Sleep(time.Millisecond * 700)
	sub, err := GetSubCompanies(c.ID)
	if err != nil {
		return []Node{}, err
	}
	if len(sub) != 0 {
		for i := range sub {
			aaa = append(aaa, &sub[i])
		}
	}

	time.Sleep(time.Millisecond * 700)
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

	time.Sleep(time.Millisecond * 700)
	workers, err := GetAssociates(c.ID)
	if err != nil {
		return []Node{}, err
	}
	if len(workers) != 0 {
		for i := range workers {
			aaa = append(aaa, &workers[i].Person)
		}
	}

	return aaa, nil
}

// GetBranch - метод PeopleInfo возвращающий список связанных Node
func (p *PeopleInfo) GetBranch() ([]Node, error) {
	aaa := []Node{}

	time.Sleep(time.Millisecond * 700)
	jobs, err := GetJobs(p.ID)
	if err != nil {
		return []Node{}, err
	}
	for i := range jobs {
		aaa = append(aaa, &jobs[i].Company)
	}

	time.Sleep(time.Millisecond * 700)
	shares, err := GetShare(p.ID)
	if err != nil {
		return []Node{}, err
	}
	for i := range shares {
		aaa = append(aaa, &shares[i])
	}

	return aaa, nil
}

// Построитель графа
func buildGraph(start Node, graph *Graph) error {
	nodes, err := start.GetBranch()
	if err != nil {
		return err
	}

	(*graph)[start.GetID()] = nodes
	if len(nodes) == 0 {
		return nil
	}

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
