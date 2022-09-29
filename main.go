package main

import (
	"html/template"
	"net/http"
)

type ContactUs struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			err := tmpl.Execute(w, nil)
			if err != nil {
				return
			}
			return
		}

		details := ContactUs{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		err := tmpl.Execute(w, details)
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":7575", nil)
	if err != nil {
		return
	}
}
