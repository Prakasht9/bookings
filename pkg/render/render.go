package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"net/http"
	"path/filepath"

	"github.com/prakasht9/bookings/pkg/config"
	"github.com/prakasht9/bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// get the template cache from app config

	// create template cache
	tc := app.TemplateCache

	// get requested template
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template to browse")
		log.Fatal(ok)
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {

	my_cache := map[string]*template.Template{}
	// build the entire template in the cache for every layout and template

	// get all file named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return my_cache, err
	}

	// range through all the files with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return my_cache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return my_cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.tmpl")
			if err != nil {
				return my_cache, err
			}
		}
		my_cache[name] = ts
	}
	return my_cache, nil
}

// var template_cache = make(map[string]*template.Template)
// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// if we already have the templpate in the cache, return from cache
// 	_, in_map := template_cache[t]
// 	if !in_map {
// 		// need to create, store and return the template
// 		log.Printf("creating template and stroing in cache")
// 		err = CreateTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		log.Println("using cached template")
// 	}
// 	tmpl = template_cache[t]
// 	err = tmpl.Execute(w, nil)
// }

// func CreateTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}
// 	template_cache[t] = tmpl
// 	return nil
// }
