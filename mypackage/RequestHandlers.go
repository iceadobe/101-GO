package mypackage

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

const (
	VIEW_PATH string = "resource/views/"
	PAGE_PATH string = "resource/pages/"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename := PAGE_PATH + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := PAGE_PATH + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(VIEW_PATH + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	/*	title := r.URL.Path[len("/views/"):]
		p, ok := LoadPage(title)
		if ok != nil {
			errorPage := Page{"Error Code 404", []byte("Page Not Found")}
			fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", errorPage.Title, errorPage.Body)
		} else {
			fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
		}
	*/
	title := r.URL.Path[len("/view/"):]
	p, err := LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
		p.Save()
	}
	renderTemplate(w, "edit", p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
