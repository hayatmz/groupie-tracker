package datastruct

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	artistsURL = "https://groupietrackers.herokuapp.com/api/artists"
	locationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	datesURL = "https://groupietrackers.herokuapp.com/api/dates"
	relationURL = "https://groupietrackers.herokuapp.com/api/relation"
)

// Fonction pour récupérer les données d'un fichier json.
// interface{} signifie qu'il accepte n'importe quel type.
func GetUrls(urls string, target interface{}) error {

	response, err := http.Get(urls)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed: %d", response.StatusCode)
	}

	return json.NewDecoder(response.Body).Decode(target)
}

// Fonction qui récupère toutes les données des liens api.
func GetDatas() []Artists {

	var artists []Artists
	if err := GetUrls(artistsURL, &artists); err != nil {
		log.Fatal(err)
	}

	var locations Locations
	if err := GetUrls(locationsURL, &locations); err != nil {
		// log.Fatal(err)
		log.Fatal(err)
	}

	var concertDates ConcertDates
	if err := GetUrls(datesURL, &concertDates); err != nil {
		log.Fatal(err)
	}

	var relations Relations
	if err := GetUrls(relationURL, &relations); err != nil {
		log.Fatal(err)
	}

	// range sur le tableau de la structure []Artists pour y ajouter les données correspondantes de locations, concertDates et relations selon leur ID.
	for i, elem := range artists {
		for j, v := range Videos {
			if (i == j) {
				artists[i].Link = v
				artists[i].LocationsRecup = locations.Index[elem.ID-1].Locations
				artists[i].ConcertDatesRecup = concertDates.Index[elem.ID-1].Dates
				artists[i].RelationsRecup = relations.Index[elem.ID-1].DatesLocations
			}
		}
	}
	
	 return artists

}