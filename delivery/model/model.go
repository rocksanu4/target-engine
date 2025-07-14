package model


type Campaign struct {
	ID               string   `json:"id"`
	Status           string   `json:"status"` // ACTIVE, PAUSED, etc.
	AppInclusion     []string `json:"app_inclusion,omitempty"`
	AppExclusion     []string `json:"app_exclusion,omitempty"`
	CountryInclusion []string `json:"country_inclusion,omitempty"`
	CountryExclusion []string `json:"country_exclusion,omitempty"`
	OSInclusion      []string `json:"os_inclusion,omitempty"`
	OSExclusion      []string `json:"os_exclusion,omitempty"`
}

type TargetingRule struct {
    CampaignID string
    Rules      []Rule
}

type Rule struct {
    Dimension string
    Type      string
    Values    []string
}

type DeliveryRequest struct {
    App     string `json:"app"`
    Country string `json:"country"`
    OS      string `json:"os"`
}

type Response struct {
	Data		any			`json:"data"`
}