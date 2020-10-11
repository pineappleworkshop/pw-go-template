package models

import (

	"time"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"{{<service_name>}}/stores"

	"github.com/labstack/echo"
)

type Product struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Code 	 string `json:"code" bson:"code"`
	CreatedAt  time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at" bson:"updated_at"`
 
}

// product model constructor
func NewProduct() (*Product, error) {
	return &Product{
		ID:        primitive.NewObjectID(),
		Name:      "test1",
		Code: 	   "1234567",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// save product example
func (p *Product) Save(c echo.Context) error {
	client := stores.DB.Mongo.Client
	collection := client.Database(stores.DB_NAME).Collection(stores.DB_COLLECTION_PRODUCTS)
	if _, err := collection.InsertOne(context.TODO(),p); err != nil {
		return err
	}

	return nil
}

type newProductResponse struct {
	Product        Product `json:"product"`
}