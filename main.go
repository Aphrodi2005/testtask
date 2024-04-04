package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	_ "strings"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "aidos"
)

func main() {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	orders := []int{10, 11, 14, 15}

	for _, orderNumber := range orders {
		rows, err := db.Query("SELECT p.name, p.product_id, s.shelf_name, o.quantity FROM orders o JOIN products p ON o.product_id = p.product_id JOIN shelves s ON p.primary_shelf = s.shelf_name WHERE o.order_number = $1", orderNumber)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		fmt.Printf("=+=+=+=\nСтраница сборки заказов %d\n\n", orderNumber)

		var productName string
		var productID int
		var shelfName string
		var quantity int
		for rows.Next() {
			if err := rows.Scan(&productName, &productID, &shelfName, &quantity); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("===Shelving %s\n", shelfName)
			fmt.Printf("%s (id=%d)\n", productName, productID)
			fmt.Printf("order %d, %d шт\n\n", orderNumber, quantity)
		}
	}
}
