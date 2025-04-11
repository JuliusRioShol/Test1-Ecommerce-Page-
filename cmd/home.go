/**
* This file is for dealing with serving the home page.
 */
package main

import (
	"net/http"
)

// Home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Ask for the data container
	data := newTemplateData()
	data.Title = "Julius Shol Ecommerce Page"

	// call the render function
	err := app.render(w, http.StatusOK, "index.tmpl", data)

	if err != nil {
		app.logger.Error("failed to render home page", "template", "home.tmpl", "error", err, "url", r.URL.Path, "method", r.Method)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
