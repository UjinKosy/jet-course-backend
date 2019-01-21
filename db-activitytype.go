package main

import (
	"net/http"

	"github.com/mholt/binding"
)

// ActivityType is a DB record of a single snippet ActivityType
type ActivityType struct {
	CommonModel
	Value string
	Icon  string
}

// FieldMap provides a field mapping
func (cf *ActivityType) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&cf.Value: "Value",
		&cf.Icon:  "Icon",
	}
}

// GetActivityType returns ActivityType by id
func (db *Context) GetActivityType(id string) *ActivityType {
	file := ActivityType{}
	db.First(&file, id)
	return &file
}

// GetAllActivityTypes returns all snippets
func (db *Context) GetAllActivityTypes() []ActivityType {
	files := []ActivityType{}
	db.Find(&files)
	return files
}

// DeleteActivityType deletes snippet from DB
func (db *Context) DeleteActivityType(id string) {
	file := db.GetActivityType(id)
	if file.ID != 0 {
		db.Delete(file)
	}
}

// UpdateActivityType deletes Activity from DB
func (db *Context) UpdateActivityType(id string, update *ActivityType) *ActivityType {
	file := db.GetActivityType(id)
	if file.ID != 0 {
		db.Model(file).Update(map[string]interface{}{
			"value"	: update.Value,
			"icon"	: update.Icon,
		})
	}
	return file
}

// AddActivityType deletes Activity from DB
func (db *Context) AddActivityType(update *ActivityType) *ActivityType {
	db.Save(&update)
	return update
}
