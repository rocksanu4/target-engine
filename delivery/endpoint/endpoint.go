// delivery/endpoint/endpoint.go
package endpoint

import (
	"context"
	"target-engine/delivery/model"
	"target-engine/delivery/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Match endpoint.Endpoint
}

func MakeDeliveryEndpoints(svc service.DeliveryService) Endpoints {
	return Endpoints{
		Match: makeMatchEndpoint(svc),
	}
}

func makeMatchEndpoint(svc service.DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.DeliveryRequest)
		return svc.GetMatchingCampaigns(ctx, req)
	}
}
