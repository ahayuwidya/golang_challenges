package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// User represents a user in the system
type User struct {
	Username string
	Password string
}

var users = []User{
	{"user1", "password1"},
	{"user2", "password2"},
	// Add more users as needed
}

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
	http.HandleFunc("/profile/", profilePage)
	http.HandleFunc("/gagal", gagallogin)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// username := r.FormValue("username")
	username := r.FormValue("useremail")
	// password := r.FormValue("password")

	fmt.Println("username test", username)
	// fmt.Println("password test", password)
	// fmt.Println("test loop")

	for i, person := range people {
		fmt.Println("here", i, person.Email)
		if username == person.Email {
			fmt.Println("here 1")
			fmt.Println("username check match", username)
			person_id := strconv.Itoa(person.Person_ID)
			fmt.Println("person_id check match", person_id)
			http.Redirect(w, r, "/profile/"+person_id, http.StatusSeeOther)
			fmt.Println("here 2")
		}
	}
	fmt.Println("here 3")
	fmt.Println("username check no match", username)
	http.Redirect(w, r, "/gagal", http.StatusSeeOther)
	fmt.Println("here 4")
}

func profilePage(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	person_id, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// msg := "Profile Page"
	for _, person := range people {
		if person_id == person.Person_ID {
			// fmt.Fprint(w, msg, person) // structnya "person"
			tmpl, err := template.ParseFiles("templates/profile.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Execute the template with the data and write the output to the response
			err = tmpl.Execute(w, person)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func gagallogin(w http.ResponseWriter, r *http.Request) {
	msg := "Email is not registered!"
	// fmt.Fprint(w, msg)
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
