package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mholt/binding"
)

// Contact is a DB record of a single snippet Contact
type Contact struct {
	CommonModel
	FirstName string
	LastName  string
	StartDate time.Time
	StatusID  uint
	Company   string
	Address   string
	Job       string
	Website   string
	Skype     string
	Phone     string
	Email     string
	Birthday  time.Time
	Photo     string
}

// FieldMap provides a field mapping
func (cf *Contact) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&cf.FirstName: "FirstName",
		&cf.LastName:  "LastName",
		&cf.StartDate: binding.Field{
			Form:       "StartDate",
			TimeFormat: "2006-01-02 15:04",
		},
		&cf.StatusID: "StatusID",
		&cf.Company:  "Company",
		&cf.Address:  "Address",
		&cf.Job:      "Job",
		&cf.Website:  "Website",
		&cf.Skype:    "Skype",
		&cf.Phone:    "Phone",
		&cf.Email:    "Email",
		&cf.Birthday: binding.Field{
			Form:       "Birthday",
			TimeFormat: "2006-01-02 15:04",
		},
		&cf.Photo: "Photo",
	}
}

func (d *Contact) MarshalJSON() ([]byte, error) {
	type Alias Contact
	return json.Marshal(&struct {
		*Alias
		Birthday  string
		StartDate string
	}{
		Alias:     (*Alias)(d),
		StartDate: d.StartDate.Format("02-01-2006"),
		Birthday:  d.Birthday.Format("02-01-2006"),
	})
}

// GetContact returns Contact by id
func (db *Context) GetContact(id string) *Contact {
	file := Contact{}
	db.First(&file, id)
	return &file
}

// GetAllContacts returns all snippets
func (db *Context) GetAllContacts() []Contact {
	files := []Contact{}
	db.Find(&files)
	return files
}

// DeleteContact deletes snippet from DB
func (db *Context) DeleteContact(id string) {
	file := db.GetContact(id)
	if file.ID != 0 {
		db.Delete(file)
	}
}

// UpdateContact deletes Contact from DB
func (db *Context) UpdateContact(id string, update *Contact) *Contact {
	file := db.GetContact(id)
	if file.ID != 0 {
		db.Model(file).Update(map[string]interface{}{
			"FirstName"	: update.FirstName,
			"LastName"	: update.LastName,
			"StartDate"	: update.StartDate,
			"StatusId"	: update.StatusID,
			"company"	: update.Company,
			"address"	: update.Address,
			"job"		: update.Job,
			"website"	: update.Website,
			"skype"		: update.Skype,
			"phone"		: update.Phone,
			"email"		: update.Email,
			"birthday"	: update.Birthday,
			"photo"		: update.Photo,
		})
	}
	return file
}

// AddContact deletes Contact from DB
func (db *Context) AddContact(update *Contact) *Contact {
	db.Save(&update)
	return update
}
