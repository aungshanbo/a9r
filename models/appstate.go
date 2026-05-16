package models

type AppState struct {
	CurrentResource *Resource
	Filter          string
	ResourceType    string
}
