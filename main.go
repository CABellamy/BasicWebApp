package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Title,
	Content string
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "Home",
	}
	t := template.Must(
		template.ParseFiles(
			"public/index.html",
			"public/header.html",
			"public/home.html",
			"public/nav.html",
			"public/footer.html"))
	contentType := "text/html"
	w.Header().Add("Content-Type", contentType)
	t.Execute(w, p)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "Login",
	}
	t := template.Must(template.ParseFiles("public/login.html", "public/header.html", "public/nav.html", "public/footer.html"))
	contentType := "text/html"
	w.Header().Add("Content-Type", contentType)
	t.Execute(w, p)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	email := r.FormValue("email")
	password := r.FormValue("password")
	if email == "testaccount@test.com" && password == "SuperSecretPassword" {
		w.Write([]byte("Login Success"))
	} else {
		w.Write([]byte("Login Failed"))
	}
}
