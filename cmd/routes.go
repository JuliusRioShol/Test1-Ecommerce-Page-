package main

import (
	"net/http"

	_ "github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	/* SERVING OUR TMPL PAGES*/
	//Get our home page
	mux.HandleFunc("GET /{$}", app.home)
	//Gets our Registration Page
	mux.HandleFunc("GET /register", app.registerPage)

	// Get out shop page
	mux.HandleFunc("GET /shop", app.shopPage) // Gets products from Database to load into webpage

	// Get the product details page
	mux.HandleFunc("GET /product/{id}", app.productDetails) // Gets products from Database to load into webpage

	// A SIMPLE SUCCES PAGE

	mux.HandleFunc("GET /successPage", app.successPage)
	// Creates a new order:
	mux.HandleFunc("POST /product/order", app.createOrder)

	return app.loggingMiddleware(mux)
}
