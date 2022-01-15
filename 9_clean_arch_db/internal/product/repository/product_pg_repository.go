package repository

import (
	"context"
	"database/sql"
	"go_practice/8_clean_arch/internal/models"
	"go_practice/8_clean_arch/internal/product"
	"log"
)

type ProductPgRepository struct {
	dbConn *sql.DB
}

func NewProductPgRepository(conn *sql.DB) product.ProductRepository {
	return &ProductPgRepository{
		dbConn: conn,
	}
}

func (r *ProductPgRepository) SelectAll() ([]*models.Product, error) {
	rows, err := r.dbConn.Query(
		`SELECT id, title, price FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		product := &models.Product{}
		err := rows.Scan(&product.Id, &product.Title, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductPgRepository) SelectById(id uint64) (*models.Product, error) {
	product := &models.Product{}
	err := r.dbConn.QueryRow(
		`SELECT id, title, price FROM products WHERE id=$1`, id).
		Scan(&product.Id, &product.Title, &product.Price)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductPgRepository) Insert(product models.Product) (uint64, error) {
	var lastId uint64
	err := r.dbConn.QueryRow(
		`INSERT INTO products ("title", "price") VALUES ($1, $2) RETURNING id`, product.Title, product.Price).
		Scan(&lastId)
	if err != nil {
		return 0, err
	}
	return lastId, nil
}

func (r *ProductPgRepository) UpdateById(productId uint64, updatedProduct models.Product) (int, error) {
	tx, err := r.dbConn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return 0, err
	}
	result, err := tx.Exec(
		`UPDATE products
				SET title=$1, price=$2
				WHERE id=$3 `,
				updatedProduct.Title, updatedProduct.Price, productId)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatal(rollbackErr)
		}
		return 0, nil
	}
	if err := tx.Commit(); err != nil {
		return 0, nil
	}
	countAffectetdRows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(countAffectetdRows), nil
}

func (r *ProductPgRepository) DeleteById(id uint64) (int, error) {
	tx, err := r.dbConn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return 0, err
	}
	result, err := tx.Exec(
		`DELETE FROM products
				WHERE id=$1 `, id)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatal(rollbackErr)
		}
		return 0, nil
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	countAffectedRows, err := result.RowsAffected()
	return int(countAffectedRows), nil
}
