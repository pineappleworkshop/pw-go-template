package services

import (
    "{{<service_name>}}/models"
)

type newResourceResponse struct {
	Resource       models.Resource `json:"resource"`
}
