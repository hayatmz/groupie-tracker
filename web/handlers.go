package web

import (
	str "groupietrack/datastruct"
	"net/http"
	"strconv"
	"text/template"
)

var data str.PageData

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	// Si le chemin de l'url n'existe pas et n'est pas une route valide, rediriger vers 404
	if r.URL.Path != "/" && r.URL.Path != "/index.html" && r.URL.Path != "/artistPage" && r.URL.Path != "/results-search" {
		http.Redirect(w, r, "/404Page.html", http.StatusSeeOther)
		return
	}

	// Si c'est la page index, parser le template
	if r.URL.Path == "/" || r.URL.Path == "/index.html" {
		tmpl, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data = str.PageData{
			ArtistHTML: str.GetDatas(),
		}
		tmpl.Execute(w, data)
		return
	}
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Je récupère les données de "value="{{$element.ID}}"" pour convertir l'id en int puis j'exécute artistPage.html avec les données liées à l'id selectionné.
	idform := r.FormValue("artistPage")
	id, _ := strconv.Atoi(idform)
	pageID := data.ArtistHTML[id-1]
	
	tmpl, err := template.ParseFiles("./templates/artistPage.html")
	if err != nil {
		http.Redirect(w, r, "/500Page.html", http.StatusSeeOther)
		return
	}
	
	tmpl.ExecuteTemplate(w, "artistPage.html", pageID)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("query")

	filteredData := SearchBar(data.ArtistHTML, query)

	tmpl, err := template.ParseFiles("./templates/recherche.html")
	if err != nil {
		// http.Redirect(w, r, "/500Page.html", http.StatusSeeOther)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.ExecuteTemplate(w, "recherche.html" ,filteredData)
}