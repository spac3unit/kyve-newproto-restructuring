package types

import "sort"

func (m *Staker) AddPool(poolId uint64) {

	m.Pools = append(m.Pools, poolId)

	// TODO assume sorted array and insert at the right index
	sort.SliceStable(m.Pools, func(i, j int) bool {
		return m.Pools[i] < m.Pools[j]
	})
}

func (m *Staker) RemovePool(poolId uint64) {

	index := sort.Search(len(m.Pools), func(i int) bool {
		return m.Pools[i] > poolId
	})

	// TODO check for boundaries
	m.Pools = append(m.Pools[0:index], m.Pools[index+1:len(m.Pools)]...)
}
