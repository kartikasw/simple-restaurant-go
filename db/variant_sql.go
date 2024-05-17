package db

import (
	"database/sql"
	"errors"
)

type VariantParams struct {
	ID          int64
	VariantName sql.NullString
	Quantity    sql.NullInt32
	ProductID   int64
}

const createVariant = `
INSERT INTO variants (
    variant_name, quantity, product_id
) VALUES (
    $1, $2, $3
) RETURNING *
`

func (store *Store) CreateVariant(arg VariantParams) (Variant, error) {
	var name string
	if !arg.VariantName.Valid {
		return Variant{}, errors.New("variant name can't be empty")

	}
	name = arg.VariantName.String

	var qty int32
	if !arg.Quantity.Valid {
		qty = 0
	}
	qty = arg.Quantity.Int32

	row := store.db.QueryRow(createVariant, name, qty, arg.ProductID)

	var v Variant

	err := row.Scan(&v.ID, &v.VariantName, &v.Quantity, &v.ProductID, &v.CreatedAt, &v.UpdatedAt)

	return v, err
}

const updateVariant = `
UPDATE variants SET 
    variant_name = COALESCE($2, variant_name), 
    quantity = COALESCE($3, quantity)
WHERE id = $1
RETURNING *
`

func (store *Store) UpdateVariant(arg VariantParams) (Variant, error) {
	row := store.db.QueryRow(updateVariant, arg.ID, arg.VariantName, arg.Quantity)

	var v Variant

	err := row.Scan(&v.ID, &v.VariantName, &v.Quantity, &v.ProductID, &v.CreatedAt, &v.UpdatedAt)

	return v, err
}

const deleteVariant = `
DELETE FROM variants WHERE id = $1
`

func (store *Store) DeleteVariant(variantID int64) error {
	_, err := store.db.Exec(deleteVariant, variantID)
	return err
}
