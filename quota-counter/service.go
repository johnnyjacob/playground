package main

type QuotaService interface {
	Count(string) (int, error)
}

type quotaService struct{}

func (quotaService) Count(string) (int, error) {
	return 0, nil
}

type ServiceMiddleware func(QuotaService) QuotaService
