// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Employee struct {
	ID         string `json:"_id"`
	Name       string `json:"Name"`
	IsTeamLead bool   `json:"IsTeamLead"`
}

type NewEmployee struct {
	Name       string `json:"Name"`
	IsTeamLead bool   `json:"IsTeamLead"`
}
