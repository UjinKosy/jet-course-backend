package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"os"
	"path/filepath"
	"strings"

	"github.com/mholt/binding"

	"github.com/go-chi/chi"
	"github.com/unrolled/render"
)

type uploaderResponse struct {
	ID uint `json:"id,omitempty"`
	Status string
	Name string 
	SName string
}

func apiRoutes(r *chi.Mux, dbc *Context, render *render.Render, cfg *AppConfig) {

	apiRoot := cfg.Server.Path + "api/v1/"

	// contacts
	contactsAPI := apiRoot + "contacts"
	r.Get(contactsAPI+"/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, 200, dbc.GetAllContacts())
	})
	r.Get(contactsAPI+"/{key}", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, 200, dbc.GetContact(chi.URLParam(r, "key")))
	})
	r.Post(contactsAPI+"/", func(w http.ResponseWriter, r *http.Request) {
		contact := new(Contact)
		binding.Bind(r, contact)

		contact = dbc.AddContact(contact)
		render.JSON(w, 200, response{ID: contact.ID})
	})
	r.Put(contactsAPI+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		contact := new(Contact)
		binding.Bind(r, contact)

		contact = dbc.UpdateContact(id, contact)
		render.JSON(w, 200, response{ID: contact.ID})
	})
	r.Delete(contactsAPI+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		dbc.DeleteContact(chi.URLParam(r, "id"))
		render.JSON(w, 200, response{})
	})

	// contacts
	countriesAPI := apiRoot + "countries"
	r.Get(countriesAPI+"/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, 200, dbc.GetAllCountrys())
	})
	r.Get(countriesAPI+"/{key}", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, 200, dbc.GetCountry(chi.URLParam(r, "key")))
	})
	r.Post(countriesAPI+"/", func(w http.ResponseWriter, r *http.Request) {
		contact := new(Country)
		binding.Bind(r, contact)

		contact = dbc.AddCountry(contact)
		render.JSON(w, 200, response{ID: contact.ID})
	})
	r.Put(countriesAPI+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		contact := new(Country)
		binding.Bind(r, contact)

		contact = dbc.UpdateCountry(id, contact)
		render.JSON(w, 200, response{ID: contact.ID})
	})
	r.Delete(countriesAPI+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		dbc.DeleteCountry(chi.URLParam(r, "id"))
		render.JSON(w, 200, response{})
	})

	// status
	statusesAPI := apiRoot + "statuses"
	r.Get(statusesAPI+"/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, 200, dbc.GetAllStatuses())
	})
	r.Get(statusesAPI+"/{key}", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, 200, dbc.GetStatus(chi.URLParam(r, "key")))
	})
	r.Post(statusesAPI+"/", func(w http.ResponseWriter, r *http.Request) {
		contact := new(Status)
		binding.Bind(r, contact)

		contact = dbc.AddStatus(contact)
		render.JSON(w, 200, response{ID: contact.ID})
	})
	r.Put(statusesAPI+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		contact := new(Status)
		binding.Bind(r, contact)

		contact = dbc.UpdateStatus(id, contact)
		render.JSON(w, 200, response{ID: contact.ID})
	})
	r.Delete(statusesAPI+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		dbc.DeleteStatus(chi.URLParam(r, "id"))
		render.JSON(w, 200, response{})
	})

	// status
	activitiesAPI := apiRoot + "activities"
	r.Get(activitiesAPI+"/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, 200, dbc.GetAllActivities())
	})
	r.Get(activitiesAPI+"/{key}", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, 200, dbc.GetActivity(chi.URLParam(r, "key")))
	})
	r.Post(activitiesAPI+"/", func(w http.ResponseWriter, r *http.Request) {
		contact := new(Activity)
		binding.Bind(r, contact)

		contact = dbc.AddActivity(contact)
		render.JSON(w, 200, response{ID: contact.ID})
	})
	r.Put(activitiesAPI+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		contact := new(Activity)
		binding.Bind(r, contact)

		contact = dbc.UpdateActivity(id, contact)
		render.JSON(w, 200, response{ID: contact.ID})
	})
	r.Delete(activitiesAPI+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		dbc.DeleteActivity(chi.URLParam(r, "id"))
		render.JSON(w, 200, response{})
	})

	// status
	activityTypesAPI := apiRoot + "activitytypes"
	r.Get(activityTypesAPI+"/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, 200, dbc.GetAllActivityTypes())
	})
	r.Get(activityTypesAPI+"/{key}", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, 200, dbc.GetActivityType(chi.URLParam(r, "key")))
	})
	r.Post(activityTypesAPI+"/", func(w http.ResponseWriter, r *http.Request) {
		contact := new(ActivityType)
		binding.Bind(r, contact)

		contact = dbc.AddActivityType(contact)
		render.JSON(w, 200, response{ID: contact.ID})
	})
	r.Put(activityTypesAPI+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		contact := new(ActivityType)
		binding.Bind(r, contact)

		contact = dbc.UpdateActivityType(id, contact)
		render.JSON(w, 200, response{ID: contact.ID})
	})
	r.Delete(activityTypesAPI+"/{id}", func(w http.ResponseWriter, r *http.Request) {
		dbc.DeleteActivityType(chi.URLParam(r, "id"))
		render.JSON(w, 200, response{})
	})

	// upload
	uploadAPI := apiRoot + "uploadFiles"

	r.Route(uploadAPI, func(root chi.Router) {
		root.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hi"))
		})

		root.Post(uploadAPI+"/", func(w http.ResponseWriter, r *http.Request) {
			// Parse our multipart form, 10 << 20 specifies a maximum
			// upload of 10 MB files.
			r.ParseMultipartForm(10 << 20)
			// FormFile returns the first file for the given key `myFile`
			// it also returns the FileHeader so we can get the Filename,
			// the Header and the size of the file
			file, handler, err := r.FormFile("upload")
			if err != nil {
				fmt.Println("Error Retrieving the File")
				fmt.Println(err)
				render.JSON(w, 200, uploaderResponse{Status: "error"})
				return
			}
			defer file.Close()
			fmt.Printf("Uploaded File: %+v\n", handler.Filename)
			fmt.Printf("File Size: %+v\n", handler.Size)
			fmt.Printf("MIME Header: %+v\n", handler.Header)
	
			// Create a temporary file within our temp-images directory that follows
			// a particular naming pattern
			fileNewName := "upload-"+handler.Filename+".xxx"
			tempFile, err := ioutil.TempFile("temp-files", fileNewName)
			if err != nil {
				fmt.Println(err)
				render.JSON(w, 200, uploaderResponse{Status: "error"})
			}
			fmt.Printf("New name of the file is %+v\n", tempFile.Name())
			defer tempFile.Close()
	
			// read all of the contents of our uploaded file into a
			// byte array
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
				render.JSON(w, 200, uploaderResponse{Status: "error"})
			}
			// write this byte array to our temporary file
			tempFile.Write(fileBytes)
	
			render.JSON(w, 200, uploaderResponse{
				Status: "server",
				Name: handler.Filename,
				SName: tempFile.Name(),
			})
		})

		workDir, _ := os.Getwd()
		filesDir := filepath.Join(workDir, "temp-files")
		FileServer(root, uploadAPI, "/files", http.Dir(filesDir))
	})

}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, basePath string, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(basePath+path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}