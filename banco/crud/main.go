package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// insert
	product := NewProduct("Notebook", 1899.90)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
	// update
	product.Price = 100.0
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}
	// select one
	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Product: %v, possui o reco de R$ %.2f", p.Name, p.Price)
	// select all
	products, err := selectAllProduct(db)
	if err != nil {
		panic(err)
	}
	for _, p := range products {
		fmt.Printf("Product: %v, possui o preco de R$ %.2f\n", p.Name, p.Price)
	}
	// delete
	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	//preper stetment - protecao contra sql inject
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		panic(err)
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	//preper stetment - protecao contra sql inject
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		panic(err)
	}
	return nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	//preper stetment - protecao contra sql inject
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	var p Product
	//seleciona apenas uma linha e seta os dados na struct
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	// err = stmt.QueryRowContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		panic(err)
	}
	return &p, nil
}

func selectAllProduct(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	//preper stetment - protecao contra sql inject
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	return nil
}
