package ogrnOnline

import (
	"fmt"
	"time"
)

// Node - вершина графа (юридическое или физическое лицо)
type Node interface {
	GetID() int
	GetBranch() ([]Node, error)
	String() string
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

// GetBranch - метод CompanyBaseInfo возвращающий список связанных Node
func (c *CompanyBaseInfo) GetBranch() ([]Node, error) {
	aff := []Node{}

	time.Sleep(time.Millisecond * 700)
	subCompanies, err := GetSubCompanies(c.ID)
	if err != nil {
		return []Node{}, err
	}
	for _, subCompany := range subCompanies {
		aff = append(aff, &subCompany)
	}

	time.Sleep(time.Millisecond * 700)
	owners, err := GetOwners(c.ID)
	if err != nil {
		return []Node{}, err
	}
	for _, owner := range owners {
		if owner.CompanyOwner.ID != 0 {
			aff = append(aff, &owner.CompanyOwner)
		}
		if owner.PersonOwner.ID != 0 {
			aff = append(aff, &owner.PersonOwner)
		}
	}

	time.Sleep(time.Millisecond * 700)
	associates, err := GetAssociates(c.ID)
	if err != nil {
		return []Node{}, err
	}
	for _, people := range associates {
		aff = append(aff, &people.Person)
	}

	return aff, nil
}

// GetBranch - метод PeopleInfo возвращающий список связанных Node
func (p *PeopleInfo) GetBranch() ([]Node, error) {
	aff := []Node{}

	time.Sleep(time.Millisecond * 700)
	jobs, err := GetJobs(p.ID)
	if err != nil {
		return []Node{}, err
	}
	for _, job := range jobs {
		aff = append(aff, &job.Company)
	}

	time.Sleep(time.Millisecond * 700)
	shares, err := GetShare(p.ID)
	if err != nil {
		return []Node{}, err
	}
	for _, share := range shares {
		aff = append(aff, &share)
	}

	return aff, nil
}

// String - метод CompanyBaseInfo возвращающий информацию о юридическом лице в виде строки
func (c *CompanyBaseInfo) String() string {
	return fmt.Sprintf(`%s (ИНН %s)`, c.Name, c.INN)
}

// String - метод PeopleInfo возвращающий информацию о физическом лице в виде строки
func (p *PeopleInfo) String() string {
	return fmt.Sprintf(`%s (ИНН %s)`, p.FullName, p.INN)
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
