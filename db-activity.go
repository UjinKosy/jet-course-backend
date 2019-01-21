package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mholt/binding"
)

// Activity is a DB record of a single snippet Activity
type Activity struct {
	CommonModel
	Details   string
	TypeID    uint
	State     string
	ContactID uint
	DueDate   time.Time
}

// FieldMap provides a field mapping
func (cf *Activity) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&cf.Details:   "Details",
		&cf.State:     "State",
		&cf.TypeID:    "TypeID",
		&cf.ContactID: "ContactID",
		&cf.DueDate: binding.Field{
			Form:       "DueDate",
			TimeFormat: "2006-01-02 15:04",
		},
	}
}

func (d *Activity) MarshalJSON() ([]byte, error) {
	type Alias Activity
	return json.Marshal(&struct {
		*Alias
		DueDate string
	}{
		Alias:   (*Alias)(d),
		DueDate: d.DueDate.Format("02-01-2006 15:04"),
	})
}

// GetActivity returns Activity by id
func (db *Context) GetActivity(id string) *Activity {
	file := Activity{}
	db.First(&file, id)
	return &file
}

// GetAllActivities returns all snippets
func (db *Context) GetAllActivities() []Activity {
	files := []Activity{}
	db.Find(&files)
	return files
}

// DeleteActivity deletes snippet from DB
func (db *Context) DeleteActivity(id string) {
	file := db.GetActivity(id)
	if file.ID != 0 {
		db.Delete(file)
	}
}

// UpdateActivity deletes Activity from DB
func (db *Context) UpdateActivity(id string, update *Activity) *Activity {
	file := db.GetActivity(id)
	if file.ID != 0 {
		db.Model(file).Update(map[string]interface{}{
			"details"	: update.Details,
			"TypeId"	: update.TypeID,
			"state"		: update.State,
			"ContactId": update.ContactID,
			"DueDate"	: update.DueDate,
		})
	}
	return file
}

// AddActivity deletes Activity from DB
func (db *Context) AddActivity(update *Activity) *Activity {
	db.Save(&update)
	return update
}
