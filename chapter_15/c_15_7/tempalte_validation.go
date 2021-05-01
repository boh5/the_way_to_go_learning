package main

import (
	"fmt"
	"text/template"
)

func main() {
	tOk := template.New("ok")
	template.Must(tOk.Parse("/* and a comment */ some static text: {{.Name}}"))
	fmt.Println("the first one pared OK.")
	fmt.Println("the next one ought to fail.")
	tErr := template.New("error_tempalte")
	template.Must(tErr.Parse("some static text {{.Name}"))
}
