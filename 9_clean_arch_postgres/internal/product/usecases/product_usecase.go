package usecases

import (
	"go_practice/8_clean_arch/internal/consts"
	"go_practice/8_clean_arch/internal/helpers/errors"
	"go_practice/8_clean_arch/internal/models"
	"go_practice/8_clean_arch/internal/product"
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
		return 0, errors.Get(consts.CodeBadRequest)
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
		return nil, errors.Get(consts.CodeInternalError)
	}
	if prod == nil {
		return nil, errors.Get(consts.CodeProductDoesNotExist)
	}
	return prod, nil
}
func (u *ProductUsecase) UpdateById(productId uint64, updatedProduct models.Product) *errors.Error {
	if updatedProduct.Price <= 0 || updatedProduct.Title == "" {
		return errors.Get(consts.CodeBadRequest)
	}
	if _, err := u.GetById(productId); err != nil {
		return err
	}
	isUpdated, err := u.productRep.UpdateById(productId, updatedProduct)
	if err != nil {
		return errors.Get(consts.CodeInternalError)
	}
	if !isUpdated {
		return errors.Get(consts.CodeProductDoesNotExist)
	}
	return nil
}

func (u *ProductUsecase) DeleteById(id uint64) *errors.Error {
	if _, err := u.GetById(id); err != nil {
		return err
	}
	isDeleted, err := u.productRep.DeleteById(id)
	if err != nil {
		return errors.Get(consts.CodeInternalError)
	}
	if !isDeleted {
		return errors.Get(consts.CodeProductDoesNotExist)
	}
	return nil
}
