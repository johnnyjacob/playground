package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
)

type quotaResponse struct {
	Status string `json:"status"`
}

type quotaRequest struct {
	ClientID uuid.UUID `json:"clientId"`
}

func makeQuotaEndpoint(svc QuotaService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(quotaRequest)
		_, err := svc.Count(req.ClientID)
		if err != nil {
			return quotaResponse{"deny"}, nil
		}
		return quotaResponse{"allow"}, nil
	}
}

func decodeQuotaRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request quotaRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeQuotaResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var response quotaResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
