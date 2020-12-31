// Code generated by entc, DO NOT EDIT.

package diagnosis

const (
	// Label holds the string label denoting the diagnosis type in the database.
	Label = "diagnosis"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDiagnosticMessages holds the string denoting the diagnosticmessages field in the database.
	FieldDiagnosticMessages = "diagnostic_messages"
	// FieldSurveillancePeriod holds the string denoting the surveillanceperiod field in the database.
	FieldSurveillancePeriod = "surveillance_period"
	// FieldDiagnosisDate holds the string denoting the diagnosisdate field in the database.
	FieldDiagnosisDate = "diagnosis_date"

	// EdgeDisease holds the string denoting the disease edge name in mutations.
	EdgeDisease = "disease"
	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"
	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"

	// Table holds the table name of the diagnosis in the database.
	Table = "diagnoses"
	// DiseaseTable is the table the holds the disease relation/edge.
	DiseaseTable = "diseases"
	// DiseaseInverseTable is the table name for the Disease entity.
	// It exists in this package in order to avoid circular dependency with the "disease" package.
	DiseaseInverseTable = "diseases"
	// DiseaseColumn is the table column denoting the disease relation/edge.
	DiseaseColumn = "disease_diagnosis"
	// PatientTable is the table the holds the patient relation/edge.
	PatientTable = "diagnoses"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "patient_diagnosis"
	// EmployeeTable is the table the holds the employee relation/edge.
	EmployeeTable = "diagnoses"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "employee_diagnosis"
)

// Columns holds all SQL columns for diagnosis fields.
var Columns = []string{
	FieldID,
	FieldDiagnosticMessages,
	FieldSurveillancePeriod,
	FieldDiagnosisDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Diagnosis type.
var ForeignKeys = []string{
	"employee_diagnosis",
	"patient_diagnosis",
}