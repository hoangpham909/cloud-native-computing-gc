package structs

type Team struct {
	ID          int
	Name        string
	City        string
	FoundedYear int
}

var teams = [20]string{"Hawks", "Eagels", "Bulls", "Titans",
	"Sharks", "Goats", "Vikings", "Lions", "Giants", "Falcons",
	"Parrots", "Dolphins", "Panthers", "Jaguars", "Capybaras",
	"Dinosaurs", "Rocks", "Meteors", "Alligators", "Raiders"}

var cities = [20]string{"Stockholm", "Linköping", "Göteborg", "Varberg", "Halmstad", "Lund", "Jönköping", "Malmö",
	"Helsingborg", "Luleå", "Kungsbacka", "Borås", "Alingsås", "Uppsala", "Örebro", "Umeå", "Norrköping", "Gävle",
	"Växjö", "Sundsvall"}