package standard

import (
	api "github.com/CoachApplication/coach-api"
)

// Standard Standard parametrized implementation of a UI that has some merging capability
type Ui struct {
	id string
	label string
	description string
	help string
}

// NewUi Create a new Ui from parametrized strings
func NewUi(id, label, description, help string) *Ui {
	return &Ui{
		id: id,
		label: label,
		description: description,
		help: help,
	}
}

// Ui convert this to a Ui interface explicitly
func (su *Ui) Ui() api.Ui {
	return api.Ui(su)
}

// Merge merge a ui into this one
func (su *Ui) Merge(merge api.Ui) {
	if val := merge.Id(); val != "" {
		su.id = val
	}
	if val := merge.Label(); val != "" {
		su.label = val
	}
	if val := merge.Description(); val != "" {
		su.description = val
	}
	if val := merge.Help(); val != "" {
		su.help = val
	}
}

// Id get the string Id for the Ui
func (su *Ui) Id() string {
	return su.id
}

// Description get the string Description
func (su *Ui) Label() string {
	return su.label
}

// Label get the string label
func (su *Ui) Description() string {
	return su.description
}

// Help get the string help
func (su *Ui) Help() string {
	return su.help
}
