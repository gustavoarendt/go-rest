package models

type Persona struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Lore string `json:"lore"`
}

var Personas []Persona
