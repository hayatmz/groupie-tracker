package web

import (
	str "groupietrack/datastruct"
	"strconv"
	"strings"
)

// Filtre les données de []Artists en fonction de la recherche.
func SearchBar(dataTab []str.Artists, query string) []str.Artists {

	dataTab = append(dataTab, data.ArtistHTML...)
	
	// Si la recherche est vide, je renvois toutes les données.
	if query == "" {
		return dataTab
	}

	// Parcourt le tableau des données des artistes pour ajouter uniquement celles correspondant à la query dans un nouveau tableau.
	var searchedData []str.Artists
	for _, elem := range dataTab {
		if strings.Contains(strings.ToLower(elem.Name), strings.ToLower(query)) || strings.Contains(strings.ToLower(strings.Join(elem.Members," ")), strings.ToLower(query)) || strings.Contains(strconv.Itoa(elem.Creation), strings.ToLower(query)) || strings.Contains(strings.ToLower(elem.FirstAlbum), strings.ToLower(query)) || strings.Contains(strings.ToLower(strings.Join(elem.LocationsRecup, " ")), strings.ToLower(query)) || strings.Contains(strings.ToLower(strings.Join(elem.ConcertDatesRecup, " ")), strings.ToLower(query)) {
			if !AlreadyExist(elem.Name, searchedData) {
				searchedData = append(searchedData, elem)
			}
		}
	}

	return searchedData
}

func AlreadyExist(name string, searchedData []str.Artists) bool {

	for _, elem := range searchedData {
		if elem.Name == name {
			return true
		}
	}

	return false
}