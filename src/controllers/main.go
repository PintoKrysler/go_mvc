package controllers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func Register(templates *template.Template) {
	// http.HandleFunc("/",
	// 	func(w http.ResponseWriter, req *http.Request) {
	// 		requestedFiled := req.URL.Path[1:]
	// 		template := templates.Lookup(requestedFiled + ".html")

	// 		var context interface{} = nil

	// 		switch requestedFiled {
	// 		case "home":
	// 			context = viewmodels.GetHome()
	// 		case "routes":
	// 			context = viewmodels.GetRoutes()
	// 		}

	// 		if template != nil {
	// 			template.Execute(w, context)
	// 		} else {
	// 			w.WriteHeader(404)
	// 		}
	// 	})

	home_controller := new(homeController)
	home_controller.template = templates.Lookup("home.html")
	http.HandleFunc("/home", home_controller.get)

	routes_controller := new(routesController)
	routes_controller.template = templates.Lookup("routes.html")
	http.HandleFunc("/routes", routes_controller.get)

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	fmt.Printf(path)
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}
	fmt.Printf(contentType)
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
