package backgo

type List struct {
	Lists []artists
}

type artists struct {
	Id		int			`json:"id"`
	Image	string		`json:"image"`
	Name	string		`json:"name"`
	Members	[]string	`json:"members"`
	CreationDate int	`json:"creationDate"`
	FirstAlbum string	`json:"firstAlbum"`
	Locations locations
	ConcertDates dates
	Relations relations
}

type dates struct {
	Id			int				`json:"id"`
	Dates		[]string		`json:"dates"`
}

type relations struct {
	Id			int						`json:"id"`
	Relations	map[string][]string		`json:"datesLocations"`
}

type locations struct {
	Id			int				`json:"id"`
	Locations	[]string		`json:"locations"`
	Dates		string			`json:"dates"`
	Country []string
	ConcertCityLocations []string
}


