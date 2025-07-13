package service

import (
	"context"
	"target-engine/delivery/model"
)

type DeliveryService interface {
	GetMatchingCampaigns(ctx context.Context, req model.DeliveryRequest) ([]model.Campaign, error)
}

type deliveryService struct{}

func NewDeliveryService() DeliveryService {
	return &deliveryService{}
}

func (s *deliveryService) GetMatchingCampaigns(ctx context.Context, req model.DeliveryRequest) ([]model.Campaign, error) {
	// TODO: implement actual filtering logic
	return []model.Campaign{}, nil
}
