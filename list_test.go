package myschoollist

import (
	"log"
	"testing"
)

func TestList(t *testing.T) {

	list, err := New()
	if nil != err {
		log.Panic(err)
	}

	//PrintStructure(list.StatesGetAll())

	// get some state information
	stateID := 7
	log.Printf("Get state information by ID %d", stateID)
	state, found := list.StateGetByID(stateID)
	if !found {
		log.Printf("State ID %d not found", stateID)
	}
	PrintStructure(state)

	stateShortname := "png"
	log.Printf("Get state information by shortname %q", stateShortname)
	state, found = list.StateGetByShortname(stateShortname)
	if !found {
		log.Printf("State shortname %q not found", stateShortname)
	}
	PrintStructure(state)

	log.Printf("Get districts in state ID %d", stateID)
	districts := list.StateGetDistricts(stateID)
	PrintStructure(districts)

	// --- end state ---

	// get some district information
	districtID := 73
	log.Printf("Get district information by ID %d", districtID)
	district, found := list.DistrictGetByID(districtID)
	if !found {
		log.Printf("District ID %d not found", districtID)
	}
	PrintStructure(district)

	log.Printf("Get state information by district ID %d", districtID)
	state, found = list.DistrictGetState(districtID)
	if !found {
		log.Printf("District ID %d not found", districtID)
	}
	PrintStructure(state)

	log.Printf("Get schools information in district ID %d", districtID)
	schools := list.DistrictGetSchools(districtID)
	PrintStructure(schools)

	// --- end district ---

	// get some school information
	schoolID := 9429
	log.Printf("Get school information by ID %d", schoolID)
	school, found := list.SchoolGetByID(schoolID)
	if !found {
		log.Printf("School ID %d not found", schoolID)
	}
	PrintStructure(school)

	schoolCode := "peb1094"
	log.Printf("Get school information by code %q", schoolCode)
	school, found = list.SchoolGetByCode(schoolCode)
	if !found {
		log.Printf("School code %q not found", schoolCode)
	}
	PrintStructure(school)

	// --- end school ---

}
