package db

import (
	"database/sql"
	"errors"
)

type ProductParams struct {
	ID   int64
	Name sql.NullString
}

const createProduct = `
INSERT INTO products (name) VALUES ($1) RETURNING *
`

func (store *Store) CreateProduct(arg ProductParams) (Product, error) {
	var name string
	if !arg.Name.Valid {
		return Product{}, errors.New("product name can't be empty")
	}
	name = arg.Name.String

	row := store.db.QueryRow(createProduct, name)

	var p Product

	err := row.Scan(&p.ID, &p.Name, &p.CreatedAt, &p.UpdatedAt)

	return p, err
}

const updateProduct = `
UPDATE products SET name = COALESCE($2, name) 
WHERE id = $1
RETURNING *
`

func (store *Store) UpdateProduct(arg ProductParams) (Product, error) {
	row := store.db.QueryRow(updateProduct, arg.ID, arg.Name)

	var p Product

	err := row.Scan(&p.ID, &p.Name, &p.CreatedAt, &p.UpdatedAt)

	return p, err
}

const getProduct = `
SELECT * FROM products 
WHERE id = $1 LIMIT 1
`

func (store *Store) GetProduct(id int64) (Product, error) {
	row := store.db.QueryRow(getProduct, id)

	var p Product

	err := row.Scan(&p.ID, &p.Name, &p.CreatedAt, &p.UpdatedAt)

	return p, err
}

const getProductWithVariant = `
SELECT 
    p.id AS product_id,
    p.name AS product_name,
    v.id AS variant_id,
    v.variant_name,
	v.quantity
FROM products p
LEFT JOIN variants v ON p.id = v.product_id
WHERE p.id = $1
ORDER BY p.id, v.id
`

func (store *Store) GetProductWithVariant(productID int64) ([]ProductWithVariat, error) {
	rows, err := store.db.Query(getProductWithVariant, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []ProductWithVariat{}

	for rows.Next() {
		var pv ProductWithVariat
		if err := rows.Scan(
			&pv.ProductID,
			&pv.ProductName,
			&pv.VariantID,
			&pv.VariantName,
			&pv.Quantity,
		); err != nil {
			return nil, err
		}

		items = append(items, pv)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
