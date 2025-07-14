// delivery/service/service.go
package service

import (
	"context"
	"sync"

	"target-engine/delivery/model"
	"target-engine/pkg/matcher"
)

type DeliveryService interface {
	GetMatchingCampaigns(ctx context.Context, req model.DeliveryRequest) ([]model.Campaign, error)
	LoadCampaigns(campaigns []model.Campaign)
}

type deliveryService struct {
	mu        sync.RWMutex
	campaigns []model.Campaign
}

func NewDeliveryService() DeliveryService {
	return &deliveryService{
		campaigns: []model.Campaign{},
	}
}

func (s *deliveryService) LoadCampaigns(campaigns []model.Campaign) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.campaigns = campaigns
}

func (s *deliveryService) GetMatchingCampaigns(ctx context.Context, req model.DeliveryRequest) ([]model.Campaign, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var matches []model.Campaign
	for _, cmp := range s.campaigns {
		if matcher.MatchCampaign(req, cmp) {
			matches = append(matches, cmp)
		}
	}
	return matches, nil
}
