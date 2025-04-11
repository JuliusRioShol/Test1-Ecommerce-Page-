package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/JuliusRioShol/E-commerceApp/internal/data"
)

func (app *application) createOrder(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	productID, err := strconv.Atoi(r.FormValue("product_id"))
	quantity, err2 := strconv.Atoi(r.FormValue("quantity"))
	price, err3 := strconv.ParseFloat(r.FormValue("price"), 64)
	size := r.FormValue("size")

	if err != nil || err2 != nil || err3 != nil || size == "" {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	order := data.Order{
		UserID:    1, // TODO: Replace with actual authenticated user ID
		ProductID: productID,
		Quantity:  quantity,
		Size:      size,
		Price:     price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// save to DB
	err = app.orderM.Insert(order)
	if err != nil {
		// Write appropriate error logger
		return
	}
	// We redirect to a success page
	http.Redirect(w, r, "/thank-you", http.StatusSeeOther)
}
