package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"target-engine/delivery/endpoint"
	"target-engine/delivery/model"

	httptransport "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"
)

func NewHTTPHandler(eps endpoint.Endpoints, logger *zap.Logger) http.Handler {
	r := http.NewServeMux()

	r.Handle("/delivery", httptransport.NewServer(
		eps.Match,
		decodeDeliveryRequest,
		encodeResponse,
	))

	return r
}

func decodeDeliveryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.DeliveryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}
