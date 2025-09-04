package web

import (
	str "groupietrack/datastruct"
	"net/http"
	"os"
	"path"
	"strconv"
	"text/template"
)

var data str.PageData

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	fileSystem := http.Dir("./templates")
	fileServer := http.FileServer(fileSystem)
	_, err := fileSystem.Open(path.Clean(r.URL.Path))
	// Si index.html n'existe pas ou que le chemin de l'url n'est ni / ni /artistPage, je redirige sur la page 404.
	if os.IsNotExist(err) && r.URL.Path != "/artistPage" && r.URL.Path != "/" && r.URL.Path != "/results-search" {
		http.Redirect(w, r, "/404Page.html", http.StatusSeeOther)
		return
	}

	// Si le chemin de l'url correspond bien à index.html ou / alors j'exécute le template index.html.
	if r.URL.Path == "/index.html" || r.URL.Path == "/" {
		tmpl, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// Pour ajouter les données de []Artists au template.
		data = str.PageData{
			ArtistHTML: str.GetDatas(),
		}

		tmpl.Execute(w, data)

	  	return
	}
	
	fileServer.ServeHTTP(w, r)
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