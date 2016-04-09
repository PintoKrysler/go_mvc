package controllers

import (
	"net/http"
	"text/template"
	"viewmodels"
)

type routesController struct {
	template *template.Template
}

func (this *routesController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetRoutes()
	w.Header().Add("Content-Type", "text/html")
	this.template.Execute(w, vm)
}
