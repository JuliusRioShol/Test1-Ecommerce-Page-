package main

import (
	"context"
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/JuliusRioShol/E-commerceApp/internal/data"
)

type application struct {
	logger        *slog.Logger
	productM      *data.ProductModel // Product model
	orderM        *data.OrderModel   // Order Model
	addr          *string
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Read in our DSN
	dsn := flag.String("dsn", "", "PostgreSQL DSN")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// the call to openDB() sets up our connection pool
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("database connection pool established")
	templateCache, err := newTemplateCache()
	/*
		WE NEED TO CONNECT OUR DB HERE !!
		sudo -u postgres psql
		FOR TESTING: db:testing;
		CREATE ROLE testing WITH LOGIN PASSWORD 'testing';
		CREATE EXTENSION IF NOT EXISTS citext;
		psql --host=localhost --dbname=testing --username=testing

		go get github.com/lib/pq@v1

		ALTER DATABASE testing OWNER TO testing;
		GRANT CREATE ON DATABASE testing TO testing;

	*/

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close() // Release the db resources before exiting
	app := &application{
		logger:        logger,
		productM:      &data.ProductModel{DB: db}, // Our Product Model
		orderM:        &data.OrderModel{DB: db},   // Our order Model
		addr:          addr,
		templateCache: templateCache,
	}

	err = app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

}

// This function opens our DB connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ping the db with 5 second timeout
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}
	// Now return the connection pool atp
	return db, nil
}
