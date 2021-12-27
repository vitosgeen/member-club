package member

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"net/mail"
	"regexp"
	"strings"
	"time"
)

func MemberHttpHandlers() {
	http.HandleFunc("/", memberHandlerIndex)
	http.HandleFunc("/member/add", memberHandlerAdd)
}

func memberHandlerIndex(w http.ResponseWriter, r *http.Request) {
	templateNew := template.Must(template.New("index.html").Funcs(template.FuncMap{
		"incrementMemberTpl":  incrementMemberTpl,
		"formatDateMemberTpl": formatDateMemberTpl,
	}).ParseFiles("templates/index.html"))
	err := templateNew.ExecuteTemplate(w, "index.html", Members)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusBadRequest)
	}
}

func memberHandlerAdd(w http.ResponseWriter, r *http.Request) {
	var email string
	var name string
	var err error
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	r.ParseForm()
	err = memberHandlerAddValidationForm(r.Form)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error: "+err.Error(), http.StatusBadRequest)
	} else {
		name = strings.Join(r.Form["memberName"], "")
		email = strings.Join(r.Form["memberEmail"], "")
		newMember := Member{
			Name:  name,
			Email: email,
			Date:  time.Now(),
		}
		err = addMember(newMember)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func memberHandlerAddValidationForm(formData map[string][]string) error {
	var err error
	log.Println(formData["memberName"], formData["memberEmail"])
	if _, ok := formData["memberEmail"]; !ok {
		err = errors.New("Empty Email")
		return err
	}
	if _, ok := formData["memberName"]; !ok {
		err = errors.New("Empty Name")
		return err
	}
	_, err = mail.ParseAddress(strings.Join(formData["memberEmail"], ""))
	if err != nil {
		return err
	}
	nameRegex := regexp.MustCompile("^[a-zA-Z0-9 ]+$")
	isValid := nameRegex.MatchString(strings.Join(formData["memberName"], ""))
	if !isValid {
		err = errors.New("Name is not valid")
	}
	return err
}

func incrementMemberTpl(cnt int) int {
	return (cnt + 1)
}

func formatDateMemberTpl(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
