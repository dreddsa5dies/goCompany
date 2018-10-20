package goCompany

import "sort"

// IsActive возвращает true, если юридическое лицо действующее
func (c *CompanyInfo) IsActive() bool {
	return c.CloseInfo.Date == ""
}

type sorterByID interface {
	sortByID()
}

// sortByID сортирует SliceCompanyInfo по ID
func (cc SliceCompanyInfo) sortByID() {
	if !sort.SliceIsSorted(cc, func(i, j int) bool {
		return cc[i].ID < cc[j].ID
	}) {
		sort.Slice(cc, func(i, j int) bool {
			return cc[i].ID < cc[j].ID
		})
	}
}

// sortByID сортирует SlicePeopleInfo по ID
func (pp SlicePeopleInfo) sortByID() {
	if !sort.SliceIsSorted(pp, func(i, j int) bool {
		return pp[i].ID < pp[j].ID
	}) {
		sort.Slice(pp, func(i, j int) bool {
			return pp[i].ID < pp[j].ID
		})
	}
}

// ExtractOwners извлекает информацию об участниках юридического лица и возвращает отдельные
// слайсы юридических и физических лиц
func (slice *SliceCompanyOwnerInfo) ExtractOwners() (SliceCompanyInfo, SlicePeopleInfo) {
	var (
		cc SliceCompanyInfo
		pp SlicePeopleInfo
	)

	for _, owner := range *slice {
		if owner.CompanyOwner.ID != 0 {
			cc = append(cc, owner.CompanyOwner)
		}
		if owner.PersonOwner.ID != 0 {
			pp = append(pp, owner.PersonOwner)
		}
	}

	return cc, pp
}

// ExtractCompany возвращает SliceCompanyInfo (организации, в которых работает человек)
func (slice *SliceCompanyAssociateInfo) ExtractCompany() SliceCompanyInfo {
	var cc SliceCompanyInfo

	for _, elem := range *slice {
		cc = append(cc, elem.Company)
	}

	return cc
}

// ExtractPeople возвращает SlicePeopleInfo (работников организации)
func (slice *SliceCompanyAssociateInfo) ExtractPeople() SlicePeopleInfo {
	var pp SlicePeopleInfo

	for _, elem := range *slice {
		pp = append(pp, elem.Person)
	}

	return pp
}

// ClearInactive очищает SliceCompanyInfo от недействующих компаний
func (cc *SliceCompanyInfo) ClearInactive() {
	var slice SliceCompanyInfo
	cc.sortByID()

	for _, company := range *cc {
		if company.IsActive() {
			slice = append(slice, company)
		}
	}

	*cc = slice
}

type clearnerDouble interface {
	ClearDouble()
}

// ClearDouble очищает SliceCompanyInfo от дублирующих компаний
func (cc *SliceCompanyInfo) ClearDouble() {
	var slice SliceCompanyInfo
	cc.sortByID()

	for i, company := range *cc {
		if i == 0 {
			slice = append(slice, company)
			continue
		}
		if company.ID != slice[len(slice)-1].ID {
			slice = append(slice, company)
		}
	}

	*cc = slice
}

// ClearDouble очищает SlicePeopleInfo от дублирующих людей
func (pp *SlicePeopleInfo) ClearDouble() {
	var slice SlicePeopleInfo
	pp.sortByID()

	for i, people := range *pp {
		if i == 0 {
			slice = append(slice, people)
			continue
		}
		if people.ID != slice[len(slice)-1].ID {
			slice = append(slice, people)
		}
	}

	*pp = slice
}
