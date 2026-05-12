package models

type Resource struct {
	Name    string
	Headers []TableColumn
	Rows    [][]string
}
