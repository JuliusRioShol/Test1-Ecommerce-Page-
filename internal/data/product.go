package data

import (
	"database/sql"
	"log"
)

// validator

/*
	Here is our Product Model



*/
// We create our "product" struct
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Brand       string  `json:"brand"`
	ImageURL    string  `json:"image_url"`
	Price       float64 `json:"price"`
	Rating      float64 `json:"rating"`
}

type ProductModel struct {
	DB *sql.DB
}

func (m *ProductModel) GetAllProducts() ([]*Product, error) {
	query := `  SELECT id, name, description, brand, image_url, price, rating 
				FROM products;`

	rows, err := m.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]*Product, 0)
	for rows.Next() {
		product := Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Brand, &product.ImageURL, &product.Price, &product.Rating)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		products = append(products, &product)

	}
	return products, err
}

func (m *ProductModel) GetProductByID(id int) (*Product, error) {
	query := `
			SELECT id, name, description, brand, image_url, price, rating 
				FROM products 
				WHERE id = $1;
	`
	rows := m.DB.QueryRow(query, id)

	product := &Product{}
	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Brand,
		&product.ImageURL,
		&product.Price,
		&product.Rating)
	if err != nil {
		return nil, err
	}
	return product, err
}

// func (m *ProductModel) getProductByID(id string ) (*Product, error ) {

// }
