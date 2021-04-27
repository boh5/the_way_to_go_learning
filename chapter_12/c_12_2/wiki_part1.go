package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	page := Page{
		"chapter_12/c_12_2/Page.md",
		[]byte("# Page\n## Section1\nThis is section1."),
	}
	page.save()

	var newPage Page
	newPage.load("chapter_12/c_12_2/Page.md")
	fmt.Println(string(newPage.Body))
}

func (p *Page) save() (err error) {
	return ioutil.WriteFile(p.Title, p.Body, 0666)
}

func (p *Page) load(title string) (err error) {
	p.Title = title
	p.Body, err = ioutil.ReadFile(title)
	return err
}
