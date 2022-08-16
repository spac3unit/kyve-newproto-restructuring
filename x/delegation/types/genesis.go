package types

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:                     DefaultParams(),
		DelegatorList:              []Delegator{},
		DelegationEntryList:        []DelegationEntry{},
		DelegationDataList:         []DelegationData{},
		UndelegationQueueEntryList: []UndelegationQueueEntry{},
		QueueStateUndelegation:     QueueState{},
		RedelegationCooldownList:   []RedelegationCooldown{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	// TODO implement genesis validation

	return gs.Params.Validate()
}
