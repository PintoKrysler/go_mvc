package viewmodels

type Home struct {
	Title  string
	Active string
}

func GetHome() Home {
	result := Home{
		Title:  "Krysler Pinto Routes",
		Active: "home",
	}
	return result
}
