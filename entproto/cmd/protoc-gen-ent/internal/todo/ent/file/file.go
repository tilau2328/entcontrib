// Code generated by ent, DO NOT EDIT.

package file

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContents holds the string denoting the contents field in the database.
	FieldContents = "contents"
	// Table holds the table name of the file in the database.
	Table = "files"
)

// Columns holds all SQL columns for file fields.
var Columns = []string{
	FieldID,
	FieldContents,
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

// Order defines the ordering method for the File queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByContents orders the results by the contents field.
func ByContents(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldContents, opts...).ToFunc()
}
