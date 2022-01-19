package usecases

import (
	"database/sql"
	systetmErrors "github.com/pkg/errors"
	"go_practice/9_clean_arch_db/internal/consts"
	"go_practice/9_clean_arch_db/internal/helpers/errors"
	"go_practice/9_clean_arch_db/internal/models"
	"go_practice/9_clean_arch_db/internal/product"
)

type ProductUsecase struct {
	productRep product.ProductRepository
}

func NewProductUsecase(rep product.ProductRepository) product.ProductUsecase {
	return &ProductUsecase{
		productRep: rep,
	}
}

func (u *ProductUsecase) List() ([]*models.Product, *errors.Error) {
	products, err := u.productRep.SelectAll()
	if err != nil {
		return nil, errors.Get(consts.CodeInternalError)
	}
	return products, nil
}

func (u *ProductUsecase) Create(product models.Product) (uint64, *errors.Error) {
	if product.Price <= 0 || product.Title == "" {
		return 0, errors.New(consts.CodeBadRequest, systetmErrors.New(
			"Error when add product. Price should be greater than 0. Title should be not empty"))
	}
	id, err := u.productRep.Insert(product)
	if err != nil {
		return 0, errors.Get(consts.CodeInternalError)
	}
	return id, nil
}
func (u *ProductUsecase) GetById(id uint64) (*models.Product, *errors.Error) {
	prod, err := u.productRep.SelectById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Get(consts.CodeProductDoesNotExist)
		}
		return nil, errors.Get(consts.CodeInternalError)
	}
	return prod, nil
}
func (u *ProductUsecase) UpdateById(productId uint64, updatedProduct models.Product) *errors.Error {
	if updatedProduct.Price <= 0 || updatedProduct.Title == "" {
		return errors.New(consts.CodeBadRequest, systetmErrors.New(
			"Error when add product. Price should be greater than 0. Title should be not empty"))
	}
	if _, err := u.GetById(productId); err != nil {
		return err
	}
	err := u.productRep.UpdateById(productId, updatedProduct)
	if err != nil {
		return errors.Get(consts.CodeInternalError)
	}
	return nil
}

func (u *ProductUsecase) DeleteById(id uint64) *errors.Error {
	if _, err := u.GetById(id); err != nil {
		return err
	}
	err := u.productRep.DeleteById(id)
	if err != nil {
		return errors.Get(consts.CodeInternalError)
	}
	return nil
}
