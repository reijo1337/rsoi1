package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func errHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>404, no such file</h1>")
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/404/", errHandler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
