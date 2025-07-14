package matcher

import (
	"strings"
	"target-engine/delivery/model"
)

func MatchCampaign(req model.DeliveryRequest, campaign model.Campaign) bool {
	if campaign.Status != "ACTIVE" {
		return false
	}

	if !matchRule(req.App, campaign.AppInclusion, campaign.AppExclusion) {
		return false
	}

	if !matchRule(req.Country, campaign.CountryInclusion, campaign.CountryExclusion) {
		return false
	}

	if !matchRule(req.OS, campaign.OSInclusion, campaign.OSExclusion) {
		return false
	}

	return true
}

func matchRule(value string, include []string, exclude []string) bool {
	val := strings.ToLower(value)

	for _, ex := range exclude {
		if val == strings.ToLower(ex) {
			return false
		}
	}

	if len(include) == 0 {
		return true // no inclusion filter => all allowed
	}

	for _, inc := range include {
		if val == strings.ToLower(inc) {
			return true
		}
	}
	return false
}
