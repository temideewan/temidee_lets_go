package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"

	"temidee_lets_go.temideewan.net/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	Form any
	Flash string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"dateFormat": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// initialize a map as an in memory cache for template
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		// parse the base template file into a template set
		// register the template func with the template by calling new on the template
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		// parse the files into a template set
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		cache[name] = ts
	}
	return cache, nil
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
		// add flash message to the template data if it exists
		Flash: app.sessionManager.PopString(r.Context(), "flash"),
	}
}
