package main

import (
	"controllers"
	"net/http"
	"os"
	"text/template"
)

func main() {
	templates := populateTemplates()
	controllers.Register(templates)
	http.ListenAndServe(":8000", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates")
	basePath := "templates"
	templateFolder, _ := os.Open("templates")
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)
	for _, pathinfo := range templatePathsRaw {
		if !pathinfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathinfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)
	return result
}
