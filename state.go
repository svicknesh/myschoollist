package myschoollist

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// State - information about a state in Malaysia
type State struct {
	ID        int    `json:"id"`
	ShortName string `json:"shortname"`
	Name      string `json:"name"`
	ISO       string `json:"iso"`
}

// states - slice of states in Malaysia and important mappings
type states struct {
	items              []State
	mapStatesID        map[int]int
	mapStatesShortname map[string]int
}

// newStates - creates new instance of States with proper mapping
func newStates(jsonDir string) (statesV *states, err error) {

	statesV = new(states)

	file, err := os.Open(filepath.Join(jsonDir, "state.json"))
	if err != nil {
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&statesV.items)
	if nil != err {
		return
	}

	statesV.mapStatesID = make(map[int]int)
	statesV.mapStatesShortname = make(map[string]int)

	// we map the slice index to a map by shortname for quicker access intead of looping through the slice one by one
	for index, state := range statesV.items {
		statesV.mapStatesID[state.ID] = index
		statesV.mapStatesShortname[state.ShortName] = index
	}

	return
}

// GetByID - returns state information given its id
func (states *states) GetByID(id int) (state State, found bool) {
	index, found := states.mapStatesID[id]
	if !found {
		return
	}

	state = states.items[index]

	return
}

// GetByShortname - returns state information given its short name
func (states *states) GetByShortname(shortname string) (state State, found bool) {
	index, found := states.mapStatesShortname[strings.ToUpper(shortname)]
	if !found {
		return
	}

	state = states.items[index]
	return
}

// GetAll - returns all states
func (states *states) GetAll() (allStates []State) {
	return states.items
}
