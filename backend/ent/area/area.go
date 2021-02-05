// Code generated by entc, DO NOT EDIT.

package area

const (
	// Label holds the string label denoting the area type in the database.
	Label = "area"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAreaName holds the string denoting the areaname field in the database.
	FieldAreaName = "area_name"
	// FieldAreaDistrict holds the string denoting the areadistrict field in the database.
	FieldAreaDistrict = "area_district"
	// FieldAreaSubDistrict holds the string denoting the areasubdistrict field in the database.
	FieldAreaSubDistrict = "area_sub_district"

	// EdgeDisease holds the string denoting the disease edge name in mutations.
	EdgeDisease = "disease"
	// EdgeStatistic holds the string denoting the statistic edge name in mutations.
	EdgeStatistic = "statistic"
	// EdgeLevel holds the string denoting the level edge name in mutations.
	EdgeLevel = "level"
	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"

	// Table holds the table name of the area in the database.
	Table = "areas"
	// DiseaseTable is the table the holds the disease relation/edge.
	DiseaseTable = "areas"
	// DiseaseInverseTable is the table name for the Disease entity.
	// It exists in this package in order to avoid circular dependency with the "disease" package.
	DiseaseInverseTable = "diseases"
	// DiseaseColumn is the table column denoting the disease relation/edge.
	DiseaseColumn = "disease_area"
	// StatisticTable is the table the holds the statistic relation/edge.
	StatisticTable = "areas"
	// StatisticInverseTable is the table name for the Statistic entity.
	// It exists in this package in order to avoid circular dependency with the "statistic" package.
	StatisticInverseTable = "statistics"
	// StatisticColumn is the table column denoting the statistic relation/edge.
	StatisticColumn = "statistic_area"
	// LevelTable is the table the holds the level relation/edge.
	LevelTable = "areas"
	// LevelInverseTable is the table name for the Level entity.
	// It exists in this package in order to avoid circular dependency with the "level" package.
	LevelInverseTable = "levels"
	// LevelColumn is the table column denoting the level relation/edge.
	LevelColumn = "level_area"
	// EmployeeTable is the table the holds the employee relation/edge.
	EmployeeTable = "areas"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "employee_area"
)

// Columns holds all SQL columns for area fields.
var Columns = []string{
	FieldID,
	FieldAreaName,
	FieldAreaDistrict,
	FieldAreaSubDistrict,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Area type.
var ForeignKeys = []string{
	"disease_area",
	"employee_area",
	"level_area",
	"statistic_area",
}

var (
	// AreaNameValidator is a validator for the "AreaName" field. It is called by the builders before save.
	AreaNameValidator func(string) error
	// AreaDistrictValidator is a validator for the "AreaDistrict" field. It is called by the builders before save.
	AreaDistrictValidator func(string) error
	// AreaSubDistrictValidator is a validator for the "AreaSubDistrict" field. It is called by the builders before save.
	AreaSubDistrictValidator func(string) error
)
