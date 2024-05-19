package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type PeopleInfo struct {
	Person_ID int
	Email     string
	Address   string
	Job       string
}

var people = []PeopleInfo{
	{Person_ID: 1, Email: "emma.johnson@example.com", Address: "23 Maple Street, Anytown, USA", Job: "Marketing Coordinator"},
	{Person_ID: 2, Email: "james.anderson@example.com", Address: "456 Oak Avenue, Springfield, USA", Job: "Software Developer"},
	{Person_ID: 3, Email: "j.smith@example.com", Address: "789 Oak Street, Cityville, USA", Job: "Software Engineer"},
	{Person_ID: 4, Email: "l.wilson@example.com", Address: "456 Pine Avenue, Suburbia, USA", Job: "Marketing Manager"},
	{Person_ID: 5, Email: "m.jones@example.org", Address: "123 Maple Road, Townsville, USA", Job: "Financial Analyst"},
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/profile/", profileHandler)
	http.HandleFunc("/usernotfound", unregisteredUserHandler)
	http.HandleFunc("/redirect", backLogoutHandler)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("useremail")

	for _, person := range people {
		if username == person.Email {
			person_id := strconv.Itoa(person.Person_ID)
			http.Redirect(w, r, "/profile/"+person_id, http.StatusSeeOther)
		}
	}

	http.Redirect(w, r, "/usernotfound", http.StatusSeeOther)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	person_id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, person := range people {
		if person_id == person.Person_ID {
			tmpl, err := template.ParseFiles("templates/profile.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, person)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func unregisteredUserHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Email is not registered!"

	tmpl, err := template.ParseFiles("templates/unregistered.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func backLogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
