# MySchoolList

Inspired by [Node MySchoolList](https://github.com/apitlekays/MySchoolList)

Library for fetching list of schools in Malaysia based on States & District.

Data taken from Ministry of Education Malaysia (Date release: January 2018)

## Using this package

Install this package using 

```bash
go get github.com/svicknesh/myschoollist
```

Import it into your appliation

```go
import "github.com/svicknesh/myschoollist"
```

There are no additional dependencies used by this package aside from what is available in Go standard.

The `list_test.go` file contains examples of using this package. 

The information about the states, districts and schools are in the `json` folder which pulls the information from [Node MySchoolList](https://github.com/apitlekays/MySchoolList) before cleaning it up a little. To rebuild the information, perform the following

```bash
cd json
go run import.go
```


### Instantiate new instance of the list

```go
list, err := myschoollist.New()
if nil != err {
    log.Panic(err)
}
```

**NOTE**: Try not to instantiate too many instances of this list, instead instantiate it once and pass it to other modules or functions as parameters. Every instantiation loads the JSON information and creates mapping between them for quick lookup, which in turn uses memory. Having multiple instances *MAY* impact memory usage in long running applications for each instantiation.


### State information

#### Get state information by ID

```go
stateID := 7
state, found := list.StateGetByID(stateID)
if !found {
    log.Printf("State ID %d not found", stateID)
}
```

#### Get state information by short name

```go
stateShortname := "pen" // case insensitive
state, found := list.StateGetByShortname(stateShortname)
if !found {
    log.Printf("State ID %d not found", stateID)
}
```

#### Get all the states information in Malaysia

```go
states := list.StatesGetAll()
```


### District information

#### Get district information by ID

```go
districtID := 73
district, found := list.DistrictGetByID(districtID)
if !found {
    log.Printf("District ID %d not found", districtID)
}
```

#### Get list of schools information in a given district ID

```go
schools := list.DistrictGetSchools(districtID)
```

#### Get state information for a given district ID

```go
districtID := 73
state, found = list.DistrictGetState(districtID)
if !found {
    log.Printf("District ID %d not found", districtID)
}
```

### School information

#### Get school information by ID

```go
schoolID := 9429
school, found := list.SchoolGetByID(schoolID)
if !found {
    log.Printf("School ID %d not found", schoolID)
}
```

#### Get school information by code

```go
schoolCode := "peb1094"
school, found = list.SchoolGetByCode(schoolCode)
if !found {
    log.Printf("School code %q not found", schoolCode)
}
```

