package matcher

import (
	"strings"
	"target-engine/delivery/model"
)

func MatchCampaign(req model.DeliveryRequest, campaign model.Campaign) bool {
	if !campaign.Active {
		return false
	}

	if !matchRule(req.App, campaign.TargetRules.IncludeApps, campaign.TargetRules.ExcludeApps) {
		return false
	}

	if !matchRule(req.Country, campaign.TargetRules.IncludeCountries, campaign.TargetRules.ExcludeCountries) {
		return false
	}

	if !matchRule(req.Os, campaign.TargetRules.IncludeOs, campaign.TargetRules.ExcludeOs) {
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
