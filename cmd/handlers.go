package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) registerPage(w http.ResponseWriter, r *http.Request) {
	// ask for the data container
	data := newTemplateData()
	data.Title = "Register User"
	data.HeaderText = "Please full form to register"

	err := app.render(w, http.StatusOK, "registration.tmpl", data)
	if err != nil {
		app.logger.Error("failed to render registration page.", "template", "registration.tmpl", "error", err, "url", r.URL.Path, "method", r.Method)
		http.Error(w, "Internal Server Error with registration template", http.StatusInternalServerError)
		return
	}
}

/*DEALING WITH SERVING THE SHOP PAGE*/
func (app *application) shopPage(w http.ResponseWriter, r *http.Request) {
	// Get all the product entries from the db
	product, err := app.productM.GetAllProducts()
	if err != nil {
		app.logger.Error("Failed to fetch products", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Prepare data to pass to the template
	// Ask for the data container
	data := newTemplateData()
	data.Title = "Shop"
	data.HeaderText = " Product in store"
	data.Products = product

	// Render the template to display the products
	err = app.render(w, http.StatusOK, "shop.tmpl", data)
	if err != nil {
		app.logger.Error("failed to render shop page.", "template", "shop.tmpl", "error", err, "url", r.URL.Path, "method", r.Method)
		http.Error(w, "Internal Server Error with shop template", http.StatusInternalServerError)
		return
	}
}

/* DEALING WITH SERVING THE PRODUCT IN DETAIL PAGE*/
func (app *application) productDetails(w http.ResponseWriter, r *http.Request) {
	// Extract product ID from URL ( assuming there is a dynamic route:  /product/{id}")

	path := r.URL.Path
	parts := strings.Split(path, "/") // SPLIT the URL into parts using "/"
	if len(parts) < 3 {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	productID := parts[2]
	log.Println("Extracted ID from URL:", productID)

	id, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID format", http.StatusBadRequest)
		return
	}

	// Fetch product data from the database,
	product, err := app.productM.GetProductByID(id)

	if err != nil {
		app.logger.Error("Failed to fetch products", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Prepare data to pass to the template
	data := newTemplateData()
	data.Title = "Product"
	data.HeaderText = " "
	data.Product = product

	// Render the template to display the products
	err = app.render(w, http.StatusOK, "product.tmpl", data)
	if err != nil {
		app.logger.Error("failed to render products page.", "template", "product.tmpl", "error", err, "url", r.URL.Path, "method", r.Method)
		http.Error(w, "Internal Server Error with products template", http.StatusInternalServerError)
		return
	}
}

func (app *application) successPage(w http.ResponseWriter, r *http.Request) {
	data := newTemplateData()
	data.Title = "Success"
	data.HeaderText = " You have made your order! Go back to shop page to continue shopping"
	// Render the success template to display
	err := app.render(w, http.StatusOK, "success.tmpl", data)

	if err != nil {
		app.logger.Error("failed to render success Page.", "template", "shop success.tmpl", "error", err, "url", r.URL.Path, "method", r.Method)
		http.Error(w, "Internal Server Error with shop template", http.StatusInternalServerError)
		return
	}

}
