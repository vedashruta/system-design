package main

import (
	"fmt"

	"github.com/google/uuid"
)

var counter = 1

type Table struct {
	ID      string
	Name    string
	Columns []Column
	Rows    map[int][]Row
}

type Column struct {
	ID         string
	TableID    string
	ColumnName string
	DataType   string
}

type Row struct {
	ID         string
	ColumnName string
	DataType   string
	Value      any
}

func NewTable(tableName string) (table Table) {
	table = Table{
		ID:      uuid.NewString(),
		Name:    tableName,
		Columns: []Column{},
		Rows:    map[int][]Row{},
	}
	return
}

func (table *Table) AddColumn(columnName string, dataType string) {
	col := Column{
		ID:         uuid.NewString(),
		TableID:    table.ID,
		ColumnName: columnName,
		DataType:   dataType,
	}
	table.Columns = append(table.Columns, col)
}

func (table *Table) InsertRow(rows []Row) {
	for i := range rows {
		if !(rows[i].ColumnName == table.Columns[i].ColumnName && rows[i].DataType == table.Columns[i].DataType) {
			return
		}
	}
	table.Rows[counter] = append(table.Rows[counter], rows...)
	counter++
}

func (table *Table) Print() {
	res := make(map[string]any)
	for _, rows := range table.Rows {
		for _, row := range rows {
			res[row.ColumnName] = row.Value
		}
	}
	fmt.Println(res)
}

func main() {
	table := NewTable("sample_table")
	columnName := "name"
	dataType := "string"
	table.AddColumn(columnName, dataType)
	columnName = "age"
	dataType = "integer"
	table.AddColumn(columnName, dataType)
	rows := []Row{
		Row{
			ColumnName: "name",
			Value:      "abcde",
			DataType:   "string",
		},
		Row{
			ColumnName: "age",
			Value:      10,
			DataType:   "integer",
		},
	}
	table.InsertRow(rows)
	table.Print()
}
