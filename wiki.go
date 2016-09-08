package main

//Ref: https://golang.org/doc/articles/wiki/#tmp_14

import "net/http"
import "./wikimod"

func main() {
	// p1 := &Page{"TestPage", []byte("This is a test page.")}
	// p1.save()

	// p2, _ := loadPage("TestPage")

	// fmt.Println(string(p2.Body))

	http.HandleFunc("/", wikimod.Handler)
	http.HandleFunc("/view/", wikimod.MakeHandler(wikimod.ViewHandler))
	http.HandleFunc("/edit/", wikimod.MakeHandler(wikimod.EditHandler))
	http.HandleFunc("/save/", wikimod.MakeHandler(wikimod.SaveHandler))
	http.ListenAndServe(":8001", nil)
}
