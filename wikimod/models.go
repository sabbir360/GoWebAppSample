package wikimod

import "io/ioutil"

//Page Attribute
type Page struct {
	Title  string
	Body   []byte
	Config WikiConfig
}

// WikiConfig Attribute
type WikiConfig struct {
	CDN string
}

//Save wiki page
func (p *Page) Save() error {
	filename := "data/" + p.Title + ".txt"

	return ioutil.WriteFile(filename, p.Body, 0600)
}

//LoadPage wikipage
func LoadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
