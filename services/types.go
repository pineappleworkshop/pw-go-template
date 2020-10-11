package services

import "{{<service_name>}}/models"

type newProductResponse struct {
	Product        models.Product `json:"product"`
}
