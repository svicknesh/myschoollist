package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
)

const (
	srcState    = "https://raw.githubusercontent.com/apitlekays/MySchoolList/master/src/state.json"
	srcDistrict = "https://raw.githubusercontent.com/apitlekays/MySchoolList/master/src/district.json"
	srcSchool   = "https://raw.githubusercontent.com/apitlekays/MySchoolList/master/src/school.json"
)

type inState struct {
	ID        string `json:"id"`
	ShortName string `json:"shortname"`
	Name      string `json:"name"`
	ISO       string `json:"iso"`
}

type outState struct {
	ID        int    `json:"id"`
	ShortName string `json:"shortname"`
	Name      string `json:"name"`
	ISO       string `json:"iso"`
}

type inDistrict struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	StateID string `json:"state_id"` // references the district this district is in
}

type outDistrict struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	StateID int    `json:"state_id"` // references the district this district is in
}

type inSchool struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	SchoolCode   string `json:"schoolcode"`
	Level        string `json:"level"`
	Address      string `json:"address"`
	Postcode     string `json:"postcode"`
	City         string `json:"city"`
	EMail        string `json:"email"`
	CoordinateXX string `json:"coordinatexx"`
	CoordinateYY string `json:"coordinateyy"`
	DisrictID    string `json:"district_id"` // references the school this school is in
}

type outSchool struct {
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
	DisrictID    int     `json:"district_id"` // references the school this school is in
}

func main() {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Panic("Error obtaining caller information")
	}

	jsonDir := filepath.Dir(filename) // our json information is stored here
	_ = jsonDir

	// import and transform STATE
	var inStates []inState
	err := httpConnect(srcState, &inStates)
	if nil != err {
		log.Panic(err)
	}

	var oS outState
	outStates := make([]outState, len(inStates))

	for index, state := range inStates {
		oS.ID, _ = strconv.Atoi(state.ID)
		oS.ISO = state.ISO
		oS.Name = state.Name
		oS.ShortName = state.ShortName
		outStates[index] = oS
	}

	dataBytes, _ := json.Marshal(outStates)
	ioutil.WriteFile(filepath.Join(jsonDir, "state.json"), dataBytes, 0600)

	// import and transform DISTRICT
	var inDistricts []inDistrict
	err = httpConnect(srcDistrict, &inDistricts)
	if nil != err {
		log.Panic(err)
	}

	var oD outDistrict
	outDistricts := make([]outDistrict, len(inDistricts))

	for index, distict := range inDistricts {
		oD.ID, _ = strconv.Atoi(distict.ID)
		oD.StateID, _ = strconv.Atoi(distict.StateID)
		oD.Name = distict.Name
		outDistricts[index] = oD
	}

	dataBytes, _ = json.Marshal(outDistricts)
	ioutil.WriteFile(filepath.Join(jsonDir, "district.json"), dataBytes, 0600)

	// import and transform SCHOOL
	var inSchools []inSchool
	err = httpConnect(srcSchool, &inSchools)
	if nil != err {
		log.Panic(err)
	}

	var oSc outSchool
	outSchools := make([]outSchool, len(inSchools))

	for index, school := range inSchools {
		oSc.ID, _ = strconv.Atoi(school.ID)
		oSc.DisrictID, _ = strconv.Atoi(school.DisrictID)
		oSc.Postcode, _ = strconv.Atoi(school.Postcode)

		oSc.CoordinateXX, _ = strconv.ParseFloat(school.CoordinateXX, 64)
		oSc.CoordinateYY, _ = strconv.ParseFloat(school.CoordinateYY, 64)

		oSc.Name = school.Name
		oSc.SchoolCode = school.SchoolCode
		oSc.Level = school.Level
		oSc.Address = school.Address
		oSc.City = school.City
		oSc.EMail = school.EMail

		outSchools[index] = oSc
	}

	dataBytes, _ = json.Marshal(outSchools)
	ioutil.WriteFile(filepath.Join(jsonDir, "school.json"), dataBytes, 0600)

}

func httpConnect(url string, store interface{}) (err error) {

	connection := &http.Client{Transport: &http.Transport{}} // supports HTTP/2

	httpRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	httpRequest.Header.Set("Content-type", "application/json")

	var httpResponse *http.Response
	httpResponse, err = connection.Do(httpRequest)
	if err != nil {
		return // return the error
	}
	defer httpResponse.Body.Close()

	json.NewDecoder(httpResponse.Body).Decode(&store)

	return
}
