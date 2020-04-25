package myschoollist

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// School - information about a school in Malaysia
type School struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	SchoolCode   string  `json:"schoolcode"`
	Level        string  `json:"level"`
	Address      string  `json:"address"`
	Postcode     int     `json:"postcode"`
	City         string  `json:"city"`
	EMail        string  `json:"email"`
	CoordinateXX float64 `json:"coordinatexx"`
	CoordinateYY float64 `json:"coordinateyy"`
	DisrictID    int     `json:"district_id"` // references the district this school is in
}

// schools - slice of schools in Malaysia and important mappings
type schools struct {
	items             []School
	mapSchoolID       map[int]int
	mapSchoolCode     map[string]int
	mapSchoolDistrict map[int][]int
}

// newSchools - creates new instance of Schools with proper mapping
func newSchools(jsonDir string) (schoolsV *schools, err error) {

	schoolsV = new(schools)

	file, err := os.Open(filepath.Join(jsonDir, "school.json"))
	if err != nil {
		return
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&schoolsV.items)
	if nil != err {
		return
	}

	schoolsV.mapSchoolID = make(map[int]int)
	schoolsV.mapSchoolCode = make(map[string]int)
	schoolsV.mapSchoolDistrict = make(map[int][]int)

	// we map the slice index to a map for quicker access intead of looping through the slice one by one
	for index, school := range schoolsV.items {
		schoolsV.mapSchoolID[school.ID] = index
		schoolsV.mapSchoolCode[school.SchoolCode] = index
		schoolsV.mapSchoolDistrict[school.DisrictID] = append(schoolsV.mapSchoolDistrict[school.DisrictID], school.ID)
	}

	return
}

// GetByID - returns school information given its id
func (schools *schools) GetByID(id int) (school School, found bool) {
	index, found := schools.mapSchoolID[id]
	if !found {
		return
	}

	school = schools.items[index]

	return
}

// GetByCode - returns school information given its code
func (schools *schools) GetByCode(schoolcode string) (school School, found bool) {
	index, found := schools.mapSchoolCode[strings.ToUpper(schoolcode)]
	if !found {
		return
	}

	school = schools.items[index]
	return
}
