package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func Home(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		Contact(w, r)
		return
	}

	page, err := template.ParseFiles("views/index.tpl")
	if err != nil {
		log.Fatal("ParseFiles:", err)
	}

	if err := page.Execute(w, nil); err != nil {
		log.Fatal("Execute:", err)
	}
}

func Contact(w http.ResponseWriter, r *http.Request) {

	m := gomail.NewMessage()
	m.SetAddressHeader("From", r.FormValue("email"), r.FormValue("name"))
	m.SetAddressHeader("Reply-To", r.FormValue("email"), r.FormValue("name"))
	m.SetAddressHeader("To", "parikshitgothwal14@gmail.com", "Parikshit Gothwal")
	m.SetHeader("Subject", "ContactForm: "+r.FormValue("subject"))
	m.SetBody("text/plain", r.FormValue("message"))

	// d := gomail.NewDialer("smtp.fastmail.com", 465, "parikshitgothwal14@gmail.com", "")
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		log.Println(m)
	}

	page, err := template.ParseFiles("views/thankyou.tpl")
	if err != nil {
		log.Fatal("ParseFiles:", err)
	}

	if err := page.Execute(w, nil); err != nil {
		log.Fatal("Execute:", err)
	}
}
