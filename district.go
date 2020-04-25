package myschoollist

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// District - information about a district and mapping to district
type District struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	StateID int    `json:"state_id"` // references the state this district is in
}

// districts - slice of districts in Malaysia and important mappings
type districts struct {
	items             []District
	mapDistrictsID    map[int]int
	mapDistrictsState map[int][]int
}

// newDistricts - creates new instance of Districts with proper mapping
func newDistricts(jsonDir string) (districtsV *districts, err error) {

	districtsV = new(districts)

	file, err := os.Open(filepath.Join(jsonDir, "district.json"))
	if err != nil {
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&districtsV.items)
	if nil != err {
		return
	}

	districtsV.mapDistrictsID = make(map[int]int)
	districtsV.mapDistrictsState = make(map[int][]int)

	// we map the slice index to a map for quicker access intead of looping through the slice one by one
	for index, district := range districtsV.items {
		districtsV.mapDistrictsID[district.ID] = index
		districtsV.mapDistrictsState[district.StateID] = append(districtsV.mapDistrictsState[district.StateID], district.ID)
	}

	return
}

// GetByID - returns district information given its id
func (districts *districts) GetByID(id int) (district District, found bool) {
	index, found := districts.mapDistrictsID[id]
	if !found {
		return
	}

	district = districts.items[index]

	return
}
