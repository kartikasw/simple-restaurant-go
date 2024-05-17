package main

import (
	"database/sql"
	"fmt"
	"log"
	db "simple-restaurant-go/db"
	"simple-restaurant-go/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	fmt.Println("config=", config)

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close()

	store := db.NewStore(conn)

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	p, err := store.CreateProduct(db.ProductParams{Name: sql.NullString{String: "Makanan", Valid: true}})
	if err != nil {
		log.Fatal("CreateProduct error,", err)
	}
	fmt.Println("CreateProduct succeeded, product=", p)

	p, err = store.GetProduct(p.ID)
	if err != nil {
		log.Fatal("GetProduct error,", err)
	}
	fmt.Println("GetProduct succeeded, product=", p)

	p, err = store.UpdateProduct(db.ProductParams{
		ID:   p.ID,
		Name: sql.NullString{String: "Makanan Enak", Valid: true},
	})
	if err != nil {
		log.Fatal("UpdateProduct with name error,", err)
	}
	fmt.Println("UpdateProduct succeeded with some changes, product=", p)

	p, err = store.UpdateProduct(db.ProductParams{ID: p.ID})
	if err != nil {
		log.Fatal("UpdateProduct without name error,", err)
	}
	fmt.Println("UpdateProduct succeeded with no changes, product=", p)

	v1, err := store.CreateVariant(db.VariantParams{
		VariantName: sql.NullString{String: "Bakso", Valid: true},
		Quantity:    sql.NullInt32{Int32: 2, Valid: true},
		ProductID:   p.ID,
	})
	if err != nil {
		log.Fatal("CreateVariant error,", err)
	}
	fmt.Println("CreateVariant succeeded, variant 1=", v1)

	v1, err = store.UpdateVariant(db.VariantParams{
		ID:          v1.ID,
		VariantName: sql.NullString{String: "Bakso Urat", Valid: true},
	})
	if err != nil {
		log.Fatal("UpdateVariant error,", err)
	}
	fmt.Println("UpdateVariant succeeded with changes, variant 1=", v1)

	v2, err := store.CreateVariant(db.VariantParams{
		VariantName: sql.NullString{String: "Mie", Valid: true},
		Quantity:    sql.NullInt32{Int32: 1, Valid: true},
		ProductID:   p.ID,
	})
	if err != nil {
		log.Fatal("CreateVariant error,", err)
	}
	fmt.Println("CreateVariant succeeded, variant 2=", v2)

	pvs, err := store.GetProductWithVariant(p.ID)
	if err != nil {
		log.Fatal("GetProductWithVariant error,", err)
	}
	fmt.Println("GetProductWithVariant succeeded, products=", pvs)

	err = store.DeleteVariant(v1.ID)
	if err != nil {
		log.Fatal("DeleteVariant error,", err)
	}
	fmt.Println("DeleteVariant succeeded")

	pvs, err = store.GetProductWithVariant(p.ID)
	if err != nil {
		log.Fatal("GetProductWithVariant error,", err)
	}
	fmt.Println("GetProductWithVariant succeeded, products=", pvs)
}
