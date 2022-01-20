package repository

import (
	"context"
	"database/sql"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/product"
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
	prod := &models.Product{}
	err := r.dbConn.QueryRow(
		`SELECT id, title, price FROM products 
                WHERE id=$1`, id).
		Scan(&prod.Id, &prod.Title, &prod.Price)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func (r *ProductPgRepository) Insert(product models.Product) (uint64, error) {
	tx, err := r.dbConn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return 0, err
	}
	var lastId int64
	err = tx.QueryRow(`INSERT INTO products(title, price) 
                              VALUES($1, $2) RETURNING id`,
		product.Title, product.Price).Scan(&lastId)
	if err != nil {
		if rollBackError := tx.Rollback(); rollBackError != nil {
			log.Fatal(rollBackError.Error())
		}
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return uint64(lastId), nil
}

func (r *ProductPgRepository) UpdateById(productId uint64, updatedProduct models.Product) error {
	tx, err := r.dbConn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		`UPDATE products
			    SET title=$1, price=$2
			    WHERE id=$3 `,
		updatedProduct.Title, updatedProduct.Price, productId)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatal(rollbackErr)
		}
		return nil
	}
	if err := tx.Commit(); err != nil {
		return nil
	}
	return nil
}

func (r *ProductPgRepository) DeleteById(id uint64) error {
	tx, err := r.dbConn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return err
	}
	_, err = tx.Exec(
		`DELETE FROM products
			    WHERE id=$1 `, id)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatal(rollbackErr)
		}
		return nil
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
