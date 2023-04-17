package structs

type Player struct {
	ID           int
	Name         string
	TeamID       int
	Team         Team
	JerseyNumber int
	BirthYear    int
}

var forenames = [20]string{"Anna", "Lisa", "Emil", "Carl", "Rick",
	"Elon", "Steve", "Juli", "Filip", "Stefan", "Anders", "Annika",
	"Lars", "Mick", "Carola", "Thea", "Lola", "Stig", "Josefin", "Helmer"}

var surnames = [20]string{"Lindgren", "Lindqvist", "Larsson", "Eliasson", "Häggkvist", "Andersson", "Svensson",
	"Höglund", "Löfven", "Thunberg", "Gustafsson", "Djarin", "Skywalker", "Jones", "Wayne", "Ågren", "Munther",
	"Maggio", "Timell", "Johansson"}