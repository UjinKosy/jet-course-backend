package main

import (
	"net/http"

	"github.com/mholt/binding"

	"github.com/go-chi/chi"
	"github.com/unrolled/render"
)

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

}
