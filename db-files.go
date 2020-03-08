package main

import (
	"encoding/json"
	"github.com/mholt/binding"
	"net/http"
	"time"
)

type File struct {
	CommonModel
	Name string
	Size uint32
	UploadDate time.Time
	Path string
}

func (f* File) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.Name: "Name",
		&f.Size: "Size",
		&f.UploadDate: binding.Field{
			Form:       "UploadDate",
			TimeFormat: "2006-01-02 15:04",
		},
		&f.Path: "Path",
	}
}

// MarshalJSON is coming as native module
func (f *File) MarshalJSON() ([]byte, error) {
	type Alias File
	return json.Marshal(&struct {
		*Alias
		UploadDate string
	}{
		Alias:   (*Alias)(f),
		UploadDate: f.UploadDate.Format("2006-01-02 15:04"),
	})
}

func (db *Context) getFile(id string) *File {
	file := File{}
	db.First(&file,id)
	return &file
}
