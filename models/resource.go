package models

type Resource struct {
	Name    string
	Headers []TableColumn
	Rows    [][]string
}

type TableColumn struct {
	Title     string
	Expansion int
}

type AppState struct {
	CurrentResource *Resource
	Filter          string
	ResourceType    string
	EC2Instances    []Ec2instance
}
