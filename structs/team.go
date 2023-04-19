package structs

type Team struct {
	ID          int
	Name        string
	City        City
	FoundedYear int
}

var names = [20]string{"Hawks", "Eagels", "Bulls", "Titans",
	"Sharks", "Goats", "Vikings", "Lions", "Giants", "Falcons",
	"Parrots", "Dolphins", "Panthers", "Jaguars", "Capybaras",
	"Dinosaurs", "Rocks", "Meteors", "Alligators", "Raiders"}
