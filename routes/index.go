package routes

import "net/http"
import "github.com/flosch/pongo2/v4"

var indexTemplate = pongo2.Must(pongo2.FromFile("templates/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()

    if query.Has("q") {
        // we need to render search results
    } else {
        // we need to render the homepage

        err := indexTemplate.ExecuteWriter(pongo2.Context{}, w)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}
