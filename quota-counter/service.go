package main

import (
	"fmt"

	"github.com/google/uuid"
)

type QuotaService interface {
	Count(uuid.UUID) (int, error)
}

type quotaService struct{}

func (quotaService) Count(id uuid.UUID) (int, error) {
	fmt.Println(id)
	return 0, nil
}

type ServiceMiddleware func(QuotaService) QuotaService
