// Code generated by ent, DO NOT EDIT.

package codesourcefile

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the codesourcefile type in the database.
	Label = "code_source_file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// EdgeDepsUsageEvidences holds the string denoting the deps_usage_evidences edge name in mutations.
	EdgeDepsUsageEvidences = "deps_usage_evidences"
	// Table holds the table name of the codesourcefile in the database.
	Table = "code_source_files"
	// DepsUsageEvidencesTable is the table that holds the deps_usage_evidences relation/edge.
	DepsUsageEvidencesTable = "deps_usage_evidences"
	// DepsUsageEvidencesInverseTable is the table name for the DepsUsageEvidence entity.
	// It exists in this package in order to avoid circular dependency with the "depsusageevidence" package.
	DepsUsageEvidencesInverseTable = "deps_usage_evidences"
	// DepsUsageEvidencesColumn is the table column denoting the deps_usage_evidences relation/edge.
	DepsUsageEvidencesColumn = "deps_usage_evidence_used_in"
)

// Columns holds all SQL columns for codesourcefile fields.
var Columns = []string{
	FieldID,
	FieldPath,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// PathValidator is a validator for the "path" field. It is called by the builders before save.
	PathValidator func(string) error
)

// OrderOption defines the ordering options for the CodeSourceFile queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// ByDepsUsageEvidencesCount orders the results by deps_usage_evidences count.
func ByDepsUsageEvidencesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDepsUsageEvidencesStep(), opts...)
	}
}

// ByDepsUsageEvidences orders the results by deps_usage_evidences terms.
func ByDepsUsageEvidences(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDepsUsageEvidencesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDepsUsageEvidencesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DepsUsageEvidencesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, DepsUsageEvidencesTable, DepsUsageEvidencesColumn),
	)
}
