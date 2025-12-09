// Package generator provides code generation functionality for functional option patterns.
package generator

// StructField represents a field within a Go struct.
type StructField struct {
	// Name is field name in the structure
	Name string
	// Type is type of field in the structure
	Type string
	// Ignore is the flag to exclude 'Functional Optional Pattern' from the automatic generation target
	Ignore bool
}
