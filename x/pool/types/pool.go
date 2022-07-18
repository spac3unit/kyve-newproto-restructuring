package types

import (
	"sort"
)

// TODO update pool total stake

func (m *Pool) UpdateFunder(funder Funder) {

	oldFunder, found := m.GetFunder(funder.Account)
	_ = oldFunder
	if !found {
		m.InsertFunder(funder)
	} else {

		sort.SliceStable(m.Funders, func(i, j int) bool {
			return m.Funders[i].Amount < m.Funders[j].Amount
		})
	}
}

func (m *Pool) InsertFunder(funder Funder) {

	m.Funders = append(m.Funders, &funder)

	sort.SliceStable(m.Funders, func(i, j int) bool {
		return m.Funders[i].Amount < m.Funders[j].Amount
	})
}

func (m *Pool) RemoveFunder(funder Funder) {

	index := sort.Search(len(m.Funders), func(i int) bool {
		return m.Funders[i].Amount > funder.Amount
	})

	// TODO check for same amount
	m.Funders = append(m.Funders[0:index], m.Funders[index+1:len(m.Funders)]...)
}

func (m *Pool) GetFunder(address string) (Funder, bool) {
	for _, v := range m.Funders {
		if v.Account == address {
			return *v, true
		}
	}
	return Funder{}, false
}

func (m *Pool) GetLowestFunder() *Funder {
	if len(m.Funders) == 0 {
		return nil
	}
	return m.Funders[0]
}
