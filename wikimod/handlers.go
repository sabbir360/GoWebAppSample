package wikimod

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

// var templates = template.Must(template.ParseFiles(getTemplatePath("edit"),
// 	getTemplatePath("view"), getTemplatePath("css")))

var templates = template.Must(template.ParseGlob("./wikimod/views/*.html"))
var validPath = regexp.MustCompile("^/(edit|view|save)/([a-zA-Z0-9/-]+)$")
var wikiConfig = func() WikiConfig {
	//TODO:: Implement a method to read config from a config file.
	return WikiConfig{CDN: "http://localhost:5003/"}
}()

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Title/URL")
	}

	return m[2], nil
}

func getTemplatePath(name string) string {
	return "./wikimod/views/" + name + ".html"
}

//WikiTemplate return with whole HTML body
func WikiTemplate(w http.ResponseWriter, title string, body []byte) {

	resp := `<html><head><title>`
	resp = resp + title + `</title></head><body><h1>` + title + "</h1>"

	fmt.Fprintf(w, resp+"%s"+"</body></html>", body)

}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// t, err := template.ParseFiles(getTemplatePath(tmpl))
	p.Config = wikiConfig
	err := templates.ExecuteTemplate(w, tmpl, p)
	fmt.Println("Template loaded...")
	if err != nil {
		// WikiTemplate(w, p.Title, []byte(err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// err = t.Execute(w, p)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}

//MakeHandler is check and process to called handler. Its a wrapper in common word
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// //WikiEditTemplate to use edit form
// func WikiEditTemplate(w http.ResponseWriter, title string, body []byte, method string, action string) {
// 	resp := `<html><title>`
// 	resp = resp + title + `</title><body><h1>` + title + "</h1>"
// 	resp = resp + "<form type=" + method + "action='" + action + "'>"
// 	fmt.Fprintf(w, resp+"%s </form></body></html>", body)
// }

//Handler the home page
func Handler(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintf(w, "Hi There! I love %s!", r.URL.Path[1:])
	// body := "Hi There! I love " + r.URL.Path[1:] + "."
	WikiTemplate(w, r.URL.Path[1:], []byte("Hi There! I love "+r.URL.Path[1:]))
}

//ViewHandler The view page
func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Println("view...")
	p, err := LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	fmt.Println("page loaded...")
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	// WikiTemplate(w, p.Title, p.Body)
	renderTemplate(w, "view", p)
}

//EditHandler This edit page
func EditHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := LoadPage(title)
	if err != nil {
		// WikiTemplate(w, title, []byte(err.Error()))
		p = &Page{Title: title}

	}

	renderTemplate(w, "edit", p)
}

// SaveHandler to save wiki pages
func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {

	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(template.HTMLEscapeString(body))}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
