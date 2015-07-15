package sobjects

import (
	"encoding/xml"
	"io/ioutil"
)

type CampaignMemberSet struct {
	Records []CampaignMember `xml:"records"`
}

type CampaignMember struct {
	Id                 string
	CampaignId         string
	ContactId          string
	CurrencyIsoCode    string
	FirstRespondedDate string
	HasResponded       bool
	LeadId             string
	RecordTypeId       string
	Name               string
	Status             string
}

func NewCampaignMemberSetFromXml(filepath string) (CampaignMemberSet, error) {
	set := CampaignMemberSet{}
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return set, err
	}
	xml.Unmarshal(bytes, &set)
	return set, nil
}
