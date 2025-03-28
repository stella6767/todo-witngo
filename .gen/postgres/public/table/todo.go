//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Todo = newTodoTable("public", "todo", "")

type todoTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnInteger
	Title     postgres.ColumnString
	Completed postgres.ColumnBool
	CreatedAt postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TodoTable struct {
	todoTable

	EXCLUDED todoTable
}

// AS creates new TodoTable with assigned alias
func (a TodoTable) AS(alias string) *TodoTable {
	return newTodoTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TodoTable with assigned schema name
func (a TodoTable) FromSchema(schemaName string) *TodoTable {
	return newTodoTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TodoTable with assigned table prefix
func (a TodoTable) WithPrefix(prefix string) *TodoTable {
	return newTodoTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TodoTable with assigned table suffix
func (a TodoTable) WithSuffix(suffix string) *TodoTable {
	return newTodoTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTodoTable(schemaName, tableName, alias string) *TodoTable {
	return &TodoTable{
		todoTable: newTodoTableImpl(schemaName, tableName, alias),
		EXCLUDED:  newTodoTableImpl("", "excluded", ""),
	}
}

func newTodoTableImpl(schemaName, tableName, alias string) todoTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		TitleColumn     = postgres.StringColumn("title")
		CompletedColumn = postgres.BoolColumn("completed")
		CreatedAtColumn = postgres.TimestampColumn("created_at")
		allColumns      = postgres.ColumnList{IDColumn, TitleColumn, CompletedColumn, CreatedAtColumn}
		mutableColumns  = postgres.ColumnList{TitleColumn, CompletedColumn, CreatedAtColumn}
	)

	return todoTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		Title:     TitleColumn,
		Completed: CompletedColumn,
		CreatedAt: CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
