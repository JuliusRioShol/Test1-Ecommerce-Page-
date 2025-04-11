package main

import (
	"github.com/JuliusRioShol/E-commerceApp/internal/data"
)

type TemplateData struct {
	Title      string
	HeaderText string
	Products   []*data.Product // To hold the list of products
	Product    *data.Product   // To hold only one product

	//
	Orders []*data.Order // To hold a list of order. (if we need to)
	Order  *data.Order   // To hold only one order

}

// Factory function
func newTemplateData() *TemplateData {
	return &TemplateData{
		Title:      "Default Title",
		HeaderText: "Default Header Text",
	}
}
