package myschoollist

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
)

// List - contains states, districts and schools in Malaysia
type List struct {
	states    *states
	districts *districts
	schools   *schools
}

// New - insantiates a new instance of states, districts and schools in Malaysia
func New() (list *List, err error) {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		err = fmt.Errorf("Error obtaining caller information")
		return
	}

	jsonDir := filepath.Join(filepath.Dir(filename), "json") // our json information is stored here

	list = new(List)

	list.states, err = newStates(jsonDir)
	if nil != err {
		return
	}

	list.districts, err = newDistricts(jsonDir)
	if nil != err {
		return
	}

	list.schools, err = newSchools(jsonDir)

	return
}

// StatesGetAll - returns list of states in Malaysia
func (list *List) StatesGetAll() (states []State) {
	return list.states.GetAll()
}

// StateGetByID - returns a state information given its ID
func (list *List) StateGetByID(stateID int) (state State, found bool) {
	return list.states.GetByID(stateID)
}

// StateGetByShortname - returns a state information given its shortname
func (list *List) StateGetByShortname(stateShortname string) (state State, found bool) {
	return list.states.GetByShortname(stateShortname)
}

// StateGetDistricts - returns slice of districts in a state given its ID
func (list *List) StateGetDistricts(stateID int) (districts []District) {

	districtsID, found := list.districts.mapDistrictsState[stateID]
	if !found {
		return
	}

	var district District
	for _, id := range districtsID {
		district, _ = list.districts.GetByID(id)
		districts = append(districts, district)
	}

	return
}

// DistrictGetByID - returns a district information given its ID
func (list *List) DistrictGetByID(districtID int) (district District, found bool) {
	return list.districts.GetByID(districtID)
}

// DistrictGetState - returns the state a district is in given its ID
func (list *List) DistrictGetState(districtID int) (state State, found bool) {

	district, found := list.districts.GetByID(districtID)
	if !found {
		return
	}

	return list.states.GetByID(district.StateID)
}

// DistrictGetSchools - returns the slice of schools in a district given its ID
func (list *List) DistrictGetSchools(districtID int) (schools []School) {

	schoolIDs, found := list.schools.mapSchoolDistrict[districtID]
	if !found {
		return
	}

	var school School
	for _, schoolID := range schoolIDs {
		school, _ = list.schools.GetByID(schoolID)
		schools = append(schools, school)
	}

	return
}

// SchoolGetByID - returns a school given its ID
func (list *List) SchoolGetByID(schoolD int) (school School, found bool) {
	return list.schools.GetByID(schoolD)
}

// SchoolGetByCode - returns a school information given its code
func (list *List) SchoolGetByCode(schoolCode string) (school School, found bool) {
	return list.schools.GetByCode(schoolCode)
}

// PrintStructure - prints a given structure to command line
func PrintStructure(object interface{}) {
	s, _ := json.MarshalIndent(object, "", "\t")
	fmt.Println(string(s))
}
