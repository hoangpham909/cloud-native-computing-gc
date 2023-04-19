package structs

import (
	"math/rand"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type City struct{
	Id int
	Name string
}

var cities = [20]string{"Stockholm", "Linköping", "Göteborg", "Varberg", "Halmstad", "Lund", "Jönköping", "Malmö",
	"Helsingborg", "Luleå", "Kungsbacka", "Borås", "Alingsås", "Uppsala", "Örebro", "Umeå", "Norrköping", "Gävle",
	"Växjö", "Sundsvall"}

func RandomCity()(city string){
	
	randNumber := rand.Int63n(20)

	city = cities[randNumber]

	return
}

func AddCities(){
	db, err := gorm.Open(sqlite.Open("db/gc.db"), &gorm.Config{})
	if err != nil {
	}

	db.AutoMigrate(&City{})

	for i := 0; i <= len(cities) -1; i++ {
        tempCity := City {Id: i+1, Name: cities[i]}
		db.Create(&tempCity)
    }
}