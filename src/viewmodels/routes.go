package viewmodels

import ()

type Routes struct {
	Title  string
	Active string
	Routes []Route
}

type Route struct {
	ImageUrl    string
	Title       string
	Description string
	Id          int
}

func GetRoutes() Routes {
	result := Routes{
		Title:  "My Running Routes",
		Active: "routes",
	}

	route1 := Route{
		Id:          1,
		ImageUrl:    "route1.png",
		Title:       "Indigo - Baldwin Park 5k",
		Description: "5K Route",
	}

	route2 := Route{
		Id:          2,
		ImageUrl:    "route2.png",
		Title:       "Home - Museum ",
		Description: "4K Route",
	}

	result.Routes = []Route{
		route1,
		route2,
	}
	return result
}
