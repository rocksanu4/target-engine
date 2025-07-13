package model


type Campaign struct {
	CampaignID	string		`json:"campaignId"`
	Name		string		`json:"name"`
	Image		string		`json:"image"`
	CTA			string		`json:"cta"`
	Active		bool		`json:"active"`
}

type TargetRule struct {

}


type Response struct {
	Data		any			`json:"data"`
}