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

func NewCampaignMemberSetFromXml(bytes []byte) (CampaignMemberSet, error) {
	set := CampaignMemberSet{}
	err := xml.Unmarshal(bytes, &set)
	return set, err
}

func NewCampaignMemberSetFromXmlFile(filepath string) (CampaignMemberSet, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return CampaignMemberSet{}, err
	}
	return NewCampaignMemberSetFromXml(bytes)
}
