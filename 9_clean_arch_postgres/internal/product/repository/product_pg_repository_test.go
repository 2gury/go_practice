package repository

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go_practice/9_clean_arch_db/internal/models"
	"testing"
)

func TestProductPgRepository_SelectAll(t *testing.T) {
	type mockBehaviour func(mock sqlmock.Sqlmock, products []*models.Product)
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		outputProducts []*models.Product
		expError error
	}{
		{
			name: "OK",
			mockBehaviour: func(mock sqlmock.Sqlmock, products []*models.Product) {
				rows := sqlmock.NewRows([]string{"id", "title", "price"})
				for _, product :=  range products {
				rows.AddRow(product.Id, product.Title,  product.Price)
			}
			mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
		},
			outputProducts: []*models.Product {
				&models.Product{
					Id: 1,
					Title: "WB",
					Price: 120,
				},
				&models.Product{
					Id: 2,
					Title: "Ozon",
					Price: 500,
				},
			},
			expError: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			productRep := NewProductPgRepository(db)
			testCase.mockBehaviour(mock, testCase.outputProducts)
			productsFromDb, err := productRep.SelectAll()

			assert.Equal(t, productsFromDb, testCase.outputProducts)
			assert.Equal(t, err, testCase.expError)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("were met expectation: %s", err)
			}
		})
	}
}

func TestProductPgRepository_SelectById(t *testing.T) {
	type mockBehaviour func(mock sqlmock.Sqlmock, inProductId uint64, outProduct *models.Product)

	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inputProductId uint64
		outputProduct *models.Product
		expError error
	}{
		{
			name: "OK",
			mockBehaviour: func(mock sqlmock.Sqlmock, inProductId uint64, outProduct *models.Product) {
				rows := sqlmock.NewRows([]string{"id", "title", "price"})
				rows.AddRow(outProduct.Id, outProduct.Title,  outProduct.Price)
				mock.ExpectQuery(`SELECT`).WithArgs(inProductId).WillReturnRows(rows)
			},
			inputProductId: 1,
			outputProduct:&models.Product{
				Id:    1,
				Title: "WB",
				Price: 120,
			},
			expError: nil,
		},
		{
			name: "Error: ErrNoRows",
			mockBehaviour: func(mock sqlmock.Sqlmock, inProductId uint64, outProduct *models.Product) {
				mock.ExpectQuery(`SELECT`).WithArgs(inProductId).WillReturnError(sql.ErrNoRows)
			},
			inputProductId: 1000,
			outputProduct: nil,
			expError: sql.ErrNoRows,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			productRep := NewProductPgRepository(db)
			testCase.mockBehaviour(mock, testCase.inputProductId, testCase.outputProduct)
			productFromDb, err := productRep.SelectById(testCase.inputProductId)

			assert.Equal(t, productFromDb, testCase.outputProduct)
			assert.Equal(t, err, testCase.expError)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("were met expectation: %s", err)
			}
		})
	}
}

func TestProductPgRepository_Insert(t *testing.T) {
	type mockBehaviour func(mock sqlmock.Sqlmock, inProduct *models.Product, outProductId uint64)

	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inputProduct *models.Product
		outputProductId uint64
		expError error
	}{
		{
			name: "OK",
			mockBehaviour: func(mock sqlmock.Sqlmock, inProduct *models.Product, outProductId uint64) {
				mock.ExpectBegin()
				insertedProductId := sqlmock.NewRows([]string{"id"}).AddRow(outProductId)
				mock.ExpectQuery(`INSERT INTO products`).
					WithArgs(inProduct.Title, inProduct.Price).
					WillReturnRows(insertedProductId)
				mock.ExpectCommit()
			},
			inputProduct: &models.Product{
				Title: "WB",
				Price: 120,
			},
			outputProductId: 1,
			expError: nil,
		},
		{
			name: "Error: No unique",
			mockBehaviour: func(mock sqlmock.Sqlmock, inProduct *models.Product, outProductId uint64) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO products`).
					WithArgs(inProduct.Title, inProduct.Price).
					WillReturnError(fmt.Errorf("error no unique"))
				mock.ExpectRollback()
			},
			inputProduct: &models.Product{
				Title: "WB",
				Price: 120,
			},
			outputProductId: 0,
			expError: fmt.Errorf("error no unique"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			productRep := NewProductPgRepository(db)
			testCase.mockBehaviour(mock, testCase.inputProduct, testCase.outputProductId)
			idProductFromDb, err := productRep.Insert(testCase.inputProduct)

			assert.Equal(t, idProductFromDb, testCase.outputProductId)
			assert.Equal(t, err, testCase.expError)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("were met expectation: %s", err)
			}
		})
	}
}

func TestProductPgRepository_UpdateById(t *testing.T) {
	type mockBehaviour func(mock sqlmock.Sqlmock, inProductId uint64, inProduct *models.Product)

	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inputProductId uint64
		inputProduct *models.Product
		expError error
	}{
		{
			name: "OK",
			mockBehaviour: func(mock sqlmock.Sqlmock, inProductId uint64, inProduct *models.Product) {
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE products`).
					WithArgs(inProduct.Title, inProduct.Price, inProductId).
					WillReturnResult(sqlmock.NewResult(int64(inProductId), 1))
				mock.ExpectCommit()
			},
			inputProductId: 1,
			inputProduct: &models.Product{
				Title: "WB",
				Price: 120,
			},
			expError: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			productRep := NewProductPgRepository(db)
			testCase.mockBehaviour(mock, testCase.inputProductId, testCase.inputProduct)
			err := productRep.UpdateById(testCase.inputProductId, testCase.inputProduct)

			assert.Equal(t, err, testCase.expError)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("were met expectation: %s", err)
			}
		})
	}
}

func TestProductPgRepository_DeleteById(t *testing.T) {
	type mockBehaviour func(mock sqlmock.Sqlmock, inProductId uint64)

	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	testTable := []struct {
		name string
		mockBehaviour mockBehaviour
		inputProductId uint64
		expError error
	}{
		{
			name: "OK",
			mockBehaviour: func(mock sqlmock.Sqlmock, inProductId uint64) {
				mock.ExpectBegin()
				mock.ExpectExec(`DELETE FROM products`).
					WithArgs(inProductId).
					WillReturnResult(driver.ResultNoRows)
				mock.ExpectCommit()
			},
			inputProductId: 1,
			expError: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			productRep := NewProductPgRepository(db)
			testCase.mockBehaviour(mock, testCase.inputProductId)
			err := productRep.DeleteById(testCase.inputProductId)

			assert.Equal(t, err, testCase.expError)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("were met expectation: %s", err)
			}
		})
	}
}
