package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go_practice/11_clean_arch_mongo_db/internal"
	"go_practice/11_clean_arch_mongo_db/internal/models"
)

type ProductMongoRepository struct {
	coll *mongo.Collection
}

func NewProductMongoRepository(collection *mongo.Collection) internal.ProductRepository {
	return &ProductMongoRepository{
		coll: collection,
	}
}

func (r *ProductMongoRepository) SelectAll() ([]*models.Product, error) {
	products := make([]*models.Product, 0, 10)
	res, err := r.coll.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = res.All(context.Background(), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductMongoRepository) SelectById(id string) (*models.Product, error) {
	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	product := &models.Product{}
	err = r.coll.FindOne(context.Background(), bson.D{{"_id", hexId}}).Decode(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductMongoRepository) Insert(product *models.Product) (string, error) {
	id := primitive.NewObjectID()
	product.Id = id
	_, err := r.coll.InsertOne(context.Background(), product)
	if err != nil {
		return "", nil
	}
	return id.Hex(), nil
}

func (r *ProductMongoRepository) Update(product *models.Product) (int64, error) {
	_, err := r.coll.ReplaceOne(context.Background(), bson.M{"_id": product.Id}, product)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (r *ProductMongoRepository) DeleteById(id string) (int64, error) {
	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}
	_, err = r.coll.DeleteOne(context.Background(), bson.M{"_id": hexId})
	if err != nil {
		return 0, err
	}
	return 1, nil
}
