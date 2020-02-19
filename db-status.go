package main

import (
	"net/http"

	"github.com/mholt/binding"
)

// Status is a DB record of a single snippet Status
type Status struct {
	CommonModel
	Value string
	Icon  string
}

// FieldMap provides a field mapping
func (cf *Status) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&cf.Value: "Value",
		&cf.Icon:  "Icon",
	}
}

// GetStatus returns Status by id
func (db *Context) GetStatus(id string) *Status {
	file := Status{}
	db.First(&file, id)
	return &file
}

// GetAllStatuses returns all snippets
func (db *Context) GetAllStatuses() []Status {
	files := []Status{}
	db.Find(&files)
	return files
}

// DeleteStatus deletes snippet from DB
func (db *Context) DeleteStatus(id string) {
	file := db.GetStatus(id)
	if file.ID != 0 {
		db.Delete(file)
	}
}

// UpdateStatus deletes Status from DB
func (db *Context) UpdateStatus(id string, update *Status) *Status {
	file := db.GetStatus(id)
	if file.ID != 0 {
		db.Model(file).Update(map[string]interface{}{"value": update.Value, "icon": update.Icon})
	}
	return file
}

// AddStatus deletes Status from DB
func (db *Context) AddStatus(update *Status) *Status {
	db.Save(&update)
	return update
}
